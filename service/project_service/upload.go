package service

import (
	"mime/multipart"

	"github.com/torinos-io/api/type/model"
	"github.com/guregu/null"
)

// UploadRequest holds uploaded files
type UploadRequest struct {
	CartfileContent    *multipart.FileHeader `json:"cartfile_content"`
	PodfileLockContent *multipart.FileHeader `json:"podfile_content"`
	PBXprojectContent  *multipart.FileHeader `json:"pbxproj_content"`
}

// Upload uploads files to Analyze service
func (s *service) Upload(userID null.Int, req *UploadRequest) (*model.Project, error) {
	return s.ProjectStore.Upload(userID, req)
}
