package routers

import (
	"MallApi/controllers"

	"github.com/gin-gonic/gin"
)

var user controllers.UserController

func SetUpUserRouters(r *gin.Engine) {
	r.GET("/hello", user.Hello)

	r.POST("/users", user.CreateUser)
}