package model

import "mime/multipart"

// UploadFiles contains file uploaded files
type UploadFiles struct {
	CartfileContent    *multipart.FileHeader
	PodfileLockContent *multipart.FileHeader
	PbxprojContent     *multipart.FileHeader
	RepositoryName     string
}
