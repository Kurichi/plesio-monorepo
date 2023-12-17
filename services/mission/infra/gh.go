package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
)

type gitHubRepository struct {
}

func NewGitHubRepository() domain.GitHubRepository {
	return &gitHubRepository{}
}

const url = "https://api.github.com/repos/Kurichi/plesio-monorepo/commits"

// GetCommits implements domain.GitHubRepository.
func (*gitHubRepository) GetCommits(ctx context.Context, token string) (*domain.CommitListResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer YOUR-TOKEN") // ここに実際のトークンを設定
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to request: %s\n", resp.Status)
		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var commits domain.CommitListResponse
	err = json.Unmarshal(body, &commits)
	if err != nil {
		panic(err)
	}

	// コミット情報の出力
	for _, commit := range commits {
		fmt.Printf("Commit SHA: %s, Author: %s, Message: %s\n",
			commit.SHA, commit.Commit.Author.Name, commit.Commit.Message)
	}

	return &commits, nil
}
