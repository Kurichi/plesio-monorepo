package application

import "context"

type MissionUsecase interface {
	GetMyDailyMissions(ctx context.Context, userID string) ([]*UserMissionDTO, error)
}
