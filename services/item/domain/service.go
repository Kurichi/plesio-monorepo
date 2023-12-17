package domain

import "context"

type ItemService interface {
	UseItem(ctx context.Context, inv *Inventory, itemID ItemID) (*ItemUsedEvent, error)
}

type itemService struct {
	repo InventoryRepository
}

func NewItemService(repo InventoryRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (svc *itemService) UseItem(ctx context.Context, inv *Inventory, itemID ItemID) (*ItemUsedEvent, error) {
	item, err := inv.GetItem(itemID)
	if err != nil {
		return nil, err
	}
	event := item.Use(inv.UserID)
	item.Use(inv.UserID)
	inv.Set(item)

	// Quantityが0になったら削除する
	if item.Quantity == 0 {
		if err := svc.repo.DeleteOneItem(ctx, inv.UserID, item.ID.String()); err != nil {
			return nil, err
		}
		return event, nil
	}

	if err := svc.repo.UpdateOneItem(ctx, inv.UserID, item); err != nil {
		return nil, err
	}

	return event, nil
}
