package domain

import (
	"context"
)

type Filter struct {
	Term Term
}

type MissionRepository interface {
	// Mission
	GetRandomMissions(ctx context.Context, count uint8, filter Filter) ([]*Mission, error)
	GetMissions(ctx context.Context, filter Filter) ([]*Mission, error)
	StoreMission(ctx context.Context, missions []*Mission) error
	// UserMission
	GetUserMissions(ctx context.Context, userID string, filter Filter) ([]*UserMission, error)
	StoreUserMissions(ctx context.Context, userID string, missions []*UserMission) error
	UpdateUserMission(ctx context.Context, userID string, mission *UserMission) error
}
