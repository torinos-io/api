package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-errors/errors"

	"github.com/torinos-io/api/type/model"
)

// FindRequest is a request object for finds project
type FindRequest struct {
	UUID string `json:"uuid" valid:"uuidv4"`
}

// Find returns project for given uuid
func (s *service) Find(req *FindRequest) (*model.Project, error) {
	if validated, err := govalidator.ValidateStruct(req); !validated || err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return s.ProjectStore.GetProjectByProjectUUID(req.UUID)
}
