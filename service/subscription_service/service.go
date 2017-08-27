package service

import (
	subscription_store "github.com/torinos-io/api/store/subscription_store"
	"github.com/torinos-io/api/type/model"
	"github.com/torinos-io/api/type/system"
)

// Context holds interfaces of external services
type Context struct {
	SubscriptionStore subscription_store.Store
	Config            *system.Config
}

// Service is an interface for authentication
type Service interface {
	Subscribe(req *SubscribeRequest) (*model.Subscription, error)
	UnSubscribe(req *UnSubscribeRequest) error
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
