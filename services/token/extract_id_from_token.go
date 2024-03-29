package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (tokenService *Service) ExtractIDFromToken(requestToken string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenService.config.Auth.AccessSecret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}
	//log.Printf("api X: error %v", claims["id"].(string))
	s := fmt.Sprintf("%.0f", claims["id"])
	return s, nil
}
