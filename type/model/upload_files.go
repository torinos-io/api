package model

import "mime/multipart"

// UploadedData contains file uploaded files
type UploadedData struct {
	CartfileContent    *multipart.FileHeader
	PodfileLockContent *multipart.FileHeader
	PbxprojContent     *multipart.FileHeader
	RepositoryName     string
}
