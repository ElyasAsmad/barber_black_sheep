package helpers

import (
	"barber_black_sheep/model"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var privateKey = []byte("SOME_SECRET_KEY")

func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi("86400")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.UserID,
		"sub":  user.Username,
		"role": user.Role,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
