package service

import (
	project_store "github.com/torinos-io/api/store/project_store"
	"github.com/torinos-io/api/type/model"
	"github.com/torinos-io/api/type/system"
)

// Context holds interfaces of external services
type Context struct {
	ProjectStore project_store.Store
	Config       *system.Config
}

// Service is an interface for authentication
type Service interface {
	Find(req *FindRequest) (*model.Project, error)
	FindAll(req *FindAllRequest) ([]*model.Project, error)
	Upload(req *UploadRequest) (*model.Project, error)
}

type service struct {
	Context
}

// New creates a new service instance from the context
func New(c Context) Service {
	return &service{
		Context: c,
	}
}
