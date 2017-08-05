package service

import (
	"github.com/go-errors/errors"
	"github.com/torinos-io/api/type/model"
	"unicode/utf8"
)

// FindRequest is a request object for finds project
type FindRequest struct {
	UUID string `json:"uuid"`
}

// Find returns project for given uuid
func (s *service) Find(req *FindRequest) (*model.Project, error) {
	if utf8.RuneCountInString(req.UUID) == 0 {
		return nil, errors.New("UUID is empty")
	}

	return s.ProjectStore.GetProjectByProjectUUID(req.UUID)
}
