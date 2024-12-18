package utils

import (
	"fmt"
	"time"

	"github.com/Utsavch189/api_mysql/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId    string    `json:"sub"`
	Username  string    `json:"name"`
	CreatedAt time.Time `json:"iat"`
	jwt.RegisteredClaims
}

var secretKey = []byte(config.GetEnv("JWT_SECRET"))

func GenerateJWT(userID, username string) (string, error) {

	claims := Claims{
		UserId:    userID,
		Username:  username,
		CreatedAt: time.Now(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myapi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
