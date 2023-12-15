package application

import "context"

type ItemUsecase interface {
	GetItemByID(ctx context.Context, id string) (*ItemDTO, error)
}
