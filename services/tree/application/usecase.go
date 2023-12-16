package application

import "context"

type TreeUsecase interface {
	InitTree(ctx context.Context, userID string) (*TreeDTO, error)
	PlantTree(ctx context.Context, userID string) (*TreeDTO, error)
	GetMyTree(ctx context.Context, userID string) (*TreeDTO, error)
	GetTreeByUserID(ctx context.Context, userID string) (*TreeDTO, error)
	GetTreeRanking(ctx context.Context, limit int) ([]*TreeDTO, error)
}
