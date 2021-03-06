package service

import (
	user_store "github.com/torinos-io/api/store/user_store"
	"github.com/torinos-io/api/type/model"
	"github.com/torinos-io/api/type/system"
)

// Context holds interfaces of external services
type Context struct {
	UserStore user_store.Store
	Config    *system.Config
}

// Service is an interface for authentication
type Service interface {
	Find(req *FindRequest) (*model.User, error)
	Save(req *SaveRequest) (*model.User, error)
	GetAuthorization() *GetAuthorizationResponse
	FindByAccessToken(token string) (*model.User, error)
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
