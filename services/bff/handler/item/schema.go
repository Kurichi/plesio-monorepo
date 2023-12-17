package item

type CreateItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Target      string `json:"target"`
}

type CreateItemResponse struct {
	Item *NewItem `json:"item"`
}

type NewItem struct {
	CreateItemRequest
	ID string `json:"id"`
}

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quantity    uint32 `json:"quantity"`
}

type GetItemsResponse struct {
	Items []*Item `json:"items"`
}

type UseItemRequest struct {
	ItemID string `json:"itemID"`
}
