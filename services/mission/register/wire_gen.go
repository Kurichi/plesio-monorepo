// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package register

import (
	"context"
	"github.com/Kurichi/plesio-monorepo/services/mission/adapter/api"
	"github.com/Kurichi/plesio-monorepo/services/mission/application"
	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
	"github.com/Kurichi/plesio-monorepo/services/mission/infra"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/database"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/pubsub"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func New(contextContext context.Context) *grpc.Server {
	dbConfig := config.NewDBConfig()
	db := database.New(dbConfig)
	missionRepository := infra.NewMissionRepository(db)
	missionService := domain.NewMissionService(missionRepository)
	pusherConfig := config.NewPusherConfig()
	topic := pubsub.New(contextContext, pusherConfig)
	publisher := infra.NewPublisher(topic)
	missionUsecase := application.NewMissionUsecase(missionRepository, missionService, publisher, db)
	missionServiceServer := api.NewMissionHandler(missionUsecase)
	server := Register(missionServiceServer)
	return server
}
