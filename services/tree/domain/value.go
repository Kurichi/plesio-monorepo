package domain

import "github.com/google/uuid"

type TreeID string

func NewTreeID() TreeID {
	id := uuid.New()
	return TreeID(id.String())
}
func NewTreeIDFromString(s string) TreeID {
	return TreeID(s)
}
func (id TreeID) String() string {
	return string(id)
}
