package address

import "go-assignment/models"

func (addressService *Service) Delete(address *models.Address) {
	addressService.DB.Delete(address)
}
