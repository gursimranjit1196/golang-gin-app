package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

var apiKey string = "0fa2c280-3694-4756-9332-a110af3e8338"

func CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(apiKey))
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(apiKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
