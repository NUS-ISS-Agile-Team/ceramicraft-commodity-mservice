package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/docs"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/api"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/metrics"
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

		baseRouter.Use(metrics.MetricsMiddleware())
		baseRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
			merchantRouter.POST("/products", api.AddProduct)
			merchantRouter.GET("/product/:id", api.GetProductMerchant)
			merchantRouter.PATCH("/product-status", api.UpdateProductStatus)
			merchantRouter.PATCH("/product-stock", api.UpdateProductStock)
			merchantRouter.POST("/images/upload-urls", api.GetImageUploadPresignURL)
			merchantRouter.GET("/products", api.GetMerchantProductList)
			merchantRouter.PUT("/products", api.EditProductInfo)
		}

		customerRouter := baseRouter.Group("/customer")
		{
			customerRouter.GET("/products", api.GetCustomerProductList)
			customerRouter.GET("/product/:id", api.GetProductCustomer)

			authed := customerRouter.Group("")
			{
				authed.Use(middleware.AuthMiddleware())
				authed.GET("/cart", api.GetUserCartInfo)
				authed.POST("/cart/items", api.CreateCartItem)
				authed.PUT("/cart/items/:item_id", api.UpdateCartItem)
				authed.DELETE("/cart/items/:item_id", api.DeleteCartItem)
				authed.GET("/cart/selected-num", api.GetCartSelctedNum)
				authed.GET("/cart/price-estimate", api.GetEstimatePrice)
			}
		}
	}
	return r
}
