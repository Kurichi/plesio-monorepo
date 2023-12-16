package domain

import "context"

type ItemService interface {
	UseItem(ctx context.Context, inv *Inventory, itemID ItemID) error
}

type itemService struct {
	repo InventoryRepository
}

func NewItemService(repo InventoryRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (svc *itemService) UseItem(ctx context.Context, inv *Inventory, itemID ItemID) error {
	item, err := inv.GetItem(itemID)
	if err != nil {
		return err
	}
	// event := item.Use(inv.UserID)
	item.Use(inv.UserID)
	inv.Set(item)

	// Quantityが0になったら削除する
	if item.Quantity == 0 {
		if err := svc.repo.DeleteOneItem(ctx, inv.UserID, item.ID.String()); err != nil {
			return err
		}
		return nil
	}

	if err := svc.repo.UpdateOneItem(ctx, inv.UserID, item); err != nil {
		return err
	}

	return nil
}
