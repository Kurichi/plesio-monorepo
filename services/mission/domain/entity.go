package domain

type Mission struct {
	ID          string
	Description string
	Term        string
	Target      string
	Amount      int
	Unit        string
	Rewards     []*Reward
}

type UserMission struct {
	UserID   string
	Mission  *Mission
	Progress int
	Deadline int64
}

func (m *UserMission) IsCompleted() bool {
	return m.Progress >= m.Mission.Amount
}

func (m *UserMission) UpdateProgress(progress int) {
	m.Progress = progress
}
