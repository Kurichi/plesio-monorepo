package application

import "context"

type MissionUsecase interface {
	GetMissions(ctx context.Context, userID string, Term *string) ([]*UserMissionDTO, error)
	ProgressMission(ctx context.Context, userID string, missionID string, progress int) (*UserMissionDTO, error)
}
