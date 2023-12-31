package router

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/item"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/labstack/echo/v4"
)

func DefineItemRouter(g *echo.Group, ic *item.ItemClient, authHandler *middleware.AuthController) {
	items := g.Group("/items", authHandler.TokenVerificationMiddleware)
	items.GET("", ic.GetItemsHandler)
	items.POST("/:id/use", ic.UseItemHandler)
	items.POST("", ic.CreateItemHandler)
}
