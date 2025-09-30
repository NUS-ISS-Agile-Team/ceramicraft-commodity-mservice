package api

import (
	"net/http"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/http/data"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/service"
	"github.com/gin-gonic/gin"
)

// GetImageUploadPresignURL godoc
// @Summary Get presigned URL for image upload
// @Description Get presigned URL for image upload
// @Tags Image
// @Accept json
// @Produce json
// @Param product body data.ImgUploadRequest true "image_type=(jpg|jpeg|png)"
// @Success 200 {object} data.ImgUploadResponse
// @Failure 400 {object} data.BaseResponse
// @Router /merchant/images/upload-urls [post]
func GetImageUploadPresignURL(c *gin.Context) {
	var req data.ImgUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Logger.Errorf("GetImageUploadPresignURL: Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, data.ResponseFailed(err.Error()))
		return
	}
	ret, err := service.GetImageService().GenUploadURL(c.Request.Context(), req.ImageType)
	if err != nil {
		log.Logger.Errorf("GetImageUploadPresignURL: Failed to generate upload url: %v", err)
		c.JSON(http.StatusInternalServerError, data.ResponseFailed("Failed to generate image uplaod url"))
		return
	}
	c.JSON(http.StatusOK, ret)
}
