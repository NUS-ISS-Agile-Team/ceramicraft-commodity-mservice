package api

import (
	"net/http"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/data"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/service"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/types"
	"github.com/gin-gonic/gin"
)

// AddProduct godoc
// @Summary 添加商品
// @Description 新增一个商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param product body types.ProductInfo true "商品信息"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /add [post]
func AddProduct(c *gin.Context) {
	var req types.ProductInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Errorf("AddProduct: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed(err.Error()))
		return
	}
	productId, err := service.GetProductServiceInstance().Create(c.Request.Context(), &req)
	if err != nil {
		log.Logger.Errorf("AddProduct: Failed to create product: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to create product"))
		return
	}
	c.JSON(http.StatusOK, data.ResponseSuccess(productId))
}
