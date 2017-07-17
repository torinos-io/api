package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/torinos-io/api/server/middleware"
	project_service "github.com/torinos-io/api/service/project_service"
	project_store "github.com/torinos-io/api/store/project_store"
)

// CreateProject creates project
func CreateProject(c *gin.Context) {

	ac := middleware.GetAppContext(c)
	userStore := project_store.New(ac.MainDB)
	service := project_service.New(project_service.Context{
		Config:       ac.Config,
		ProjectStore: userStore,
	})

	cartfile, _ := c.FormFile("cartfile_content")
	podfile, _ := c.FormFile("podfile_content")
	pbxproj, _ := c.FormFile("pbxporj_content")

	uploadRequest := &project_service.UploadRequest{
		CartfileContent:    cartfile,
		PodfileLockContent: podfile,
		PBXprojectContent:  pbxproj,
	}

	service.Upload(uploadRequest)

	c.JSON(http.StatusOK, gin.H{
		"message": "Project created",
	})
}

// GetProject returns the project
func GetProject(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetProject",
		"uuid":    uuid,
		"result":  "",
	})
}
