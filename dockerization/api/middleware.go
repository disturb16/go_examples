package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var users = map[string]string{"user1": "password1", "user2": "password2"}

func requestIDHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		requestID := c.Request().Header.Get("X-Request-ID")

		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		c.Response().Header().Set("X-Request-ID", requestID)
		return next(c)
	}
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Request().Header.Get("Authorization")

		if users[user] == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
