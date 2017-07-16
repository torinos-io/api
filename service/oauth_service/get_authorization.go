package service

import "fmt"

// Request is a request object for ...
type GetAuthorizationResponse struct {
	URL string `json:"url"`
}

func (s *service) GetAuthorization() (*GetAuthorizationResponse)  {
	response := GetAuthorizationResponse{}
	// https://github.com/login/oauth/authorize?scope=user:email&client_id=
	url := fmt.Sprintf("https://github.com/login/oauth/authorize?scope=user:email&client_id=%s", s.Config.GithubClientID)
	response.URL = url
	return &response
}
