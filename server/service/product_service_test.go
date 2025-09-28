package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/dao/mocks"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/types"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

// mockgen -source=dao/productDao.go -destination=dao/mocks/productDao_mock.go -package=mocks

func TestProductServiceImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)

	productModel := &model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}

	m.EXPECT().CreateProduct(context.Background(), gomock.Eq(productModel)).Return(1, nil)

	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	productInfo := &types.ProductInfo{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}

	_, err := testProductServiceImpl.Create(context.Background(), productInfo)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func init() {
	// 初始化测试用logger
	logger, _ := zap.NewDevelopment()
	log.Logger = logger.Sugar()
}

func TestProductServiceImpl_GetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)

	m.EXPECT().GetProductByID(context.Background(), 1).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	productInfo, err := testProductServiceImpl.GetProductByID(context.Background(), 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedProductInfo := &types.ProductInfo{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}

	if !reflect.DeepEqual(productInfo, expectedProductInfo) {
		t.Errorf("Product info mismatch:\ngot: %+v\nwant: %+v", productInfo, expectedProductInfo)
	}

	m.EXPECT().GetProductByID(context.Background(), 2).Return(nil, errors.New("product not found"))

	_, err = testProductServiceImpl.GetProductByID(context.Background(), 2)
	if err == nil {
		t.Errorf("Expected error when getting a non-existent product, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 3).Return(nil, nil)

	productInfo, _ = testProductServiceImpl.GetProductByID(context.Background(), 3)
	if productInfo != nil {
		t.Errorf("Expected nil product info for non-existent product, got %+v", productInfo)
	}
}

func TestProductServiceImpl_PublishProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)

	m.EXPECT().GetProductByID(context.Background(), 1).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	m.EXPECT().UpdateProductStatus(context.Background(), 1, 1).Return(nil)

	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	err := testProductServiceImpl.PublishProduct(context.Background(), 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	m.EXPECT().GetProductByID(context.Background(), 2).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           1,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	err = testProductServiceImpl.PublishProduct(context.Background(), 2)
	if err == nil {
		t.Errorf("Expected error when publishing an already published product, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 3).Return(nil, errors.New("product not found"))

	err = testProductServiceImpl.PublishProduct(context.Background(), 3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 4).Return(nil, nil)
	err = testProductServiceImpl.PublishProduct(context.Background(), 4)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 5).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	m.EXPECT().UpdateProductStatus(context.Background(), 5, 1).Return(errors.New("database error"))
	err = testProductServiceImpl.PublishProduct(context.Background(), 5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestProductServiceImpl_UnpublishProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)

	m.EXPECT().GetProductByID(context.Background(), 1).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           1,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	m.EXPECT().UpdateProductStatus(context.Background(), 1, 0).Return(nil)

	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	err := testProductServiceImpl.UnpublishProduct(context.Background(), 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	m.EXPECT().GetProductByID(context.Background(), 2).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           0,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	err = testProductServiceImpl.UnpublishProduct(context.Background(), 2)
	if err == nil {
		t.Errorf("Expected error when unpublishing an already unpublished product, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 3).Return(nil, errors.New("product not found"))

	err = testProductServiceImpl.UnpublishProduct(context.Background(), 3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 4).Return(nil, nil)
	err = testProductServiceImpl.UnpublishProduct(context.Background(), 4)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	m.EXPECT().GetProductByID(context.Background(), 5).Return(&model.Product{
		Name:             "Test Product",
		Price:            200,
		Desc:             "This is a test product",
		Stock:            50,
		PicInfo:          "http://example.com/pic.jpg",
		Status:           1,
		Category:         "Test Category",
		Weight:           "1kg",
		Material:         "Plastic",
		Capacity:         "500ml",
		Dimensions:       "10x10x10cm",
		CareInstructions: "Handle with care",
	}, nil)

	m.EXPECT().UpdateProductStatus(context.Background(), 5, 0).Return(errors.New("database error"))

	err = testProductServiceImpl.UnpublishProduct(context.Background(), 5)
	if err == nil {
		t.Errorf("Expected database error, got nil")
	}
}

func TestProductServiceImpl_UpdateProductStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)
	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	// 测试成功更新库存的情况
	m.EXPECT().GetProductByID(context.Background(), 1).Return(&model.Product{
		Name:   "Test Product",
		Stock:  50,
		Status: 0,
	}, nil)
	m.EXPECT().UpdateProductStock(context.Background(), 1, 60).Return(nil)

	err := testProductServiceImpl.UpdateProductStock(context.Background(), 1, 60)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 测试产品不存在的情况
	m.EXPECT().GetProductByID(context.Background(), 2).Return(nil, errors.New("product not found"))

	err = testProductServiceImpl.UpdateProductStock(context.Background(), 2, 30)
	if err == nil {
		t.Errorf("Expected error for non-existent product, got nil")
	}

	// 测试产品为nil的情况
	m.EXPECT().GetProductByID(context.Background(), 3).Return(nil, nil)

	err = testProductServiceImpl.UpdateProductStock(context.Background(), 3, 30)
	if err == nil {
		t.Errorf("Expected error for nil product, got nil")
	}

	// 测试更新库存失败的情况
	m.EXPECT().GetProductByID(context.Background(), 4).Return(&model.Product{
		Name:   "Test Product",
		Stock:  50,
		Status: 0,
	}, nil)
	m.EXPECT().UpdateProductStock(context.Background(), 4, 70).Return(errors.New("database error"))

	err = testProductServiceImpl.UpdateProductStock(context.Background(), 4, 70)
	if err == nil {
		t.Errorf("Expected database error, got nil")
	}

	// 测试更新负数库存的情况
	err = testProductServiceImpl.UpdateProductStock(context.Background(), 5, -10)
	if err == nil {
		t.Errorf("Expected error for negative stock, got nil")
	}

	// 测试更新库存为零的情况
	m.EXPECT().GetProductByID(context.Background(), 6).Return(&model.Product{
		Name:   "Test Product",
		Stock:  50,
		Status: 0,
	}, nil)
	m.EXPECT().UpdateProductStock(context.Background(), 6, 0).Return(nil)

	err = testProductServiceImpl.UpdateProductStock(context.Background(), 6, 0)
	if err != nil {
		t.Errorf("Expected no error for zero stock, got %v", err)
	}

	// 测试更新已上架商品库存的情况
	m.EXPECT().GetProductByID(context.Background(), 7).Return(&model.Product{
		Name:   "Test Product",
		Stock:  50,
		Status: 1,
	}, nil)
	
	err = testProductServiceImpl.UpdateProductStock(context.Background(), 7, 60)
	if err == nil {
		t.Errorf("Expected error when updating stock for published product, got nil")
	}
}
