package model

// User represents user
type User struct {
	ID                int    `json:"id"`
	UserName          string `json:"user_name"`
	GitHubUUID        string `json:"github_uuid"`
	GitHubAccessToken string `json:"github_access_token"`
}
