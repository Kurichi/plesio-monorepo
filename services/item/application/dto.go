package application

import "github.com/Kurichi/plesio-monorepo/services/item/domain"

type ItemDTO struct {
	ID          string
	Name        string
	Description string
}

func NewItemFromEntity(item *domain.Item) *ItemDTO {
	return &ItemDTO{
		ID:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
	}
}
