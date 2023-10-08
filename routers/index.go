package routers

import (
	"gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
	route.GET("/v1/example/", controllers.GetData)

	route.GET("/v1/category", controllers.GetCategories)
	route.POST("/v1/category", controllers.CreateCategory)
	route.GET("/v1/category/:id", controllers.GetCategoryById)
	route.PUT("/v1/category/:id", controllers.UpdateCategory)
	route.DELETE("/v1/category/:id", controllers.DeleteCategory)

	route.GET("/v1/provider", controllers.GetProviders)
	route.POST("/v1/provider", controllers.CreateProvider)
	route.GET("/v1/provider/:id", controllers.GetProviderById)
	route.PUT("/v1/provider/:id", controllers.UpdateProvider)
	route.DELETE("/v1/provider/:id", controllers.DeleteProvider)

	route.GET("/v1/production", controllers.GetProviderProductions)
	route.POST("/v1/production", controllers.CreateProviderProduction)
	route.GET("/v1/production/:id", controllers.GetProviderProductionById)
	route.PUT("/v1/production/:id", controllers.UpdateProviderProduction)
	route.DELETE("/v1/production/:id", controllers.DeleteProviderProduction)

	
	route.GET("/v1/deal", controllers.GetProductDeals)
	route.POST("/v1/deal", controllers.CreateProductDeal)
	route.GET("/v1/deal/:id", controllers.GetProductDealById)
	route.PUT("/v1/deal/:id", controllers.UpdateProductDeal)
	route.DELETE("/v1/deal/:id", controllers.DeleteProductDeal)


	route.GET("/v1/rule", controllers.GetProviderRules)
	route.POST("/v1/rule", controllers.CreateProviderRule)
	route.GET("/v1/rule/:id", controllers.GetProviderRuleById)
	route.PUT("/v1/rule/:id", controllers.UpdateProviderRule)
	route.DELETE("/v1/rule/:id", controllers.DeleteProviderRule)


	//Add All route
	//TestRoutes(route)
}
