package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (tokenService *Service) IsAuthorized(requestToken string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenService.config.Auth.AccessSecret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
