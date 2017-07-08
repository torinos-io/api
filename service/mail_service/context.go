package service

// Context holds interfaces of external services
type Context struct {
}

// Service is an interface for ...
type Service interface {
	Perform(req *Request) error
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
