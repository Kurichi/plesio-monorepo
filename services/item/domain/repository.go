package domain

import "context"

type ItemRepository interface {
	GetByID(ctx context.Context, id ItemID) (*Item, error)
}
