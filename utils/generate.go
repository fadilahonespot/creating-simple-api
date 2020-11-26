package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)


func GenerateToken(id int, pass string) (string, error) {
	claims := jwt.MapClaims{}

	claims["user_id"] = id
	claims["pass"] = pass

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		fmt.Printf("[utils.GenerateToken] Error to signature generate token, %v \n", err)
		return "", fmt.Errorf("Failed generate token")
	}
	return tokenString, nil
}