package service

import (
	"github.com/go-errors/errors"
	"github.com/go-resty/resty"
	"net/http"

	"github.com/torinos-io/api/type/model"
	"fmt"
)

// Request is a request object for ...
type SaveRequest struct {
	AuthorizationCode string `json:"authorization_code"`
}

func (s *service) Save(req *SaveRequest) (*model.User, error) {

	// TODO: GitHubのやつでいいかんじにしていく

	token, err := s.exchangeCodeForAccessToken(req.AuthorizationCode)

	if err != nil {
		return nil, err
	}

	user, err := s.getGithubUser(token)

	return s.UserStore.CreateUserFromGithub(user, token)
}

func (s *service) exchangeCodeForAccessToken(code string) (string, error) {
	tokenResponse := struct {
		AccessToken string `json:"access_token"`
	}{}

	resp, err := resty.R().
		SetBody(map[string]string{
			"code":          code,
			"client_id":     s.Config.GithubClientID,
			"client_secret": s.Config.GithubClientSecret,
		}).
		SetResult(&tokenResponse).
		Get("https://github.com/login/oauth/access_token")

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	if code := resp.StatusCode(); code != http.StatusCreated {
		return nil, errors.Errorf("%d.auth_service.github", code)
	}

	return tokenResponse.AccessToken, nil
}

func (s *service) getGithubUser(accessToken string) (*model.GithubUser, error) {
	githubUser := model.GithubUser{}

	resp, err := resty.R().
		SetHeader("Authorization", fmt.Sprintf("token %s", accessToken)).
		SetResult(&githubUser).
		Get("https://api.github.com/user")

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	if code := resp.StatusCode(); code != http.StatusOK {
		return nil, errors.Errorf("%d.auth_service.github", code)
	}

	return &githubUser, nil
}
