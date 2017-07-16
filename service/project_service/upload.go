package service

import (
	"mime/multipart"
)

type UploadRequest struct {
	ProjectFiles []*multipart.FileHeader
}

// Upload uploads files to Analyze service
func (s *service) Upload() error {
	return s.ProjectStore.Upload()
}
