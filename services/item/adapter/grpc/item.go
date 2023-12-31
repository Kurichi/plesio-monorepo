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
	// return &itemController{
	// 	usecase: usecase,
	// }
	return nil
}

// CreateItem implements grpc.ItemServiceServer.
func (ctrl *itemController) CreateItem(ctx context.Context, req *itempb.CreateItemRequest) (*itempb.CreateItemResponse, error) {
	item, err := ctrl.usecase.CreateItem(ctx, req.Name, req.Description, req.Target, int(req.Amount))
	if err != nil {
		return nil, err
	}
	return &itempb.CreateItemResponse{
		Item: &itempb.Item{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		},
	}, nil
}

// GetMyInventory implements grpc.ItemServiceServer.
func (ctrl *itemController) GetMyInventory(ctx context.Context, req *itempb.GetMyInventoryRequest) (*itempb.GetMyInventoryResponse, error) {
	inv, err := ctrl.usecase.GetMyInventory(ctx, req.GetUserId())
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
func (ctrl *itemController) UseItem(ctx context.Context, req *itempb.UseItemRequest) (*itempb.UseItemResponse, error) {
	tree, err := ctrl.usecase.UseItem(ctx, req.GetUserId(), req.GetItemId())
	if err != nil {
		return nil, err
	}

	res := &itempb.UseItemResponse{
		Stage:       int32(tree.Stage),
		Water:       int32(tree.Water),
		Felitilizer: int32(tree.Fertilizer),
	}
	return res, nil
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
