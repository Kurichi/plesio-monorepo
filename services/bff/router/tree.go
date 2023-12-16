package router

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/tree"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/labstack/echo/v4"
)

func DefineTreeRouter(g *echo.Group, treeHandler *tree.TreeClient, authHandler *middleware.AuthController) {
	g.GET("/trees/me", treeHandler.GetMyTreeHandler, authHandler.TokenVerificationMiddleware)
	g.GET("/users/:userId/tree", treeHandler.GetTreeByUserId, authHandler.TokenVerificationMiddleware)
	g.GET("/trees/ranking", treeHandler.GetTreeRanking, authHandler.TokenVerificationMiddleware)
	g.POST("/trees/plant", treeHandler.PlantTreeHandler, authHandler.TokenVerificationMiddleware)
	g.POST("/trees/init", treeHandler.InitTreeHandler, authHandler.TokenVerificationMiddleware)
}
