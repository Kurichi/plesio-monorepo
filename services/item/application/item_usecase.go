package application

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/domain"
)

type itemUsecase struct {
	itemRepo domain.ItemRepository
	invRepo  domain.InventoryRepository
	svc      domain.ItemService
}

func NewItemUsecase(
	itemRepo domain.ItemRepository,
	invRepo domain.InventoryRepository,
	svc domain.ItemService,
) ItemUsecase {
	return &itemUsecase{
		itemRepo: itemRepo,
		invRepo:  invRepo,
		svc:      svc,
	}
}

// GetMyInventory implements ItemUsecase.
func (uc *itemUsecase) GetMyInventory(ctx context.Context, userID string) ([]*ItemDTO, error) {
	inv, err := uc.invRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	items := make([]*ItemDTO, 0, len(inv.Items))
	for _, i := range inv.Items {
		items = append(items, NewItemFromEntity(i))
	}

	return items, nil
}

// UseItem implements ItemUsecase.
func (uc *itemUsecase) UseItem(ctx context.Context, userID string, itemID string) error {
	inv, err := uc.invRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if err := uc.svc.UseItem(ctx, inv, domain.NewItemIDFromString(itemID)); err != nil {
		return err
	}
	return nil
}

// GetItem implements ItemUsecase.
func (uc *itemUsecase) GetItem(ctx context.Context, userID string, itemID string) (*ItemDTO, error) {
	panic("unimplemented")
}
