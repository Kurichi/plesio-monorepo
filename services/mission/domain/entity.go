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

func (m *UserMission) IsCompleted() (bool, *CompleteMissionEvent) {
	if m.Progress >= m.Mission.Amount {
		return true, &CompleteMissionEvent{
			UserID:  m.UserID,
			Rewards: m.Mission.Rewards,
		}
	}

	return false, nil
}

func (m *UserMission) UpdateProgress(progress int) {
	m.Progress = progress
}

func (m *UserMission) CheckProgress(cont *CommitListResponse) bool {
	if m.Mission.Target == "commit" && m.Mission.Amount <= len(*cont) {
		return true
	}

	return false
}
