package usecase

import "context"

type TreeUsecase interface {
	GetTree(ctx context.Context, id string) (*TreeDTO, error)
}
