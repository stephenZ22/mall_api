package routers

import "github.com/gin-gonic/gin"

func RegisterRouters() *gin.Engine {

	r := gin.Default()

	SetUpUserRouters(r)

	return r
}
