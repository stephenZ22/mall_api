package routers

import (
	"MallApi/controllers"
	"MallApi/middlewares"

	"github.com/gin-gonic/gin"
)

var user controllers.UserController

func SetUpUserRouters(r *gin.Engine) {
	r.GET("/hello", user.Hello)

	r.POST("/users", user.CreateUser)
	r.GET("/users", user.GetAllUser)
	r.POST("/users/login", user.UserLogin)
	r.DELETE(("/users/:id"), middlewares.CheckLogin(), user.DeletedUser)
}
