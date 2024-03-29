package address

import (
	"go-assignment/models"
)

func (addressService *Service) Create(address *models.Address) {
	addressService.DB.Create(address)
}
