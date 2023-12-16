package application

import "context"

type ItemUsecase interface {
	CreateItem(ctx context.Context, name string, description string, target string, amount int) (*ItemDTO, error)
	GetMyInventory(ctx context.Context, userID string) ([]*ItemWithQuantityDTO, error)
	UseItem(ctx context.Context, userID, itemID string) error
	AddItem(ctx context.Context, userID string, rewards []*RewqrdDTO) error
	GetItems(ctx context.Context, itemIDs []string) ([]*ItemDTO, error)
}
