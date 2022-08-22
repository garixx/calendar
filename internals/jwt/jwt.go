package jwt

import (
	"calendar/internals/ssm"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey, _ = ssm.GetSecret("stub")

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
