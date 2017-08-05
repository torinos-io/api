package service

import (
	"encoding/base64"
	"io"
	"unicode/utf8"

	"github.com/go-errors/errors"
	"github.com/guregu/null"

	"github.com/torinos-io/api/type/model"
	"io/ioutil"
)

// UploadRequest holds uploaded files
type UploadRequest struct {
	CartfileContent    io.Reader `json:"-"`
	PodfileLockContent io.Reader `json:"-"`
	RepositoryName     string    `json:"repository_name"`
}

// Upload uploads files to Analyze service
func (s *service) Upload(userID null.Int, req *UploadRequest) (*model.Project, error) {
	if req.CartfileContent == nil && req.PodfileLockContent == nil {
		return nil, errors.New("CartFile and PodFile are both empty")
	}

	data := &model.UploadedData{}
	if s, err := toBase64(req.PodfileLockContent); utf8.RuneCountInString(s) != 0 && err == nil {
		data.PodfileLockContent = s
	}
	if s, err := toBase64(req.CartfileContent); utf8.RuneCountInString(s) != 0 && err == nil {
		data.CartfileContent = s
	}
	data.RepositoryName = req.RepositoryName
	if isEmpty(data.CartfileContent) && isEmpty(data.PodfileLockContent) {
		return nil, errors.New("CartFile and PodFile are both empty")
	}
	if isEmpty(data.RepositoryName) {
		return nil, errors.New("Repository name is empty")
	}

	return s.ProjectStore.Upload(userID, data)
}

func toBase64(buffer io.Reader) (string, error) {
	b, err := ioutil.ReadAll(buffer)

	if err != nil {
		return "", errors.Wrap(err, 0)
	}

	encoded := base64.StdEncoding.EncodeToString(b)

	return encoded, nil
}

func isEmpty(s string) bool {
	return s == ""
}
