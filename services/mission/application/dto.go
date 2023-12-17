package application

import "github.com/Kurichi/plesio-monorepo/services/mission/domain"

type (
	MissionDTO struct {
		ID          string
		Description string
		Term        string
		Target      string
		Amount      int
		Unit        string
		Rewards     []*RewardDTO
	}

	RewardDTO struct {
		ItemID string
		Amount int
	}

	UserMissionDTO struct {
		*MissionDTO
		Progress    int
		Deadline    int64
		IsCompleted bool
	}
)

func (r *RewardDTO) ConvertToEntity() *domain.Reward {
	return &domain.Reward{
		ItemID: r.ItemID,
		Amount: r.Amount,
	}
}

func ConvertRewardsDTOToEntities(rs []*RewardDTO) []*domain.Reward {
	rewards := make([]*domain.Reward, 0, len(rs))
	for _, r := range rs {
		rewards = append(rewards, r.ConvertToEntity())
	}
	return rewards
}

func (r *MissionDTO) ConvertToEntity() *domain.Mission {
	return &domain.Mission{
		ID:          r.ID,
		Description: r.Description,
		Term:        r.Term,
		Target:      r.Target,
		Amount:      r.Amount,
		Unit:        r.Unit,
		Rewards:     ConvertRewardsDTOToEntities(r.Rewards),
	}
}

func NewMissionFromEntity(e *domain.Mission) *MissionDTO {
	return &MissionDTO{
		ID:          e.ID,
		Description: e.Description,
		Term:        e.Term,
		Target:      e.Target,
		Amount:      e.Amount,
		Unit:        e.Unit,
		Rewards:     NewRewardsFromEntity(e.Rewards),
	}
}

func NewRewardsFromEntity(e []*domain.Reward) []*RewardDTO {
	rewards := make([]*RewardDTO, len(e))
	for i, r := range e {
		rewards[i] = NewRewardFromEntity(r)
	}
	return rewards
}

func NewRewardFromEntity(e *domain.Reward) *RewardDTO {
	return &RewardDTO{
		ItemID: e.ItemID,
		Amount: e.Amount,
	}
}

func NewUserMissionFromEntity(e *domain.UserMission, isCompleted bool) *UserMissionDTO {
	return &UserMissionDTO{
		MissionDTO:  NewMissionFromEntity(e.Mission),
		Progress:    e.Progress,
		Deadline:    e.Deadline,
		IsCompleted: isCompleted,
	}
}

func NewUserMissionsFromEntity(e []*domain.UserMission) []*UserMissionDTO {
	userMissions := make([]*UserMissionDTO, 0, len(e))

	for _, item := range e {
		isCompleted, _ := item.IsCompleted()
		userMissions = append(userMissions, NewUserMissionFromEntity(item, isCompleted))
	}

	return userMissions
}
