package address

import (
	"go-assignment/models"
	"go-assignment/requests"

	"gorm.io/gorm"
)

type ServiceWrapper interface {
	Create(address *models.Address)
	Delete(address *models.Address)
	Update(address *models.Address, updateAddressRequest *requests.UpdateAddressRequest)
}

type Service struct {
	DB *gorm.DB
}

func NewAddressService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
