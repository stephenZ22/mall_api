package routers

import (
	"MallApi/controllers"
	"MallApi/middlewares"

	"github.com/gin-gonic/gin"
)

var project controllers.ProjectController

func SetProjectRouter(r *gin.Engine) {
	r.POST("/projects", middlewares.CheckLogin(), project.CreateProject)

	r.GET("/projects", middlewares.CheckLogin(), project.GetAllProject)
	r.GET("/project/:id", middlewares.CheckLogin(), project.GetProjectByID)
	r.DELETE("/projects/:id", middlewares.CheckLogin(), project.DeleteProject)
	r.PUT("/projects/:id", middlewares.CheckLogin(), project.UpdateProject)
}
