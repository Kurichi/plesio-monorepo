package application

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
)

type missionUsecase struct {
	repo domain.MissionRepository
	svc  domain.MissionService
	pub  domain.Publisher
	tx   domain.TxRepository
}

func NewMissionUsecase(
	repo domain.MissionRepository,
	svc domain.MissionService,
	pub domain.Publisher,
	tx domain.TxRepository,
) MissionUsecase {
	return &missionUsecase{
		repo: repo,
		svc:  svc,
		pub:  pub,
		tx:   tx,
	}
}

// GetMyMissions implements MissionUsecase.
func (uc *missionUsecase) GetMissions(ctx context.Context, userID string, term *string) ([]*UserMissionDTO, error) {
	var filter domain.Filter
	if term != nil {
		if *term == "daily" {
			filter.Term = domain.TermDaily
		} else if *term == "weekly" {
			filter.Term = domain.TermDaily
		}
	}
	dailyMissions, err := uc.repo.GetUserMissions(ctx, userID, filter)
	if err != nil {
		return nil, err
	}

	if len(dailyMissions) > 0 {
		return NewUserMissionsFromEntity(dailyMissions), nil
	}

	missions, err := uc.svc.GenerateDailyUserMission(ctx, userID)

	if err != nil {
		return nil, err
	}

	return NewUserMissionsFromEntity(missions), nil
}

func (uc *missionUsecase) ProgressMission(ctx context.Context, userID string, missionID string, progress int) (*UserMissionDTO, error) {
	userMission := &domain.UserMission{}
	err := uc.tx.DoInTx(ctx, func(ctx context.Context) error {
		var err error
		userMission, err = uc.repo.GetUserMissionByID(ctx, userID, missionID)
		if err != nil {
			return err
		}
		userMission.UpdateProgress(progress)
		if err = uc.repo.UpdateUserMission(ctx, userID, userMission); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	isCompleted, event := userMission.IsCompleted()
	if !isCompleted {
		err := uc.pub.Publish(ctx, event)
		if err != nil {
			return nil, err
		}
	}

	return NewUserMissionFromEntity(userMission, isCompleted), nil
}
