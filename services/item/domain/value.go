package domain

import "github.com/google/uuid"

type (
	ItemID string
	Target string
)

const (
	TargetWater      Target = "water"
	TargetFertilizer Target = "fertilizer"
)

func NewItemID() ItemID {
	id := uuid.New()
	return ItemID(id.String())
}

func NewItemIDFromString(id string) ItemID {
	return ItemID(id)
}

func (id ItemID) String() string {
	return string(id)
}

func NewTargetFromString(target string) Target {
	return Target(target)
}

func (t Target) String() string {
	return string(t)
}

func (t Target) IsValid() bool {
	return t == TargetWater || t == TargetFertilizer
}

type ItemUsedEvent struct {
	UserID string
	Target Target
	Amount int
}

type Tree struct {
	Stage      int
	Water      int
	Fertilizer int
}
