package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/trinitt/utils"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {

			return c.String(401, "Unauthorized")
		}

		if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.String(401, "Unauthorized")
		}

		token := authHeader[7:]

		userId, err := utils.ValidateToken(token)

		if err != nil {
			return c.String(401, "Unauthorized")
		}

		c.Set("userId", userId)

		return next(c)
	}
}
