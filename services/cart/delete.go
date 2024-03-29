package cart

import "go-assignment/models"

func (cartService *Service) Delete(cartItem *models.Cart) {
	cartService.DB.Delete(cartItem)
}
