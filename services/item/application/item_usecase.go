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
	tx       domain.TxRepository
	treeSvc  domain.TreeService
}

func NewItemUsecase(
	itemRepo domain.ItemRepository,
	invRepo domain.InventoryRepository,
	svc domain.ItemService,
	tx domain.TxRepository,
	treeSvc domain.TreeService,
) ItemUsecase {
	return &itemUsecase{
		itemRepo: itemRepo,
		invRepo:  invRepo,
		svc:      svc,
		tx:       tx,
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
func (uc *itemUsecase) UseItem(ctx context.Context, userID string, itemID string) (*TreeDTO, error) {
	var event *domain.ItemUsedEvent
	uc.tx.DoInTx(ctx, func(ctx context.Context) error {
		inv, err := uc.invRepo.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}

		event, err = uc.svc.UseItem(ctx, inv, domain.NewItemIDFromString(itemID))
		if err != nil {
			return err
		}

		return nil
	})

	tree, err := uc.treeSvc.GrowthEvent(ctx, event)
	if err != nil {
		return nil, err
	}

	dto := &TreeDTO{
		Stage:      tree.Stage,
		Water:      tree.Water,
		Fertilizer: tree.Fertilizer,
	}

	return dto, nil
}

// GetItem implements ItemUsecase.
func (uc *itemUsecase) AddItem(ctx context.Context, userID string, rewards []*RewqrdDTO) error {
	err := uc.tx.DoInTx(ctx, func(ctx context.Context) error {
		inv, err := uc.invRepo.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}

		for _, r := range rewards {
			item := &domain.ItemWithQuantity{
				Item:     &domain.Item{ID: domain.NewItemIDFromString(r.ItemID)},
				Quantity: r.Amount,
			}
			inv.AddItem(item)
		}

		if err := uc.invRepo.Update(ctx, inv); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
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
