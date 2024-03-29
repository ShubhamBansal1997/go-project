package responses

import "go-assignment/models"

type CartResponse struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func NewCartResponse(cartItems []models.Cart) *[]CartResponse {
	cartResponse := make([]CartResponse, 0)

	for i := range cartItems {
		cartResponse = append(cartResponse, CartResponse{
			ProductID: cartItems[i].ProductID,
			Quantity:  cartItems[i].ID,
		})
	}

	return &cartResponse
}
