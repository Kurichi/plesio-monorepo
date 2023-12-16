package application

import "context"

type ItemUsecase interface {
	CreateItem(ctx context.Context, name string, description string, target string, amount int) (*ItemDTO, error)
	GetMyInventory(ctx context.Context, userID string) ([]*ItemWithQuantityDTO, error)
	UseItem(ctx context.Context, userID, itemID string) error
	GetItem(ctx context.Context, userID, itemID string) (*ItemDTO, error)
	GetItems(ctx context.Context, itemIDs []string) ([]*ItemDTO, error)
}
