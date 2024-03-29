package cart

import "go-assignment/models"

func (cartService *Service) Create(cartItem *models.Cart) {
	cartService.DB.Create(cartItem)
}
