package domain

import "context"

type ItemRepository interface {
	GetByID(ctx context.Context, id ItemID) (*Item, error)
	GetByIDs(ctx context.Context, id []ItemID) ([]*Item, error)
	StoreItem(ctx context.Context, item *Item) error
}

type InventoryRepository interface {
	GetByUserID(ctx context.Context, userID string) (*Inventory, error)
	Update(ctx context.Context, inventory *Inventory) error
	UpdateOneItem(ctx context.Context, userID string, item *ItemWithQuantity) error
	DeleteOneItem(ctx context.Context, userID string, itemID string) error
}

type TxRepository interface {
	DoInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type TreeService interface {
	GrowthEvent(ctx context.Context, event *ItemUsedEvent) (*Tree, error)
}
