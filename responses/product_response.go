package responses

import (
	"go-assignment/models"
)

type ProductResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Sku         string  `json:"sku"`
	ID          uint    `json:"id"`
}

func NewProductResponse(products []models.Product) *[]ProductResponse {
	productResponse := make([]ProductResponse, 0)
	for i := range products {
		productResponse = append(productResponse, ProductResponse{
			Name:        products[i].Name,
			Description: products[i].Description,
			Image:       products[i].Image,
			Price:       products[i].Price,
			Category:    products[i].Category,
			Sku:         products[i].Sku,
			ID:          products[i].ID,
		})
	}
	return &productResponse

}
