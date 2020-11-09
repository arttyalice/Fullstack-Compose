package routes

import (
	"deeploy-exam/app/controllers"

	"github.com/gin-gonic/gin"
)

func StoreRoutes(app *gin.RouterGroup) {
	storeCtrl := controllers.NewStoreController()
	storeProductCtrl := controllers.NewStoreProductController()

	store := app.Group("/store")
	{
		store.GET("/list", storeCtrl.GetList)
		store.GET("/info/:store_id", storeCtrl.GetInfo)
		store.POST("/create", storeCtrl.Create)
		store.PUT("/update", storeCtrl.Update)
		store.DELETE("/delete/:store_id", storeCtrl.Delete)
	}

	product := store.Group("/product")
	{
		product.POST("/create", storeProductCtrl.Create)
		product.PUT("/update", storeProductCtrl.Update)
		product.DELETE("/delete/:req_id", storeProductCtrl.Delete)
	}
}
