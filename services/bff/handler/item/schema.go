package item

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Thumbnail   string `json:"thumbnail"`
}

type GetItemResponse struct {
	Item *Item `json:"item"`
}

type UseItemResponse struct {
	Item *Item `json:"item"`
}

func NewGetItemResponse() *GetItemResponse {
	return &GetItemResponse{
		Item: &Item{
			ID:          "1",
			Name:        "name",
			DisplayName: "displayName",
			Thumbnail:   "thumbnail",
		},
	}
}

func NewGetItemsResponse() []*GetItemResponse {
	items := make([]*GetItemResponse, 0, 1)

	items = append(items, &GetItemResponse{
		Item: &Item{
			ID:          "1",
			Name:        "name",
			DisplayName: "displayName",
			Thumbnail:   "thumbnail",
		},
	})

	return items
}

func NewUseItemResponse() *UseItemResponse {
	return &UseItemResponse{
		Item: &Item{
			ID:          "1",
			Name:        "name",
			DisplayName: "displayName",
			Thumbnail:   "thumbnail",
		},
	}
}
