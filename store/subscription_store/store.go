package store

import (
	"regexp"
	"time"

	"github.com/go-errors/errors"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"

	"github.com/torinos-io/api/type/model"
)

var emailPattern = regexp.MustCompile(`^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`)

type concreteStore struct {
	db *gorm.DB
}

// Store is an interface for CRUD category records
type Store interface {
	CreateSubscription(userID int, email string, projectUUID string) (*model.Subscription, error)
	DeleteSubscription(user *model.User, projectUUID string) error
}

// New creates a store
func New(db *gorm.DB) Store {
	return &concreteStore{
		db: db,
	}
}

func (s *concreteStore) CreateSubscription(userID int, email string, projectUUID string) (*model.Subscription, error) {
	subscription := &model.Subscription{}

	finder := s.db.
		Where("user_id = ?", userID).
		Where("project_uuid", projectUUID).
		Find(subscription)

	if err := finder.Error; err != nil && err != gorm.ErrRecordNotFound {
		return subscription, errors.Wrap(err, 0)
	}

	if !emailPattern.MatchString(email) {
		return subscription, errors.New("Invalid email")
	}

	subscription.UserID = null.IntFrom(int64(userID))
	subscription.Email = email
	subscription.DeletedAt = null.Time{}
	subscription.ProjectUUID = projectUUID

	db := s.db.Save(subscription)

	if err := db.Error; err != nil {
		return subscription, errors.Wrap(err, 0)
	}

	return subscription, nil
}

func (s *concreteStore) DeleteSubscription(user *model.User, projectUUID string) error {
	subscription := &model.Subscription{}

	finder := s.db.
		Where("user_id = ?", user.ID).
		Where("project_uuid", projectUUID).
		Find(subscription)

	if err := finder.Error; err != nil && err != gorm.ErrRecordNotFound {
		return errors.Wrap(err, 0)
	}

	subscription.DeletedAt = null.TimeFrom(time.Now())
	db := s.db.Save(subscription)

	if err := db.Error; err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
