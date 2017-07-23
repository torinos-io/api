package store

import (
	"bufio"
	"fmt"
	"mime/multipart"

	"github.com/go-errors/errors"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"

	project_service "github.com/torinos-io/api/service/project_service"
	"github.com/torinos-io/api/type/model"
)

type concreteStore struct {
	db *gorm.DB
}

// Store is an interface for CRUD category records
type Store interface {
	Upload(userID null.Int, req *project_service.UploadRequest) (*model.Project, error)
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
func (s *concreteStore) Upload(userID null.Int, req *project_service.UploadRequest) (*model.Project, error) {
	project := &model.Project{}

	project.UserID = userID

	if carFileContent, err := readFile(req.CartfileContent); err != nil {
		return nil, err
	} else {
		project.CartfileContent = carFileContent
	}

	if podFileLockContent, err := readFile(req.PodfileLockContent); err != nil {
		return nil, err
	} else {
		project.PodfileLockContent = podFileLockContent
	}

	if pbxProjectContent, err := readFile(req.PBXprojectContent); err != nil {
		return nil, err
	} else {
		project.PBXprojectContent = pbxProjectContent
	}

	db := s.db.Save(project)

	if err := db.Error; err != nil {
		return project, errors.Wrap(err, 0)
	}

	return project, nil
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

func readFile(fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader == nil {
		return nil, nil
	}

	file, err := fileHeader.Open()

	defer file.Close()

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	scanner := bufio.NewScanner(file)
	var content string

	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	fmt.Println(content)

	return content, nil
}
