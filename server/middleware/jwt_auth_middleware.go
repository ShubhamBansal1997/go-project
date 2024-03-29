package middleware

import (
	s "go-assignment/server"
	tokenservice "go-assignment/services/token"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func JwtAuthMiddleware(server *s.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			tokenService := tokenservice.NewTokenService(server.Config)
			authorized, err := tokenService.IsAuthorized(authToken)
			if authorized {
				userID, err := tokenService.ExtractIDFromToken(authToken)
				if err != nil {
					c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				log.Printf("api XYYY: error %v", userID)
				c.Set("user", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
