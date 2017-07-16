package service

import "github.com/torinos-io/api/type/model"

// Save saves new project
func (s *service) Save() (*model.Project, error) {
	return s.ProjectStore.CreateOrUpdateProject()
}
