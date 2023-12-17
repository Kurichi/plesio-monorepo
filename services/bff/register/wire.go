//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/item"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/mission"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/tree"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/user"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/Kurichi/plesio-monorepo/services/bff/pkg/firebase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func New(grpc.ClientConnInterface) *echo.Echo {
	wire.Build(
		// TreeClientの生成
		tree.NewTreeClient,
		NewTreeClient,

		// UserClientの生成
		user.NewUserClient,
		NewUserClient,

		// ItemClientの生成
		item.NewItemClient,
		NewItemClient,

		// MissionClientの生成
		mission.NewMissionClient,
		NewMissionClient,

		Register,

		// 認証ミドルウェアの生成
		middleware.NewAuthController,
		firebase.InitializeFirebaseApp,
		firebase.InitializeFBAuthClient,
	)

	return nil
}
