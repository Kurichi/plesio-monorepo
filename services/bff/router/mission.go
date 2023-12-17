package router

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/mission"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/labstack/echo/v4"
)

func DefineMissionRouter(g *echo.Group, mc *mission.MissionClient, authHandler *middleware.AuthController) {
	missions := g.Group("/missions")
	missions.GET("", mc.GetMissions, authHandler.TokenVerificationMiddleware)
	missions.POST("/:id", mc.ProgressMission, authHandler.TokenVerificationMiddleware)
	missions.POST("", mc.CreateMission)
}
