package domain

type User struct {
	ID       string
	Name     string
	Avatar   string
	GithubID string
}

func NewUser(id, name, avatar, githubID string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Avatar:   avatar,
		GithubID: githubID,
	}
}
