package routers

import (
	"MallApi/controllers"

	"github.com/gin-gonic/gin"
)

var store controllers.StoreController

func SetUpStoreRouters(r *gin.Engine) {
	r.POST("/stores", store.CreateStore)
	r.PUT("store/:id", store.UpdateStore)
}
