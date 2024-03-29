package responses

import "go-assignment/models"

type AddressResponse struct {
	Value string `json:"value"`
	ID    uint   `json:"id"`
}

func NewAddressResponse(addresses []models.Address) *[]AddressResponse {
	addressResponse := make([]AddressResponse, 0)
	for i := range addresses {
		addressResponse = append(addressResponse, AddressResponse{
			Value: addresses[i].Value,
			ID:    addresses[i].ID,
		})
	}
	return &addressResponse
}
