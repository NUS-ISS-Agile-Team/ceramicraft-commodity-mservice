package api

import (
	"net/http"
	"strconv"

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

// GetProduct godoc
// @Summary 获取商品详情
// @Description 根据商品ID获取商品详细信息
// @Tags 商品
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} data.BaseResponse{data=types.ProductInfo} "成功"
// @Failure 400 {object} data.BaseResponse "请求参数错误"
// @Failure 404 {object} data.BaseResponse "商品不存在"
// @Failure 500 {object} data.BaseResponse "服务器内部错误"
// @Router /product/{id} [get]
func GetProduct(c *gin.Context) {
	// 解析路径参数
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Logger.Errorf("GetProduct: Invalid product ID: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed("Invalid product ID"))
		return
	}

	// 调用 service 层获取商品信息
	product, err := service.GetProductServiceInstance().GetProductByID(c.Request.Context(), id)
	if err != nil {
		log.Logger.Errorf("GetProduct: Failed to get product details: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to get product details"))
		return
	}

	// 如果没找到商品
	if product == nil {
		c.JSON(http.StatusNotFound, data.ResponseFailed("Product not found"))
		return
	}

	// 返回商品信息
	c.JSON(http.StatusOK, data.ResponseSuccess(product))
}
