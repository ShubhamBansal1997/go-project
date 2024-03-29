package handlers

import (
	"go-assignment/models"
	"go-assignment/repositories"
	"go-assignment/requests"
	"go-assignment/responses"
	s "go-assignment/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	server *s.Server
}

func NewProductHandlers(server *s.Server) *ProductHandlers {
	return &ProductHandlers{server: server}
}

func (p *ProductHandlers) GetProducts(c *gin.Context) {
	filterProductRequest := new(requests.FilterProductRequest)

	if err := c.ShouldBind(&filterProductRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := filterProductRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var products []models.Product
	filters := map[string]interface{}{
		"category":   filterProductRequest.Category,
		"name":       filterProductRequest.Name,
		"sku":        filterProductRequest.SKU,
		"price__gte": filterProductRequest.PriceGte,
		"price__lte": filterProductRequest.PriceLte,
		"order_by":   filterProductRequest.OrderBy,
	}
	productRepository := repositories.NewProductRepository(p.server.DB)
	productRepository.GetProducts(&products, filters)
	response := responses.NewProductResponse(products)
	responses.Response(c, http.StatusOK, response)
	return

}
