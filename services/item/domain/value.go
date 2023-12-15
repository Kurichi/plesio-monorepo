package domain

import "github.com/google/uuid"

type (
	ItemID string
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
