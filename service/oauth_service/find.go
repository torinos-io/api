package service

import "github.com/torinos-io/api/type/model"

// FindRequest is a request object for finds user
type FindRequest struct {
	UUID string `json:"uuid"`
}

// Find returns the found user
func (s *service) Find(req *FindRequest) (*model.User, error) {
	return s.UserStore.FindByGithubUser(req.UUID)
}

func (s *service) FindByAuthToken(token string) (*model.User, error) {
	return s.UserStore.FindByGithubUser(token)
}
