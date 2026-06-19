package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(
	roles ...string,
) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			user := c.Get("user").(*JWTClaims)

			for _, role := range roles {

				if role == user.Role {
					return next(c)
				}
			}

			return c.JSON(
				http.StatusForbidden,
				map[string]string{
					"message": "forbidden",
				},
			)
		}
	}
}
