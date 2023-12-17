package domain

import (
	"time"
)

type CommitListResponse []struct {
	URL         string   `json:"url"`
	SHA         string   `json:"sha"`
	NodeID      string   `json:"node_id"`
	HtmlURL     string   `json:"html_url"`
	CommentsURL string   `json:"comments_url"`
	Commit      Commit   `json:"commit"`
	Author      User     `json:"author"`
	Committer   User     `json:"committer"`
	Parents     []Parent `json:"parents"`
	Stats       Stats    `json:"stats"`
	Files       []File   `json:"files"`
}

type Commit struct {
	URL          string       `json:"url"`
	Author       GitUser      `json:"author"`
	Committer    GitUser      `json:"committer"`
	Message      string       `json:"message"`
	CommentCount int          `json:"comment_count"`
	Tree         Tree         `json:"tree"`
	Verification Verification `json:"verification"`
}

type GitUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HtmlURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	StarredAt         string `json:"starred_at"`
}

type Parent struct {
	SHA     string `json:"sha"`
	URL     string `json:"url"`
	HtmlURL string `json:"html_url"`
}

type Stats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

type File struct {
	SHA              string `json:"sha"`
	Filename         string `json:"filename"`
	Status           string `json:"status"`
	Additions        int    `json:"additions"`
	Deletions        int    `json:"deletions"`
	Changes          int    `json:"changes"`
	BlobURL          string `json:"blob_url"`
	RawURL           string `json:"raw_url"`
	ContentsURL      string `json:"contents_url"`
	Patch            string `json:"patch"`
	PreviousFilename string `json:"previous_filename"`
}

type Tree struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type Verification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
}
