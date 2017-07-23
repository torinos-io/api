package store

import (
	"bufio"
	"fmt"
	"mime/multipart"

	"github.com/go-errors/errors"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

	"github.com/torinos-io/api/type/model"
	"unicode/utf8"
)

type concreteStore struct {
	db *gorm.DB
}

// Store is an interface for CRUD category records
type Store interface {
	Upload(userID null.Int, files *model.UploadFiles) (*model.Project, error)
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
func (s *concreteStore) Upload(userID null.Int, files *model.UploadFiles) (*model.Project, error) {
	project := &model.Project{}
	finder := s.db.Where("repository = ?", files.RepositoryName).Find(project)

	if err := finder.Error; err != nil {
		return project, errors.Wrap(err, 0)
	}

	project.UserID = userID

	if carFileContent, err := readFile(files.CartfileContent); err == nil {
		project.CartfileContent = carFileContent
	}

	if podFileLockContent, err := readFile(files.PodfileLockContent); err == nil {
		project.PodfileLockContent = podFileLockContent
	}

	if pbxProjContent, err := readFile(files.PbxprojContent); err == nil {
		project.PbxprojContent = pbxProjContent
	}

	var db *gorm.DB

	if utf8.RuneCountInString(project.Repository) > 0 {
		db = s.db.Updates(project)
	} else {
		project.UUID = uuid.NewV4().String()
		db = s.db.Save(project)
	}

	if err := db.Error; err != nil {
		return project, errors.Wrap(err, 0)
	}

	return project, nil
}

// GetAllProjectsByUserID returns all projects
func (s *concreteStore) GetAllProjectsByUserID(userID null.Int) (*[]model.Project, error) {
	projects := &[]model.Project{}
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

func readFile(fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader == nil {
		return "", nil
	}

	file, err := fileHeader.Open()

	defer file.Close()

	if err != nil {
		return "", errors.Wrap(err, 0)
	}

	scanner := bufio.NewScanner(file)
	var content string

	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", errors.Wrap(err, 0)
	}

	fmt.Println(content)

	return content, nil
}
