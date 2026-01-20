package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(userID uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       userID,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	return token.SignedString(secretKey)
}

func VerifyTokenAndGetUsername(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", fmt.Errorf("username not found in token")
		}
		return username, nil
	}

	return "", fmt.Errorf("invalid token")
}
