package repositories

import (
	"go-assignment/models"

	"gorm.io/gorm"
)

type AddressRepositoryQ interface {
	GetAddresses(address *[]models.Address, userId string)
	GetAddress(address *models.Address, id int)
}

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (addressRepository *AddressRepository) GetAddresses(addresses *[]models.Address, userId string) {
	addressRepository.DB.Where("user_id = ? ", userId).Find(&addresses)
}

func (addressRepository *AddressRepository) GetAddress(address *models.Address, id int) {
	addressRepository.DB.Where("id = ? ", id).First(address)
}
