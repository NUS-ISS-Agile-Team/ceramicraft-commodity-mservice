package router

import (
	"github.com/gin-gonic/gin"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/api"
	_ "github.com/NUS-ISS-Agile-Team/ceramicraft-user-mservice/server/docs"
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
		v1.POST("/add", api.AddProduct)
		v1.GET("/product/:id", api.GetProduct)
		v1.POST("/publish", api.PublishProduct)
		v1.POST("/unpublish", api.UnpublishProduct)
	}
	return r
}
