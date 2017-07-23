package model

// User represents user
type User struct {
	ID         int    `json:"id" gorm:"primary_key"`
	GithubUUID string `json:"github_uuid"`

	UserName          string `json:"user_name"`
	Email             string `json:"email"`
	GithubAccessToken string `json:"github_access_token"`
}
