package middleware

import (
	"net/http"
	"os"
	jwtutil "restaurant-mie-api/util/jwt"
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

		token, err := jwtlib.ParseWithClaims(
			tokenString,
			&jwtutil.Claims{},
			func(t *jwtlib.Token) (any, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token"})
		}

		claims, ok := token.Claims.(*jwtutil.Claims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid claims",
			})
		}

		c.Set("user", &JWTClaims{
			ID:    claims.ID,
			Email: claims.Subject,
			Role:  claims.Role,
		})

		return next(c)
	}
}
