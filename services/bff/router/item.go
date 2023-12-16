package router

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/item"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/labstack/echo/v4"
)

func DefineItemRouter(g *echo.Group, ic *item.ItemClient, authHandler *middleware.AuthController) {
	g.GET("/items", ic.GetItemsHandler, authHandler.TokenVerificationMiddleware)
	g.GET("/items/:id", ic.GetItemDetailHandler, authHandler.TokenVerificationMiddleware)
	g.DELETE("/items/:id", ic.DeleteItemHandler, authHandler.TokenVerificationMiddleware)
	g.POST("/items/:id/use", ic.UseItemHandler, authHandler.TokenVerificationMiddleware)
}
