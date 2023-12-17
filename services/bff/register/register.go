package register

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/item"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/mission"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/tree"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/user"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/Kurichi/plesio-monorepo/services/bff/router"
	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func Register(
	treeHandler *tree.TreeClient,
	userHandler *user.UserClient,
	itemHandler *item.ItemClient,
	missionHandler *mission.MissionClient,
	authHandler *middleware.AuthController,
) *echo.Echo {
	e := echo.New()

	g := e.Group("/api/v1")
	g.Use(echoMid.Logger())
	e.Use(echoMid.Recover())
	e.Use(echoMid.Gzip())
	e.Use(echoMid.CORSWithConfig(echoMid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization},
	}))

	router.DefineUserRouter(g, userHandler, authHandler)
	router.DefineTreeRouter(g, treeHandler, authHandler)
	router.DefineItemRouter(g, itemHandler, authHandler)
	router.DefineMissionRouter(g, missionHandler, authHandler)
	return e
}
