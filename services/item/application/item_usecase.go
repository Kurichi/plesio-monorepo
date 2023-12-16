package application

import (
	"context"
	"errors"

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

// CreateItem implements ItemUsecase.
func (uc *itemUsecase) CreateItem(ctx context.Context, name string, description string, target string, amount int) (*ItemDTO, error) {
	formattedTarget := domain.Target(target)
	if !formattedTarget.IsValid() {
		return nil, errors.New("target is invalid")
	}
	item := domain.NewItem(name, description, formattedTarget, amount)
	uc.itemRepo.StoreItem(ctx, item)
	return NewItemFromEntity(item), nil
}

// GetMyInventory implements ItemUsecase.
func (uc *itemUsecase) GetMyInventory(ctx context.Context, userID string) ([]*ItemWithQuantityDTO, error) {
	inv, err := uc.invRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	items := make([]*ItemWithQuantityDTO, 0, len(inv.Items))
	for _, i := range inv.Items {
		items = append(items, NewItemWithQuantityFromEntity(i))
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

// GetItems implements ItemUsecase.
func (uc *itemUsecase) GetItems(ctx context.Context, itemIDs []string) ([]*ItemDTO, error) {
	ids := make([]domain.ItemID, 0, len(itemIDs))
	for _, id := range itemIDs {
		ids = append(ids, domain.NewItemIDFromString(id))
	}
	items, err := uc.itemRepo.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	res := make([]*ItemDTO, 0, len(items))
	for _, i := range items {
		res = append(res, NewItemFromEntity(i))
	}

	return res, nil
}
