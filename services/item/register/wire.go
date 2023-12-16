//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
	"github.com/google/wire"
	"google.golang.org/grpc"

	api "github.com/Kurichi/plesio-monorepo/services/item/adapter/grpc"
	"github.com/Kurichi/plesio-monorepo/services/item/application"
	"github.com/Kurichi/plesio-monorepo/services/item/infra"
)

func New() *grpc.Server {
	wire.Build(
		config.NewDBConfig,
		database.New,
		infra.NewItemRepository,
		infra.NewInventoryRepository,
		domain.NewItemService,
		application.NewItemUsecase,
		api.NewItemController,
		Register,
	)

	return nil
}
