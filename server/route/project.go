package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/guregu/null"

	"github.com/torinos-io/api/server/middleware"
	project_service "github.com/torinos-io/api/service/project_service"
	project_store "github.com/torinos-io/api/store/project_store"
)

// CreateProject creates project
func CreateProject(c *gin.Context) {
	ac := middleware.GetAppContext(c)

	projectStore := project_store.New(ac.MainDB)
	service := project_service.New(project_service.Context{
		Config:       ac.Config,
		ProjectStore: projectStore,
	})

	request := &project_service.UploadRequest{}

	{
		header, err := c.FormFile("cartfile_content")
		if err != nil {
			c.Error(errors.New("Cartfile is empty"))
			return
		}

		file, err := header.Open()
		if err != nil {
			c.Error(errors.Wrap(err, 0))
			return
		}

		request.CartfileContent = file
	}

	{
		header, err := c.FormFile("podfile_content")
		if err != nil {
			c.Error(errors.New("Podfile is empty"))
			return
		}

		file, err := header.Open()
		if err != nil {
			c.Error(errors.Wrap(err, 0))
			return
		}

		request.PodfileLockContent = file
	}

	name, _ := c.GetPostForm("repository_name")

	request.RepositoryName = name

	userID := null.Int{}
	if user := middleware.GetCurrentUser(c); user != nil {
		userID.Int64 = int64(user.ID)
		userID.Valid = true
	}

	request.UserID = userID

	if project, err := service.Upload(request); err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

// GetProject returns the project
func GetProject(c *gin.Context) {
	uuid := c.Param("uuid")

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
