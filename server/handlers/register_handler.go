package handlers

import (
	"go-assignment/models"
	"go-assignment/repositories"
	"go-assignment/requests"
	"go-assignment/responses"
	s "go-assignment/server"
	"go-assignment/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	server *s.Server
}

func NewRegisterHandler(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

// Register godoc
//
//	@Summary		Register
//	@Description	New user registration
//	@ID				user-register
//	@Tags			User Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.RegisterRequest	true	"User's email, user's password"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Router			/register [post]
func (registerHandler *RegisterHandler) Register(c *gin.Context) {
	registerRequest := new(requests.RegisterRequest)

	err := c.ShouldBind(&registerRequest)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := registerRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
		return
	}

	existUser := models.User{}
	userRepository := repositories.NewUserRepository(registerHandler.server.DB)
	userRepository.GetUserByEmail(&existUser, registerRequest.Email)

	if existUser.ID != 0 {
		responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
		return
	}

	userService := user.NewUserService(registerHandler.server.DB)
	if err := userService.Register(registerRequest); err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
		return
	}

	responses.MessageResponse(c, http.StatusCreated, "User successfully created")
	return
}
