package store

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"

	"github.com/torinos-io/api/type/model"
)

type concreteStore struct {
	db *gorm.DB
}

// Store is an interface for CRUD category records
type Store interface {
	CreateUserFromGithub(githubUser *model.GithubUser, accessToken string) (*model.User, error)
	FindByGithubUser(uuid string) (*model.User, error)
}

// New creates a store
func New(db *gorm.DB) Store {
	return &concreteStore{
		db: db,
	}
}

func (s *concreteStore) CreateUserFromGithub(githubUser *model.GithubUser, accessToken string) (*model.User, error) {
	user := &model.User{}
	finder := s.db.Where("github_uuid = ?", githubUser.UUID()).Find(user)

	if err := finder.Error; err != nil && err != gorm.ErrRecordNotFound {
		return user, errors.Wrap(err, 0)
	}

	user.GithubAccessToken = accessToken

	db := s.db.Save(user)

	if err := db.Error; err != nil {
		return user, errors.Wrap(err, 0)
	}

	return user, nil
}

func (s *concreteStore) FindByGithubUser(uuid string) (*model.User, error) {
	user := &model.User{}
	finder := s.db.Where("github_uuid = ?", uuid).Find(user)

	if err := finder.Error; err != nil {
		return user, errors.Wrap(err, 0)
	}

	return user, nil
}
