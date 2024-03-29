package handlers

import (
	"go-assignment/models"
	"go-assignment/repositories"
	"go-assignment/requests"
	"go-assignment/responses"
	s "go-assignment/server"
	addressService "go-assignment/services/address"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandlers struct {
	server *s.Server
}

func NewAddressHandlers(server *s.Server) *AddressHandlers {
	return &AddressHandlers{server: server}
}

func (p *AddressHandlers) CreateAddress(c *gin.Context) {
	createAddressRequest := new(requests.CreateAddressRequest)

	if err := c.ShouldBind(&createAddressRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := createAddressRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
		return
	}
	user := c.GetString("user")
	userId, err := strconv.ParseUint(user, 10, 64) // base 10, up to 64 bits
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Invalid User Id")
		return
	}

	address := models.Address{
		Value:  createAddressRequest.Value,
		UserID: uint(userId),
	}
	addressService := addressService.NewAddressService(p.server.DB)
	addressService.Create(&address)

	responses.MessageResponse(c, http.StatusCreated, "Address successfully created")
	return
}

func (p *AddressHandlers) DeleteAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	address := models.Address{}

	addressRepository := repositories.NewAddressRepository(p.server.DB)
	addressRepository.GetAddress(&address, id)
	if address.ID == 0 {
		responses.ErrorResponse(c, http.StatusNotFound, "Address not found")
		return
	}

	addressService := addressService.NewAddressService(p.server.DB)
	addressService.Delete(&address)

	responses.MessageResponse(c, http.StatusNoContent, "Address deleted successfully")
	return
}

func (p *AddressHandlers) GetAddress(c *gin.Context) {
	var addresses []models.Address
	//user := c.GetUint("user")
	userId := c.GetString("user")

	addressRepository := repositories.NewAddressRepository(p.server.DB)
	addressRepository.GetAddresses(&addresses, userId)

	response := responses.NewAddressResponse(addresses)
	responses.Response(c, http.StatusOK, response)
	return
}

func (p *AddressHandlers) UpdateAddress(c *gin.Context) {
	updateAddressRequest := new(requests.UpdateAddressRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&updateAddressRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := updateAddressRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
		return
	}

	user := c.GetString("user")
	userId, err := strconv.ParseUint(user, 10, 64) // base 10, up to 64 bits
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Invalid User Id")
		return
	}

	address := models.Address{}

	addressRepository := repositories.NewAddressRepository(p.server.DB)
	addressRepository.GetAddress(&address, id)

	if address.ID == 0 {
		responses.ErrorResponse(c, http.StatusNotFound, "Address not found")
		return
	}

	addressService := addressService.NewAddressService(p.server.DB)
	addressService.Update(&address, *updateAddressRequest, uint(userId))

	responses.MessageResponse(c, http.StatusOK, "Address successfully updated")
	return
}
