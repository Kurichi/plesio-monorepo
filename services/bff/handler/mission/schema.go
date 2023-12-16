package mission

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Mission struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Items       []*Item `json:"items"`
}

type GetMissionsResponse struct {
	Missions []*Mission `json:"missions"`
}
