package store

import (
	"regexp"

	"github.com/go-errors/errors"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"

	"github.com/torinos-io/api/type/model"
)

type concreteStore struct {
	db *gorm.DB
}

// Store is an interface for CRUD category records
type Store interface {
	CreateSubscription(user *model.User, projectUUID string) (*model.Subscription, error)
	DeleteSubscription() error
}

// New creates a store
func New(db *gorm.DB) Store {
	return &concreteStore{
		db: db,
	}
}

func (s *concreteStore) CreateSubscription(user *model.User, projectUUID string) (*model.Subscription, error) {
	subscription := &model.Subscription{}

	finder := s.db.
		Where("user_id = ?", user.ID).
		Where("project_uuid", projectUUID).
		Find(subscription)

	if err := finder.Error; err != nil && err != gorm.ErrRecordNotFound {
		return subscription, errors.Wrap(err, 0)
	}

	emailPattern := regexp.MustCompile(`^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`)

	if !emailPattern.MatchString(user.Email) {
		return subscription, errors.New("Invalid email")
	}

	subscription.UserID = null.IntFrom(int64(user.ID))
	subscription.Email = user.Email
	subscription.DeletedAt = nil
	subscription.ProjectUUID = projectUUID

	db := s.db.Save(subscription)

	if err := db.Error; err != nil {
		return subscription, errors.Wrap(err, 0)
	}

	return subscription, nil
}

func (s *concreteStore) DeleteSubscription() error {
	return nil
}
