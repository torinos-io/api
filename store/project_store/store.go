package store

import (
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
	Upload() error
	GetAllProjectsByUserID(userID null.Int) (*[]model.Project, error)
	GetProjectByProjectUUID(uuid string) (*model.Project, error)
}

// New creates a store
func New(db *gorm.DB) Store {
	return &concreteStore{
		db: db,
	}
}

// Upload uploads files to Analyze service
func (s *concreteStore) Upload() error {
	// TODO: Upload to Analyze server
	return nil
}

// GetAllProjectsByUserID returns all projects
func (s *concreteStore) GetAllProjectsByUserID(userID null.Int) (*[]model.Project, error) {
	projects := &[]model.Project{}
	finder := s.db.Where("user_id", userID).Find(projects)

	if err := finder.Error; err != nil {
		return projects, errors.Wrap(err, 0)
	}

	return projects, nil
}

// GetProjectByProjectUUID returns project
func (s *concreteStore) GetProjectByProjectUUID(uuid string) (*model.Project, error) {
	project := &model.Project{}

	finder := s.db.Where("uuid", uuid).Find(project)

	if err := finder.Error; err != nil {
		return project, errors.Wrap(err, 0)
	}

	return project, nil
}
