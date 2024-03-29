package handlers

import (
	"fmt"
	"go-assignment/models"
	"go-assignment/repositories"
	"go-assignment/requests"
	"go-assignment/responses"
	s "go-assignment/server"
	tokenservice "go-assignment/services/token"
	"net/http"

	"github.com/gin-gonic/gin"

	jwtGo "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	server *s.Server
}

func NewAuthHandler(server *s.Server) *AuthHandler {
	return &AuthHandler{server: server}
}

func (authHandler *AuthHandler) Login(c *gin.Context) {
	loginRequest := new(requests.LoginRequest)

	if err := c.ShouldBind(&loginRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := loginRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
		return
	}

	user := models.User{}
	userRepository := repositories.NewUserRepository(authHandler.server.DB)
	userRepository.GetUserByEmail(&user, loginRequest.Email)

	if user.ID == 0 || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil) {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	tokenService := tokenservice.NewTokenService(authHandler.server.Config)
	accessToken, exp, err := tokenService.CreateAccessToken(&user)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	refreshToken, err := tokenService.CreateRefreshToken(&user)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	res := responses.NewLoginResponse(accessToken, refreshToken, exp)

	responses.Response(c, http.StatusOK, res)
	return
}

func (authHandler *AuthHandler) RefreshToken(c *gin.Context) {
	refreshRequest := new(requests.RefreshRequest)
	if err := c.ShouldBind(&refreshRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := jwtGo.Parse(refreshRequest.Token, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(authHandler.server.Config.Auth.RefreshSecret), nil
	})

	if err != nil {
		responses.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	claims, ok := token.Claims.(jwtGo.MapClaims)
	if !ok && !token.Valid {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	user := new(models.User)
	authHandler.server.DB.First(&user, int(claims["id"].(float64)))

	if user.ID == 0 {
		responses.ErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	tokenService := tokenservice.NewTokenService(authHandler.server.Config)
	accessToken, exp, err := tokenService.CreateAccessToken(user)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	refreshToken, err := tokenService.CreateRefreshToken(user)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	res := responses.NewLoginResponse(accessToken, refreshToken, exp)

	responses.Response(c, http.StatusOK, res)
	return
}
