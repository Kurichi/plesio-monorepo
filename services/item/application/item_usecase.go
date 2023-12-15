package application

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/domain"
)

type itemUsecase struct {
	itemRepo domain.ItemRepository
}

func NewItemUsecase(itemRepo domain.ItemRepository) ItemUsecase {
	return &itemUsecase{itemRepo: itemRepo}
}

// GetItemByID implements ItemUsecase.
func (uc *itemUsecase) GetItemByID(ctx context.Context, id string) (*ItemDTO, error) {
	item, err := uc.itemRepo.GetByID(ctx, domain.NewItemIDFromString(id))
	if err != nil {
		return nil, err
	}

	return NewItemFromEntity(item), nil
}
