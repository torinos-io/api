package route

import (
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/guregu/null"

	"github.com/torinos-io/api/server/middleware"
	project_service "github.com/torinos-io/api/service/project_service"
	project_store "github.com/torinos-io/api/store/project_store"
	"github.com/torinos-io/api/type/model"
)

// CreateProject creates project
func CreateProject(c *gin.Context) {
	ac := middleware.GetAppContext(c)

	projectStore := project_store.New(ac.MainDB)
	service := project_service.New(project_service.Context{
		Config:       ac.Config,
		ProjectStore: projectStore,
	})

	cartfile, cartErr := c.FormFile("cartfile_content")
	podfile, podErr := c.FormFile("podfile_content")
	pbxproj, pbxprojErr := c.FormFile("pbxproj_content")

	if pbxprojErr != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.Wrap(pbxprojErr, 0))
		return
	}

	if cartErr != nil && podErr != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New("Podfile and Cartfile is empty"))
		return
	}

	uploadRequest := &project_service.UploadRequest{
		CartfileContent:    cartfile,
		PodfileLockContent: podfile,
		PbxprojContent:     pbxproj,
	}

	var project *model.Project
	var err error

	userID := null.Int{}

	if user := middleware.GetCurrentUser(c); user != nil {
		userID.Int64 = int64(user.ID)
		userID.Valid = true
	}

	service.Upload(userID, uploadRequest)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, project)
}

// GetProject returns the project
func GetProject(c *gin.Context) {
	uuid := c.Param("uuid")
	if utf8.RuneCountInString(uuid) == 0 {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	ac := middleware.GetAppContext(c)

	projectStore := project_store.New(ac.MainDB)
	service := project_service.New(project_service.Context{
		Config:       ac.Config,
		ProjectStore: projectStore,
	})

	request := &project_service.FindRequest{
		UUID: uuid,
	}

	project, err := service.Find(request)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, project)
}

// ListProjects returns all projects
func ListProjects(c *gin.Context) {
	user := middleware.GetCurrentUser(c)

	if user == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ac := middleware.GetAppContext(c)
	projectStore := project_store.New(ac.MainDB)
	service := project_service.New(project_service.Context{
		Config:       ac.Config,
		ProjectStore: projectStore,
	})

	request := &project_service.FindAllRequest{
		UserID: null.IntFrom(int64(user.ID)),
	}

	projects, err := service.FindAll(request)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, projects)
}
