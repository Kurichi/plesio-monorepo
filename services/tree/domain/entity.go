package domain

type Tree struct {
	ID   TreeID
	Name string
}

func NewTree(name string) (*Tree, error) {
	//TODO: validation

	return &Tree{
		ID:   NewTreeID(),
		Name: name,
	}, nil
}
