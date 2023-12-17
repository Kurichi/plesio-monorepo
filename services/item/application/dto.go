package application

import "github.com/Kurichi/plesio-monorepo/services/item/domain"

type ItemDTO struct {
	ID          string
	Name        string
	Description string
}

type ItemWithQuantityDTO struct {
	ID          string
	Name        string
	Description string
	Quantity    int
}

type RewqrdDTO struct {
	ItemID string
	Amount int
}

type TreeDTO struct {
	Stage      int
	Water      int
	Fertilizer int
}

func NewItemFromEntity(item *domain.Item) *ItemDTO {
	return &ItemDTO{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
	}
}

func NewItemWithQuantityFromEntity(item *domain.ItemWithQuantity) *ItemWithQuantityDTO {
	return &ItemWithQuantityDTO{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
}
