package application

import "context"

type ItemUsecase interface {
	GetMyInventory(ctx context.Context, userID string) ([]*ItemWithQuantityDTO, error)
	UseItem(ctx context.Context, userID, itemID string) error
	GetItem(ctx context.Context, userID, itemID string) (*ItemDTO, error)
	GetItems(ctx context.Context, itemIDs []string) ([]*ItemDTO, error)
}
