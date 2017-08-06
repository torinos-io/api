package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-errors/errors"

	"github.com/torinos-io/api/type/model"
)

// FindRequest is a request object for finds project
type FindRequest struct {
	UUID string `json:"uuid" valid:"uuid"`
}

// Find returns project for given uuid
func (s *service) Find(req *FindRequest) (*model.Project, error) {
	if !govalidator.IsUUID(req.UUID) {
		return nil, errors.New("UUID is empty")
	}

	return s.ProjectStore.GetProjectByProjectUUID(req.UUID)
}
