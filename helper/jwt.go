package helper

import (
	"go-diary-app/model"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user model.User) (string, error) {
	
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Hour * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)

}