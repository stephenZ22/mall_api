package routers

import (
	"MallApi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouters() *gin.Engine {

	r := gin.Default()
	r.Use(middlewares.Cors())

	SetUpUserRouters(r)
	SetProjectRouter(r)

	return r
}
