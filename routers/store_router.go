package routers

import (
	"MallApi/controllers"
	"MallApi/middlewares"

	"github.com/gin-gonic/gin"
)

var store controllers.StoreController

func SetUpStoreRouters(r *gin.Engine) {
	r.POST("/stores", middlewares.CheckLogin(), store.CreateStore)

	r.GET("/stores", middlewares.CheckLogin(), store.GetAllStores)
	r.GET("store/:id", middlewares.CheckLogin(), store.GetStoreById)
	r.DELETE("store:id", middlewares.CheckLogin(), store.DeleteStoreById)
	r.PUT("store/:id", middlewares.CheckLogin(), store.UpdateStore)
}
