package infra

import (
	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
	"github.com/uptrace/bun"
)

type mission struct {
	bun.BaseModel `bun:"missions"`

	ID          string
	Description string
	Term        string `bun:",index"`
	Target      string
	Amount      int
	Unit        string
	Rewards     []*reward `bun:"rel:has-many,join:id=mission_id"`
}

type reward struct {
	bun.BaseModel `bun:"rewards"`

	MissionID string `bun:",pk"`
	ItemID    string `bun:",pk"`
	Amount    int
}

type userMission struct {
	bun.BaseModel `bun:"user_missions"`

	UserID    string
	MissionID string
	Mission   *mission `bun:"rel:belongs-to,join:mission_id=id"`
	Progress  int
	Deadline  int64
}

func (m *mission) ConvertToEntity() *domain.Mission {
	return &domain.Mission{
		ID:          m.ID,
		Description: m.Description,
		Term:        m.Term,
		Target:      m.Target,
		Amount:      m.Amount,
		Unit:        m.Unit,
		Rewards:     m.convertRewardsToEntity(),
	}
}

func (m *mission) convertRewardsToEntity() []*domain.Reward {
	rewards := make([]*domain.Reward, len(m.Rewards))
	for i, r := range m.Rewards {
		rewards[i] = r.ConvertToEntity()
	}
	return rewards
}

func (r *reward) ConvertToEntity() *domain.Reward {
	return &domain.Reward{
		ItemID: r.ItemID,
		Amount: r.Amount,
	}
}

func (um *userMission) ConvertToEntity() *domain.UserMission {
	return &domain.UserMission{
		UserID:   um.UserID,
		Mission:  um.Mission.ConvertToEntity(),
		Progress: um.Progress,
		Deadline: um.Deadline,
	}
}

