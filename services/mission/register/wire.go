//go:build wireinject

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
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func New(context.Context) *grpc.Server {
	wire.Build(
		config.NewDBConfig,
		database.New,
		config.NewPusherConfig,
		pubsub.New,
		infra.NewPublisher,
		domain.NewMissionService,
		infra.NewMissionRepository,
		application.NewMissionUsecase,
		api.NewMissionHandler,
		Register,
		wire.Bind(new(domain.TxRepository), new(*database.DB)),
	)

	return nil
}
