package router

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/user"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"

	"github.com/labstack/echo/v4"
)

func DefineUserRouter(g *echo.Group, uc *user.UserClient, authHandler *middleware.AuthController) {
	g.POST("/signup", uc.SignUp, authHandler.TokenVerificationMiddleware)
	g.GET("/user/:userId", uc.GetUser, authHandler.TokenVerificationMiddleware)
	g.PUT("/users/me", uc.UpdateUser, authHandler.TokenVerificationMiddleware)
}
