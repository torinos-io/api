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
	CreateSubscription(userID null.Int, email string, projectUUID string) (*model.Subscription, error)
	DeleteSubscription(userID null.Int, projectUUID string) error
}

// New creates a store
func New(db *gorm.DB) Store {
	return &concreteStore{
		db: db,
	}
}

func (s *concreteStore) CreateSubscription(userID null.Int, email string, projectUUID string) (*model.Subscription, error) {
	subscription := &model.Subscription{}

	finder := s.db.
		Where("user_id = ?", userID.Int64).
		Where("project_uuid", projectUUID).
		Find(subscription)

	if err := finder.Error; err != nil && err != gorm.ErrRecordNotFound {
		return subscription, errors.Wrap(err, 0)
	}

	if !emailPattern.MatchString(email) {
		return subscription, errors.New("Invalid email")
	}

	subscription.UserID = userID
	subscription.Email = email
	subscription.DeletedAt = null.Time{}
	subscription.ProjectUUID = projectUUID

	db := s.db.Save(subscription)

	if err := db.Error; err != nil {
		return subscription, errors.Wrap(err, 0)
	}

	return subscription, nil
}

func (s *concreteStore) DeleteSubscription(userID null.Int, projectUUID string) error {
	subscription := &model.Subscription{}

	finder := s.db.
		Where("user_id = ?", userID.Int64).
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
