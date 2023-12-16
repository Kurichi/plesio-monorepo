//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/tree/adapter/api"
	"github.com/Kurichi/plesio-monorepo/services/tree/application"
	"github.com/Kurichi/plesio-monorepo/services/tree/domain"
	"github.com/Kurichi/plesio-monorepo/services/tree/infra"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/database"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	wire.Build(
		config.NewDBConfig,
		database.New,
		infra.NewTreeRepository,
		application.NewTreeUsecase,
		api.NewTreeHandler,
		Register,
		wire.Bind(new(domain.TxRepository), new(*database.DB)),
	)

	return nil
}
