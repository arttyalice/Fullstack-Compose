package routes

import (
	"deeploy-exam/app/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(app *gin.RouterGroup) {
	prodCtrl := controllers.NewProductController()
	product := app.Group("/product")
	product.GET("/list", prodCtrl.GetList)
	product.GET("/info/:product_id", prodCtrl.GetInfo)
	product.POST("/create", prodCtrl.Create)
	product.PUT("/update", prodCtrl.Update)
	product.DELETE("/delete/:product_id", prodCtrl.Delete)

	catCtrl := controllers.NewProductCategoryController()
	category := product.Group("/category")
	category.GET("/list", catCtrl.GetList)
	category.GET("/info/:category_id", catCtrl.GetInfo)
	category.POST("/create", catCtrl.Create)
	category.PUT("/update", catCtrl.Update)
	category.DELETE("/delete/:category_id", catCtrl.Delete)
}
