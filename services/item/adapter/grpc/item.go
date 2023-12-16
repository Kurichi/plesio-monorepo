package api

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/application"
	itempb "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type itemController struct {
	itempb.UnsafeItemServiceServer

	usecase application.ItemUsecase
}

func NewItemController(usecase application.ItemUsecase) itempb.ItemServiceServer {
	return &itemController{
		usecase: usecase,
	}
}

// GetMyInventory implements grpc.ItemServiceServer.
func (ctrl *itemController) GetMyInventory(ctx context.Context, req *itempb.GetMyInventoryRequest) (*itempb.GetMyInventoryResponse, error) {
	inv, err := ctrl.usecase.GetMyInventory(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	items := make([]*itempb.Item, 0, len(inv))
	for _, i := range inv {
		items = append(items, &itempb.Item{
			Id:          i.ID,
			Name:        i.Name,
			Description: i.Description,
			Quantity:    uint32(i.Quantity),
		})
	}

	return &itempb.GetMyInventoryResponse{
		Items: items,
	}, nil
}

// UseItem implements grpc.ItemServiceServer.
func (ctrl *itemController) UseItem(ctx context.Context, req *itempb.UseItemRequest) (*emptypb.Empty, error) {
	if err := ctrl.usecase.UseItem(ctx, req.UserId, req.ItemId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// GetItemsByIDs implements grpc.ItemServiceServer.
func (ctrl *itemController) GetItemsByIDs(ctx context.Context, req *itempb.GetItemsByIDsRequest) (*itempb.GetItemsByIDsResponse, error) {
	dto, err := ctrl.usecase.GetItems(ctx, req.GetItemIds())
	if err != nil {
		return nil, err
	}

	items := make([]*itempb.Item, 0, len(dto))
	for _, i := range dto {
		items = append(items, &itempb.Item{
			Id:          i.ID,
			Name:        i.Name,
			Description: i.Description,
		})
	}

	return &itempb.GetItemsByIDsResponse{
		Items: items,
	}, nil
}
