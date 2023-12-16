package application

import "context"

type ItemUsecase interface {
	GetMyInventory(ctx context.Context, userID string) ([]*ItemDTO, error)
	UseItem(ctx context.Context, userID, itemID string) error
	GetItem(ctx context.Context, userID, itemID string) (*ItemDTO, error)
}
