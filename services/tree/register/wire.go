//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/tree/adapter/api"
	"github.com/Kurichi/plesio-monorepo/services/tree/usecase"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	wire.Build(
		usecase.NewTreeUsecase,
		api.NewTreeHandler,
		Register,
	)

	return nil
}
