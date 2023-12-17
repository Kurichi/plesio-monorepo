//go:build wireinject

package register

import (
	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/treepb"
	"github.com/google/wire"

	api "github.com/Kurichi/plesio-monorepo/services/item/adapter/grpc"
	"github.com/Kurichi/plesio-monorepo/services/item/adapter/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/item/application"
	"github.com/Kurichi/plesio-monorepo/services/item/infra"
)

func New() *Server {
	wire.Build(
		config.NewDBConfig,
		database.New,
		infra.NewItemRepository,
		infra.NewInventoryRepository,
		domain.NewItemService,
		treepb.New,
		application.NewItemUsecase,
		api.NewItemController,
		pubsub.NewItemController,
		Register,
		Subscriber,
		NewServer,
		infra.NewTreeService,
		wire.Bind(new(domain.TxRepository), new(*database.DB)),
	)

	return nil
}
