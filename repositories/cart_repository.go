package repositories

import (
	"go-assignment/models"

	"gorm.io/gorm"
)

type CartRepositoryQ interface {
	GetCartItemsForUser(cartItems *[]models.Cart, userId string)
	GetCartItem(cartItem *models.Cart, userId uint, productId uint)
}

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (cartRepository *CartRepository) GetCartItemsForUser(cartItems *[]models.Cart, userId string) {
	cartRepository.DB.Where("user_id = ? ", userId).Find(cartItems)
}

func (cartRepository *CartRepository) GetCartItem(cartItem *models.Cart, userId uint, productId uint) {
	cartRepository.DB.Where("user_id = ? ", userId).Where("product_id = ?", productId).First(cartItem)
}
