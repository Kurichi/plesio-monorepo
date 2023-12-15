package api

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/application"
	itempb "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
)

type itemController struct {
	itempb.UnsafeItemServiceServer

	usecase application.ItemUsecase
}

func NewItemController(usecase application.ItemUsecase) itempb.ItemServiceServer {
	return &itemController{usecase: usecase}
}

// GetItemByID implements grpc.ItemServiceServer.
func (ctrl *itemController) GetItemByID(ctx context.Context, req *itempb.GetItemRequest) (*itempb.GetItemResponse, error) {
	item, err := ctrl.usecase.GetItemByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	res := &itempb.GetItemResponse{
		Item: &itempb.Item{
			Id:   item.ID,
			Name: item.Name,
		},
	}
	return res, nil
}
