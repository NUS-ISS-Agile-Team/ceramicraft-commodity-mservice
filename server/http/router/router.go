package router

import (
	"github.com/gin-gonic/gin"

	_ "github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/docs"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/api"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-user-mservice/common/middleware"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

const (
	servicePrefix = "/product-ms/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group(servicePrefix)
	{
		// swagger router
		baseRouter.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

		baseRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		merchantRouter := baseRouter.Group("/merchant")
		{
			merchantRouter.Use(middleware.AuthMiddleware())
			merchantRouter.POST("/add", api.AddProduct)
			merchantRouter.GET("/product/:id", api.GetProduct)
			merchantRouter.POST("/publish", api.PublishProduct)
			merchantRouter.POST("/unpublish", api.UnpublishProduct)
			merchantRouter.POST("/updateStock", api.UpdateProductStock)
			merchantRouter.POST("/images/upload-urls", api.GetImageUploadPresignURL)
			merchantRouter.GET("/list", api.GetMerchantProductList)
		}

		customerRouter := baseRouter.Group("/customer")
		{
			customerRouter.GET("/list", api.GetCustomerProductList)

			authed := customerRouter.Group("")
			{
				authed.Use(middleware.AuthMiddleware())
			}
		}
	}
	return r
}
