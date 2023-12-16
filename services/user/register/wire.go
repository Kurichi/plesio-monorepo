//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/user/adapter/api"
	"github.com/Kurichi/plesio-monorepo/services/user/application"
	"github.com/Kurichi/plesio-monorepo/services/user/infra"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/database"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/firebase"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	wire.Build(
		config.NewDBConfig,
		database.New,
		firebase.InitializeFBAuthClient,
		firebase.InitializeFirebaseApp,
		infra.NewUserRepository,
		infra.NewIdentityRepository,
		application.NewUserUsecase,
		api.NewUserHandler,
		Register,
	)

	return nil
}
