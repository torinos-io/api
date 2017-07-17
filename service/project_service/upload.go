package service

import (
	"mime/multipart"
)

// UploadRequest holds uploaded files
type UploadRequest struct {
	CartfileContent    *multipart.FileHeader `json:"cartfile_content"`
	PodfileLockContent *multipart.FileHeader `json:"podfile_content"`
	PBXprojectContent  *multipart.FileHeader `json:"pbxporj_content"`
}

// Upload uploads files to Analyze service
func (s *service) Upload(req *UploadRequest) error {
	// TODO: ここでFileをtext化する
	return s.ProjectStore.Upload()
}
