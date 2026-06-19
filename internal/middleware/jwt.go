package middleware

import (
	"net/http"
	"os"
	"strings"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	ID    uint
	Email string
	Role  string
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		auth := c.Request().Header.Get("Authorization")

		if auth == "" {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{
					"message": "missing token",
				},
			)
		}

		tokenString := strings.TrimPrefix(
			auth,
			"Bearer ",
		)

		token, err := jwtlib.Parse(
			tokenString,
			func(token *jwtlib.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{
					"message": "invalid token",
				},
			)
		}

		mapClaims := token.Claims.(jwtlib.MapClaims)

		c.Set(
			"user",
			&JWTClaims{
				ID:    uint(mapClaims["id"].(float64)),
				Email: mapClaims["email"].(string),
				Role:  mapClaims["role"].(string),
			},
		)

		return next(c)
	}
}
