package utils

import (
	"ffserver/env"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateDefaultJwt() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	return token.SignedString(env.AuthSecret)
}
