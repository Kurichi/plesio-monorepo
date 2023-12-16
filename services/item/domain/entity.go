package domain

import "github.com/pkg/errors"

type Item struct {
	ID          ItemID
	Name        string
	Description string
	Target      Target
	Amount      int
}

func NewItem(name, description string, target Target, amount int) *Item {
	return &Item{
		ID:          NewItemID(),
		Name:        name,
		Description: description,
		Target:      target,
		Amount:      amount,
	}
}

func (i *ItemWithQuantity) Use(userID string) *ItemUsedEvent {
	i.Quantity--
	return &ItemUsedEvent{
		UserID: userID,
		Target: i.Target,
		Amount: i.Amount,
	}
}

type ItemWithQuantity struct {
	*Item
	Quantity int
}

type Inventory struct {
	UserID string
	Items  []*ItemWithQuantity
}

func (i *Inventory) AddItem(item *ItemWithQuantity) {
	for _, v := range i.Items {
		if v.ID == item.ID {
			v.Quantity += item.Quantity
			return
		}
	}

	i.Items = append(i.Items, item)
}

func (i *Inventory) GetItem(itemID ItemID) (*ItemWithQuantity, error) {
	for _, v := range i.Items {
		if v.ID == itemID {
			return v, nil
		}
	}

	return nil, errors.New("item not found")
}

func (i *Inventory) RemoveItem(itemID ItemID) error {
	for idx, v := range i.Items {
		if v.ID == itemID {
			if v.Quantity == 1 {
				i.Items = append(i.Items[:idx], i.Items[idx+1:]...)
			} else {
				v.Quantity--
			}
			return nil
		}
	}

	return errors.New("item not found")
}

func (i *Inventory) Set(item *ItemWithQuantity) {
	for idx, v := range i.Items {
		if v.ID == item.ID {
			i.Items[idx] = item
			return
		}
	}

	i.Items = append(i.Items, item)
}
