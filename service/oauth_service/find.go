package service

import "github.com/torinos-io/api/type/model"

// FindRequest is a request object for ...
type FindRequest struct {
	UUID string `json:"uuid"`
}

// Find return the found user
func (s *service) Find(req *FindRequest) (*model.User, error) {
	return s.UserStore.FindByGithubUser(req.UUID)
}
