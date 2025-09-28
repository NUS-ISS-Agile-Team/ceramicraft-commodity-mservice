package router

import (
	"github.com/gin-gonic/gin"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/api"
	_ "github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/docs"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// swagger router
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/product-ms/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		merchantRouter := v1.Group("/merchant")
		{
			merchantRouter.POST("/add", api.AddProduct)
			merchantRouter.GET("/product/:id", api.GetProduct)
			merchantRouter.POST("/publish", api.PublishProduct)
			merchantRouter.POST("/unpublish", api.UnpublishProduct)
			merchantRouter.POST("/updateStock", api.UpdateProductStock)
		}

	}
	return r
}
