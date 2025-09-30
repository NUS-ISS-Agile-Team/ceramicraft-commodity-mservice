package router

import (
	"github.com/gin-gonic/gin"

	_ "github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/docs"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/api"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/product-ms/v1")
	{
		// swagger router
		v1.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

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
			merchantRouter.POST("/images/upload-urls", api.GetImageUploadPresignURL)
		}

	}
	return r
}
