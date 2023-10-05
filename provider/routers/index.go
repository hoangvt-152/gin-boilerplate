package routers

import (
	"gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
	route.GET("/v1/example/", controllers.GetData)
	route.GET("/v1/category",controllers.GetCategories)
	route.POST("/v1/category",controllers.CreateCategories)
	route.GET("/v1/category/:id",controllers.GetCategoryById)


	//Add All route
	//TestRoutes(route)
}
