package service

import (
	"mime/multipart"
)

// UploadRequest holds uploaded files
type UploadRequest struct {
	CartfileContent    *multipart.FileHeader `json:"cartfile_content"`
	PodfileLockContent *multipart.FileHeader `json:"podfile_content"`
	PBXprojectContent  *multipart.FileHeader `json:"pbxproj_content"`
}

// Upload uploads files to Analyze service
func (s *service) Upload(req *UploadRequest) error {
	// TODO: Make files to texts
	return s.ProjectStore.Upload()
}
