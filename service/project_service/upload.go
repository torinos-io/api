package service

import (
	"mime/multipart"

	"github.com/guregu/null"

	"github.com/torinos-io/api/type/model"
)

// UploadRequest holds uploaded files
type UploadRequest struct {
	CartfileContent    *multipart.FileHeader `json:"cartfile_content"`
	PodfileLockContent *multipart.FileHeader `json:"podfile_content"`
	PbxprojContent     *multipart.FileHeader `json:"pbxproj_content"`
	RepositoryName     string                `json:"repository_name"`
}

// Upload uploads files to Analyze service
func (s *service) Upload(userID null.Int, req *UploadRequest) (*model.Project, error) {

	files := &model.UploadedData{
		CartfileContent:    req.CartfileContent,
		PodfileLockContent: req.PodfileLockContent,
		PbxprojContent:     req.PbxprojContent,
		RepositoryName:     req.RepositoryName,
	}

	return s.ProjectStore.Upload(userID, files)
}
