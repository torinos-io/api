package service

import (
	"fmt"
	"net/http"

	"github.com/go-errors/errors"
	"github.com/go-resty/resty"

	"github.com/torinos-io/api/type/model"
)

// FindRequest is a request object for finds user
type FindRequest struct {
	UUID string `json:"uuid"`
}

// Find returns the found user
func (s *service) Find(req *FindRequest) (*model.User, error) {
	return s.UserStore.FindByGithubUUID(req.UUID)
}

// FindByAccessToken returns user that found by given access token
func (s *service) FindByAccessToken(token string) (*model.User, error) {
	githubUser := &model.GithubUser{}

	resp, _ := resty.
		SetDebug(s.Config.Env == "development").
		R().
		SetHeader("Authorization", fmt.Sprintf("token %s", token)).
		SetResult(githubUser).
		Get("https://api.github.com/user")

	if status := resp.StatusCode(); status != http.StatusOK {
		return nil, errors.Errorf("%d.auth_service.github", status)
	}

	return s.UserStore.FindByGithubUUID(githubUser.UUID())
}
