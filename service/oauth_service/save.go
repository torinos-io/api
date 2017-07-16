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

	resp, err := resty.
		SetDebug(s.Config.Env == "development").
		R().
		SetBody(map[string]string{
			"code":          code,
			"client_id":     s.Config.GithubClientID,
			"client_secret": s.Config.GithubClientSecret,
		}).
		SetHeader("Accept", "application/json").
		SetResult(tokenResponse).
		Post("https://github.com/login/oauth/access_token")

	if resp.StatusCode() != http.StatusOK {
		return "", errors.Wrap(err, 0)
	}

	return tokenResponse.AccessToken, nil
}

func (s *service) getGithubUser(accessToken string) (*model.GithubUser, error) {
	githubUser := &model.GithubUser{}

	resp, err := resty.
		SetDebug(s.Config.Env == "development").
		R().
		SetHeader("Authorization", fmt.Sprintf("token %s", accessToken)).
		SetResult(githubUser).
		Get("https://api.github.com/user")

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.Wrap(err, 0)
	}

	return githubUser, nil
}
