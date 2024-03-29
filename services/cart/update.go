package cart

import (
	"go-assignment/models"
	"go-assignment/requests"
)

func (cartService *Service) Update(cartItem *models.Cart, updateCartRequest *requests.UpdateCartItemRequest, userID uint) {
	cartItem.Quantity = uint(updateCartRequest.Quantity)
	cartItem.ProductID = updateCartRequest.ProductID
	cartItem.UserID = userID
	cartService.DB.Save(cartItem)
}
