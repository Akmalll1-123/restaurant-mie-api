package jwtutil

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`

	jwt.RegisteredClaims
}

func GenerateToken(
	id uint,
	email string,
	role string,
) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{
			ID:   id,
			Role: role,
			RegisteredClaims: jwt.RegisteredClaims{
				Subject: email,
			},
		},
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}
