package middleware

import (
	"myproject/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(idGuest int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["idGuest"] = idGuest
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
