package model

// User represents user
type User struct {
	ID         int    `json:"id"`
	GitHubUUID string `json:"github_uuid"`

	UserName          string `json:"user_name"`
	GitHubAccessToken string `json:"github_access_token"`
}
