package model

// User represents user
type User struct {
	ID         int    `json:"id"`
	GithubUUID string `json:"github_uuid"`

	UserName          string `json:"user_name"`
	GithubAccessToken string `json:"github_access_token"`
}
