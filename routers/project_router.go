package routers

import (
	"MallApi/controllers"

	"github.com/gin-gonic/gin"
)

var project controllers.ProjectController

func SetProjectRouter(r *gin.Engine) {
	r.POST("/projects", project.CreateProject)

	r.GET("/projects", project.GetAllProject)
	r.GET("/project/:id", project.GetProjectByID)
	r.DELETE("/projects/:id", project.DeleteProject)
	r.PUT("/projects/:id", project.UpdateProject)
}
