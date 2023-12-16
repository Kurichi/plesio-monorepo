package infra

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/Kurichi/plesio-monorepo/services/user/domain"
	"github.com/pkg/errors"
)

type GitHubUser struct {
	Login string `json:"login"`
}

type identityRepositoryImpl struct {
	client *auth.Client
}

func NewIdentityRepository(client *auth.Client) domain.IdentityRepository {
	return &identityRepositoryImpl{
		client: client,
	}
}

// FetchUser implements domain.IdentityRepository.
func (c *identityRepositoryImpl) FetchUser(ctx context.Context, uid string) (*domain.User, error) {
	user, err := c.client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	if len(user.ProviderUserInfo) == 0 {
		return nil, errors.New("provider user info is not found")
	}

	githubID, err := getUserByGithubID(user.ProviderUserInfo[0].UID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       user.UID,
		Name:     githubID.Login,
		Avatar:   user.PhotoURL,
		GithubID: user.ProviderUserInfo[0].UID,
	}, nil
}

func getUserByGithubID(githubID string) (*GitHubUser, error) {
	resp, err := http.Get("https://api.github.com/user/" + githubID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := &GitHubUser{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return user, nil
}
