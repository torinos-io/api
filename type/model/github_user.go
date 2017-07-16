package model

import "strconv"

// GithubUser is response object from github api
type GithubUser struct {
	ID       int    `json:"id"`
	UserName string `json:"login"`
	Email    string `json:"email"`
}

// UUID return string github user id
func (u *GithubUser) UUID() string {
	return strconv.Itoa(u.ID)
}
