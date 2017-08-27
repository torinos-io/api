package service

import (
	"github.com/guregu/null"

	"github.com/torinos-io/api/type/model"
)

// FindAllRequest is a request object for finds all projects
type FindAllRequest struct {
	UserID null.Int `json:"-"`
}

// FindAll returns all projects related to given user id
func (s *service) FindAll(req *FindAllRequest) ([]*model.Project, error) {
	return s.ProjectStore.GetAllProjectsByUserID(req.UserID)
}
