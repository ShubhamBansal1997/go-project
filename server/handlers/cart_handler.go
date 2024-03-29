package handlers

import (
	"go-assignment/models"
	"go-assignment/repositories"
	"go-assignment/requests"
	"go-assignment/responses"
	"go-assignment/server"
	s "go-assignment/server"
	cartService "go-assignment/services/cart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandlers struct {
	server *s.Server
}

func NewCartHandlers(server *server.Server) *CartHandlers {
	return &CartHandlers{server: server}
}

func (p *CartHandlers) GetCart(c *gin.Context) {
	var cartItems []models.Cart

	userId := c.GetString("user")

	cartRepository := repositories.NewCartRepository(p.server.DB)
	cartRepository.GetCartItemsForUser(&cartItems, userId)

	response := responses.NewCartResponse(cartItems)
	responses.Response(c, http.StatusOK, response)
	return

}

func (p *CartHandlers) UpdateCart(c *gin.Context) {
	updateCartRequest := new(requests.UpdateCartItemRequest)
	if err := c.ShouldBind(&updateCartRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := updateCartRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cart := models.Cart{}
	user := c.GetString("user")
	userId, err := strconv.ParseUint(user, 10, 64)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Invalid User Id")
	}

	cartRepository := repositories.NewCartRepository(p.server.DB)
	cartRepository.GetCartItem(&cart, uint(userId), updateCartRequest.ProductID)
	cartService := cartService.NewCartService(p.server.DB)

	if cart.ID == 0 && updateCartRequest.Quantity != 0 {
		// create a new item in the database
		cartItem := models.Cart{
			UserID:    uint(userId),
			ProductID: updateCartRequest.ProductID,
			Quantity:  updateCartRequest.Quantity,
		}

		cartService.Create(&cartItem)
		responses.MessageResponse(c, http.StatusCreated, "Product added to the cart successfully")
		return
	}
	if updateCartRequest.Quantity == 0 {
		cartService.Delete(&cart)
		responses.MessageResponse(c, http.StatusNoContent, "Product removed from the cart successfully")
		return
	}
	cartService.Update(&cart, updateCartRequest, uint(userId))
	responses.MessageResponse(c, http.StatusOK, "Product Quantity Updated Successfully")
	return
}
