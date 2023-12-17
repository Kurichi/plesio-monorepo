package mission

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Reward struct {
	ItemID string `json:"item_id"`
	Amount int    `json:"amount"`
}

type CreateMission struct {
	Description string    `json:"description"`
	Target      string    `json:"target"`
	Amount      int       `json:"amount"`
	Unit        string    `json:"unit"`
	Term        string    `json:"term"`
	Rewards     []*Reward `json:"rewards"`
}

type Mission struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Items       []*Item `json:"items"`
}

type DetailedMission struct {
	CreateMission
	ID string `json:"id"`
}

type CreateMissionResponse struct {
	Mission *DetailedMission `json:"mission"`
}

type GetMissionsResponse struct {
	Missions []*Mission `json:"missions"`
}
