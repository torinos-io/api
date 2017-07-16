package model

import "strconv"

type GithubUser struct {
	ID       int    `json:"id"`
	UserName string `json:"login"`
	Email    string `json:"email"`
}

func (u *GithubUser) UUID() string {
	return strconv.Itoa(u.ID)
}
