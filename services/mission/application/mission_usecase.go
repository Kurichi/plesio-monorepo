package application

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
	"github.com/Kurichi/plesio-monorepo/services/mission/infra"
	"github.com/google/uuid"
)

type missionUsecase struct {
	repo domain.MissionRepository
	svc  domain.MissionService
	pub  domain.Publisher
	gh   domain.GitHubRepository
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
		gh:   infra.NewGitHubRepository(),
		tx:   tx,
	}
}

// CreateMission implements MissionUsecase.
func (uc *missionUsecase) CreateMission(ctx context.Context, description string, target string, amount int, unit string, term string, rewards []*RewardDTO) (*MissionDTO, error) {
	missionDTO := MissionDTO{
		ID:          uuid.New().String(),
		Description: description,
		Term:        term,
		Target:      target,
		Amount:      amount,
		Unit:        unit,
		Rewards:     rewards,
	}
	err := uc.repo.StoreMission(ctx, []*domain.Mission{
		missionDTO.ConvertToEntity(),
	})
	if err != nil {
		return nil, err
	}
	return &missionDTO, nil
}

// GetMyMissions implements MissionUsecase.
func (uc *missionUsecase) GetMissions(ctx context.Context, userID string, term *string, token string) ([]*UserMissionDTO, error) {
	cons, err := uc.gh.GetCommits(ctx, token)
	if err != nil {
		return nil, err
	}

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

	ret := make([]*UserMissionDTO, 0, len(dailyMissions))
	for _, mission := range dailyMissions {
		var isCompleted bool
		if isCompleted := mission.CheckProgress(cons); isCompleted {
			mission.UpdateProgress(1)
			if err := uc.repo.UpdateUserMission(ctx, userID, mission); err != nil {
				return nil, err
			}
		}
		ret = append(ret, NewUserMissionFromEntity(mission, isCompleted))
	}

	if len(ret) > 0 {
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
