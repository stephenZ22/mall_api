package routers

import (
	"MallApi/controllers"

	"github.com/gin-gonic/gin"
)

var project controllers.ProjectController

func SetProjectRouter(r *gin.Engine) {
	r.POST("/projects", project.CreateProject)
}
