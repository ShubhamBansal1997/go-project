package cart

import (
	"go-assignment/models"
	"go-assignment/requests"

	"gorm.io/gorm"
)

type ServiceWrapper interface {
	Create(cartItem *models.Cart)
	Delete(cartItem *models.Cart)
	Update(cartItem *models.Cart, updateCartItemsRequest *requests.UpdateCartItemRequest)
}

type Service struct {
	DB *gorm.DB
}

func NewCartService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
