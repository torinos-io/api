package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"

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
		PBXprojectContent:  pbxproj,
	}

	service.Upload(uploadRequest)

	// TODO: Create response by service
	cartfileName := ""
	podFileName := ""

	if cartfile != nil {
		cartfileName = cartfile.Filename
	}
	if podfile != nil {
		podFileName = podfile.Filename
	}

	//TODO: Return response immediately
	c.JSON(http.StatusOK, gin.H{
		"pbxprojfile_name": pbxproj.Filename,
		"cartfile_name":    cartfileName,
		"podfile_name":     podFileName,
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
