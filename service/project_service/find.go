package service

import "github.com/torinos-io/api/type/model"

// FindRequest is a request object for finds project
type FindRequest struct {
	UUID string `json:"uuid"`
}

// Find returns project for given uuid
func (s *service) Find(req *FindRequest) (*model.Project, error) {
	return s.ProjectStore.GetProjectByProjectUUID(req.UUID)
}
