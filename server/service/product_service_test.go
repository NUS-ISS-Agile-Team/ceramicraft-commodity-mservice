package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/dao/mocks"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/types"
	"github.com/golang/mock/gomock"
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
		t.Errorf("Expected error when publishing an already published product, got nil")
	}
}
