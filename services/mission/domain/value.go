package domain

type (
	Reward struct {
		ItemID string
		Amount int
	}

	Term uint
)

const (
	TermDaily Term = iota + 1
	TermWeekly
)

func (t Term) IsValid() bool {
	return t == TermDaily || t == TermWeekly
}

func (t Term) Int() int {
	return int(t)
}
