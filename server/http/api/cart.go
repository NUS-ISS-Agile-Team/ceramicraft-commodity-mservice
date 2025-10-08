package api

import (
	"net/http"
	"strconv"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/data"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/service"
	"github.com/gin-gonic/gin"
)

// CreateCartItem godoc
// @Summary Create a cart item
// @Description Create a cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_item body data.CartItemBasicVO true "cart item info"
// @Success 200 {object} data.BaseResponse{data=data.CartItemBasicVO}
// @Failure 400 {object} data.BaseResponse
// @Failure 401 {object} data.BaseResponse
// @Failure 500 {object} data.BaseResponse
// @Router /customer/cart/items [post]
func CreateCartItem(c *gin.Context) {
	var req data.CartItemBasicVO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Errorf("CreateCartItem: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed(err.Error()))
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		log.Logger.Error("CreateCartItem: User ID not found in context")
		c.JSON(http.StatusUnauthorized, data.ResponseFailed("User not authenticated"))
		return
	}
	req.UserID = userID.(int)
	req.Selected = true
	bizErr := service.GetCartService().AddItem(c.Request.Context(), &req)
	if bizErr == nil {
		c.JSON(http.StatusOK, data.ResponseSuccess(req))
		return
	}
	if bizErr.Code == service.ProductCheckStatus_DBError {
		log.Logger.Errorf("CreateCartItem: Failed to add cart item: %v", bizErr)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to add cart item"))
		return
	}
	log.Logger.Errorf("CreateCartItem: Failed to add cart item: %v", bizErr)
	c.JSON(http.StatusBadRequest, data.ResponseFailed(bizErr.Message))
}

// UpdateCartItem godoc
// @Summary Update a cart item
// @Description Update a cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_item body data.CartItemBasicVO true "cart item info"
// @Success 200 {object} data.BaseResponse{data=data.CartItemBasicVO}
// @Failure 400 {object} data.BaseResponse
// @Failure 401 {object} data.BaseResponse
// @Failure 500 {object} data.BaseResponse
// @Router /customer/cart/items/:item_id [put]
func UpdateCartItem(c *gin.Context) {
	var req data.CartItemBasicVO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Errorf("UpdateCartItem: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed(err.Error()))
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		log.Logger.Error("UpdateCartItem: User ID not found in context")
		c.JSON(http.StatusUnauthorized, data.ResponseFailed("User not authenticated"))
		return
	}
	req.UserID = userID.(int)
	itemId := c.Param("item_id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		log.Logger.Errorf("UpdateCartItem: Invalid item_id parameter: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed("Invalid item_id parameter"))
		return
	}
	req.ID = id
	if req.ID <= 0 {
		log.Logger.Errorf("UpdateCartItem: item_id must be positive")
		c.JSON(http.StatusBadRequest, data.ResponseFailed("illegal item_id parameter"))
		return
	}
	bizErr := service.GetCartService().UpdateItem(c.Request.Context(), &req)
	if bizErr == nil {
		c.JSON(http.StatusOK, data.ResponseSuccess(req))
		return
	}
	if bizErr.Code == service.ProductCheckStatus_DBError {
		log.Logger.Errorf("UpdateCartItem: Failed to update cart item: %v", bizErr)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to update cart item"))
		return
	}
	log.Logger.Errorf("UpdateCartItem: Failed to update cart item: %v", bizErr)
	c.JSON(http.StatusBadRequest, data.ResponseFailed(bizErr.Message))
}

// DeleteCartItem godoc
// @Summary Delete a cart item
// @Description Delete a cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Param item_id path int true "cart item ID"
// @Success 200 {object} data.BaseResponse{data=string}
// @Failure 400 {object} data.BaseResponse
// @Failure 401 {object} data.BaseResponse
// @Failure 500 {object} data.BaseResponse
// @Router /customer/cart/items/:item_id [delete]
func DeleteCartItem(c *gin.Context) {
	itemId := c.Param("item_id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		log.Logger.Errorf("UpdateCartItem: Invalid item_id parameter: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed("Invalid item_id parameter"))
		return
	}
	if id <= 0 {
		log.Logger.Errorf("DeleteCartItem: item_id must be positive")
		c.JSON(http.StatusBadRequest, data.ResponseFailed("illegal item_id parameter"))
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		log.Logger.Error("DeleteCartItem: User ID not found in context")
		c.JSON(http.StatusUnauthorized, data.ResponseFailed("User not authenticated"))
		return
	}
	userId := userID.(int)
	err = service.GetCartService().DeleteItem(c.Request.Context(), id, userId)
	if err != nil {
		log.Logger.Errorf("DeleteCartItem: Failed to delete cart item: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to delete cart item"))
		return
	}
	// todo delete with itemId+userId
	c.JSON(http.StatusOK, data.ResponseSuccess(nil))
}

// GetUserCartInfo godoc
// @Summary Get user's cart info
// @Description Get user's cart info
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} data.BaseResponse{data=data.CartListVO}
// @Failure 401 {object} data.BaseResponse
// @Failure 500 {object} data.BaseResponse
// @Router /customer/cart [get]
func GetUserCartInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		log.Logger.Error("GetUserCartItemList: User ID not found in context")
		c.JSON(http.StatusUnauthorized, data.ResponseFailed("User not authenticated"))
		return
	}
	userID = userID.(int)
	log.Logger.Infof("GetUserCartInfo: userID=%d", userID)
	ret, err := service.GetCartService().GetCartItems(c.Request.Context(), userID.(int))
	if err != nil {
		log.Logger.Errorf("GetUserCartInfo: Failed to get cart items: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to get cart items"))
		return
	}
	c.JSON(http.StatusOK, data.ResponseSuccess(ret))
}

// GetCartSelctedNum godoc
// @Summary Get number of selected items in cart
// @Description Get number of selected items in cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} data.BaseResponse{data=int}
// @Failure 401 {object} data.BaseResponse
// @Failure 500 {object} data.BaseResponse
// @Router /customer/cart/selected-num [get]
func GetCartSelctedNum(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		log.Logger.Error("GetCartSelctedNum: User ID not found in context")
		c.JSON(http.StatusUnauthorized, data.ResponseFailed("User not authenticated"))
		return
	}
	userId := userID.(int)
	log.Logger.Infof("GetCartSelctedNum: userID=%d", userId)
	ret, err := service.GetCartService().GetCartSelectedItemCnt(c.Request.Context(), userId)
	if err != nil {
		log.Logger.Errorf("GetCartSelctedNum: Failed to get selected item count: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to get selected item count"))
		return
	}
	c.JSON(http.StatusOK, data.ResponseSuccess(ret))
}

// CartPriceEstimate godoc
// @Summary Calculate order price
// @Description Calculate order price
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} data.BaseResponse{data=data.CartPriceEstimateResult}
// @Failure 401 {object} data.BaseResponse
// @Failure 500 {object} data.BaseResponse
// @Router /customer/cart/price-estimate [get]
func GetEstimatePrice(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userIdInt := userId.(int)
	log.Logger.Infof("CalOrderPrice: userID=%d", userIdInt)
	ret, err := service.GetCartService().EstimatePrice(c.Request.Context(), userIdInt)
	if err != nil {
		log.Logger.Errorf("CalOrderPrice: Failed to estimate price: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to estimate price"))
		return
	}
	c.JSON(http.StatusOK, data.ResponseSuccess(ret))
}
