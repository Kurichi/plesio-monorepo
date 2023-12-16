package infra

import (
	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	"github.com/uptrace/bun"
)

type item struct {
	bun.BaseModel `bun:"items"`

	ID          string `bun:",pk"`
	Name        string
	Description string
	Target      string
	Amount      int
}

func (item *item) ConvertToEntity() *domain.Item {
	return &domain.Item{
		ID:          domain.NewItemIDFromString(item.ID),
		Name:        item.Name,
		Description: item.Description,
		Target:      domain.NewTargetFromString(item.Target),
		Amount:      item.Amount,
	}
}

type inventory struct {
	bun.BaseModel `bun:"inventories"`

	UserID   string `bun:",pk"`
	ItemID   string `bun:",pk"`
	Item     *item  `bun:"rel:belongs-to"`
	Quantity int
}

func (inventory *inventory) ConvertToEntity() *domain.ItemWithQuantity {
	return &domain.ItemWithQuantity{
		Item:     inventory.Item.ConvertToEntity(),
		Quantity: inventory.Quantity,
	}
}
