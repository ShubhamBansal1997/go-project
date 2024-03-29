package address

import (
	"go-assignment/models"
	"go-assignment/requests"
)

func (addressService *Service) Update(address *models.Address, updateAddressRequest requests.UpdateAddressRequest, userId uint) {
	address.Value = updateAddressRequest.Value
	address.UserID = uint(userId)
	addressService.DB.Save(address)
}
