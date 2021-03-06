package service

import (
	neturl "net/url"
)

// GetAuthorizationResponse is a response object for returns url
type GetAuthorizationResponse struct {
	URL string `json:"url"`
}

const (
	scope    = "repo,read:org,user"
	endpoint = "https://github.com/login/oauth/authorize"
)

// GetAuthorization returns the url for oauth
func (s *service) GetAuthorization() *GetAuthorizationResponse {
	response := GetAuthorizationResponse{}

	url, _ := neturl.Parse(endpoint)

	query := neturl.Values{}
	query.Add("access_type", "online")
	query.Add("scope", scope)
	query.Add("client_id", s.Config.GithubClientID)

	url.RawQuery = query.Encode()

	response.URL = url.String()

	return &response
}
