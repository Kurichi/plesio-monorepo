package domain

import (
	"context"
)

type TreeRepository interface {
	CreateTree(ctx context.Context, tree *Tree) error
	UpdateTree(ctx context.Context, tree *Tree) error
	GetTreeByUserID(ctx context.Context, userID string) (*Tree, error)
	GetTreesRanking(ctx context.Context, limit int) ([]*Tree, error)
}

type TxRepository interface {
	DoInTx(ctx context.Context, fn func(ctx context.Context) error) error
}
