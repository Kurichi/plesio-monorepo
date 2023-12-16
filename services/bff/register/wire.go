//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/item"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/tree"
	"github.com/Kurichi/plesio-monorepo/services/bff/handler/user"
	"github.com/Kurichi/plesio-monorepo/services/bff/middleware"
	"github.com/Kurichi/plesio-monorepo/services/bff/pkg/firebase"
	itemGrpc "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
	treeGrpc "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func New(grpc.ClientConnInterface) *echo.Echo {
	wire.Build(
		// TreeClientの生成
		tree.NewTreeClient,
		treeGrpc.NewTreeServiceClient,

		// UserClientの生成
		user.NewUserClient,

		// ItemClientの生成
		item.NewItemClient,
		itemGrpc.NewItemServiceClient,

		Register,

		// 認証ミドルウェアの生成
		middleware.NewAuthController,
		firebase.InitializeFirebaseApp,
		firebase.InitializeFBAuthClient,
	)

	return nil
}
