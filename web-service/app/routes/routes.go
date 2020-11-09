package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {
	StoreRoutes(app)
	ProductRoutes(app)
}
