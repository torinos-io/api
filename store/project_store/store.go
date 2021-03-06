package store

import (
	"github.com/go-errors/errors"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

	"github.com/torinos-io/api/type/model"
)

type concreteStore struct {
	db *gorm.DB
}

// Store is an interface for CRUD category records
type Store interface {
	Upload(userID null.Int, files *model.UploadedData) (*model.Project, error)
	GetAllProjectsByUserID(userID null.Int) ([]*model.Project, error)
	GetProjectByProjectUUID(uuid string) (*model.Project, error)
}

// New creates a store
func New(db *gorm.DB) Store {
	return &concreteStore{
		db: db,
	}
}

// Upload uploads files to Analyze service and save project
func (s *concreteStore) Upload(userID null.Int, data *model.UploadedData) (*model.Project, error) {
	project := &model.Project{}
	finder := s.db.Where("repository = ?", data.RepositoryName).First(project)

	err := finder.Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return project, errors.Wrap(err, 0)
	}

	if err == gorm.ErrRecordNotFound {
		project.UUID = uuid.NewV4().String()
		project.UserID = userID
	}

	project.PodfileLockContent = data.PodfileLockContent
	project.CartfileContent = data.CartfileContent

	db := s.db.Save(project)

	if err := db.Error; err != nil {
		return project, errors.Wrap(err, 0)
	}

	// TODO: Upload data to analyzer server

	return project, nil
}

// GetAllProjectsByUserID returns all projects
func (s *concreteStore) GetAllProjectsByUserID(userID null.Int) ([]*model.Project, error) {
	projects := []*model.Project{}
	finder := s.db.Where("user_id = ?", userID).Find(projects)

	if err := finder.Error; err != nil {
		return projects, errors.Wrap(err, 0)
	}

	return projects, nil
}

// GetProjectByProjectUUID returns project
func (s *concreteStore) GetProjectByProjectUUID(uuid string) (*model.Project, error) {
	project := &model.Project{}
	finder := s.db.Where("uuid = ?", uuid).Find(project)

	if err := finder.Error; err != nil {
		return project, errors.Wrap(err, 0)
	}

	return project, nil
}
