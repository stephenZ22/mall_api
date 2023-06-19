package routers

import (
	"MallApi/controllers"
	"MallApi/middlewares"

	"github.com/gin-gonic/gin"
)

var user controllers.UserController

func SetUpUserRouters(r *gin.Engine) {
	r.GET("/hello", user.Hello)

	r.POST("/users", middlewares.CheckLogin(), user.CreateUser)
	r.POST("/users/login", user.UserLogin)
	r.GET("/users", middlewares.CheckLogin(), user.GetAllUser)
	r.DELETE(("/users/:id"), middlewares.CheckLogin(), user.DeletedUser)
}
