// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/tree/adapter/api"
	"github.com/Kurichi/plesio-monorepo/services/tree/application"
	"github.com/Kurichi/plesio-monorepo/services/tree/infra"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/database"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func New() *grpc.Server {
	dbConfig := config.NewDBConfig()
	db := database.New(dbConfig)
	treeRepository := infra.NewTreeRepository(db)
	treeUsecase := application.NewTreeUsecase(treeRepository, db)
	treeServiceServer := api.NewTreeHandler(treeUsecase)
	server := Register(treeServiceServer)
	return server
}
