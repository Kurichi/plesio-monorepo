package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	client *auth.Client
}

func NewAuthController(client *auth.Client) *AuthController {
	return &AuthController{
		client: client,
	}
}

func (ac *AuthController) TokenVerificationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Required token")
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(token, prefix) {
			return c.JSON(http.StatusUnauthorized, "Invalid authorization header format")
		}

		token = strings.TrimPrefix(token, prefix)

		ctx := context.Background()

		authToken, err := ac.client.VerifyIDToken(ctx, token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("userId", authToken.UID)
		c.Set("token", token)

		return next(c)
	}
}
