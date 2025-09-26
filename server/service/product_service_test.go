package service

import (
	"context"
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
		Name:    "Test Product",
		Price:   200,
		Desc:    "This is a test product",
		Stock:   50,
		PicInfo: "http://example.com/pic.jpg",
		Status:  0,
	}

	m.EXPECT().CreateProduct(context.Background(), gomock.Eq(productModel)).Return(1, nil)

	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	productInfo := &types.ProductInfo{
		Name:    "Test Product",
		Price:   200,
		Desc:    "This is a test product",
		Stock:   50,
		PicInfo: "http://example.com/pic.jpg",
		Status:  0,
	}

	err := testProductServiceImpl.Create(context.Background(), productInfo)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
