package service

import (
	"fmt"
	"net/http"

	"github.com/go-errors/errors"
	"github.com/go-resty/resty"

	"github.com/torinos-io/api/type/model"
)

// SaveRequest is a request object for returns authorization code
type SaveRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

// Save stores the user
func (s *service) Save(req *SaveRequest) (*model.User, error) {

	token, err := s.exchangeCodeForAccessToken(req.AuthorizationCode)

	if err != nil {
		return nil, err
	}

	user, err := s.getGithubUser(token)

	if err != nil {
		return nil, err
	}

	return s.UserStore.CreateUserFromGithub(user, token)
}

func (s *service) exchangeCodeForAccessToken(code string) (string, error) {
	tokenResponse := &struct {
		AccessToken string `json:"access_token"`
	}{}

	resp, _ := resty.
		SetDebug(s.Config.IsDevelopment()).
		R().
		SetBody(map[string]string{
			"code":          code,
			"client_id":     s.Config.GithubClientID,
			"client_secret": s.Config.GithubClientSecret,
		}).
		SetHeader("Accept", "application/json").
		SetResult(tokenResponse).
		Post("https://github.com/login/oauth/access_token")

	if status := resp.StatusCode(); status != http.StatusOK {
		return "", errors.Errorf("%d.auth_service.github", status)
	}

	return tokenResponse.AccessToken, nil
}

func (s *service) getGithubUser(accessToken string) (*model.GithubUser, error) {
	githubUser := &model.GithubUser{}

	resp, _ := resty.
		SetDebug(s.Config.IsDevelopment()).
		R().
		SetHeader("Authorization", fmt.Sprintf("token %s", accessToken)).
		SetResult(githubUser).
		Get("https://api.github.com/user")

	if status := resp.StatusCode(); status != http.StatusOK {
		return nil, errors.Errorf("%d.auth_service.github", status)
	}

	return githubUser, nil
}
