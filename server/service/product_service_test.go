package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/dao"
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

func TestProductServiceImpl_GetPublishedProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)
	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	// 测试获取已上架商品
	publishedProduct := &model.Product{
		Name:             "已上架商品",
		Category:         "茶具",
		Price:            10000,
		Desc:             "精美陶瓷茶具",
		Stock:            100,
		Status:           1, // 已上架
		PicInfo:          "pic1.jpg",
		Dimensions:       "10x10x10",
		Material:         "陶瓷",
		Weight:           "1kg",
		Capacity:         "500ml",
		CareInstructions: "小心轻放",
	}
	m.EXPECT().GetProductByID(context.Background(), 1).Return(publishedProduct, nil)

	productInfo, err := testProductServiceImpl.GetPublishedProductByID(context.Background(), 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if productInfo == nil {
		t.Error("Expected product info, got nil")
	} else {
		// 验证返回的商品信息是否正确
		if productInfo.Name != publishedProduct.Name {
			t.Errorf("Expected name %s, got %s", publishedProduct.Name, productInfo.Name)
		}
		if productInfo.Status != publishedProduct.Status {
			t.Errorf("Expected status %d, got %d", publishedProduct.Status, productInfo.Status)
		}
	}

	// 测试获取未上架商品（应返回nil）
	unpublishedProduct := &model.Product{
		Name:   "未上架商品",
		Status: 0, // 未上架
	}
	m.EXPECT().GetProductByID(context.Background(), 2).Return(unpublishedProduct, nil)

	productInfo, err = testProductServiceImpl.GetPublishedProductByID(context.Background(), 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if productInfo != nil {
		t.Error("Expected nil for unpublished product, got product info")
	}

	// 测试获取不存在的商品
	m.EXPECT().GetProductByID(context.Background(), 3).Return(nil, nil)

	productInfo, err = testProductServiceImpl.GetPublishedProductByID(context.Background(), 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if productInfo != nil {
		t.Error("Expected nil for non-existent product, got product info")
	}

	// 测试数据库错误的情况
	m.EXPECT().GetProductByID(context.Background(), 4).Return(nil, errors.New("database error"))

	productInfo, err = testProductServiceImpl.GetPublishedProductByID(context.Background(), 4)
	if err == nil {
		t.Error("Expected database error, got nil")
	}
	if productInfo != nil {
		t.Error("Expected nil product info when database error occurs")
	}
}

func TestProductServiceImpl_GetProductList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockProductDao(ctrl)

	// 准备测试数据
	mockProducts := []*model.Product{
		{
			Name:             "陶瓷茶具1",
			Category:         "茶具",
			Price:            10000,
			Desc:             "精美陶瓷茶具",
			Stock:            100,
			Status:           1, // 已上架
			PicInfo:          "pic1.jpg",
			Dimensions:       "10x10x10",
			Material:         "陶瓷",
			Weight:           "1kg",
			Capacity:         "500ml",
			CareInstructions: "小心轻放",
		},
		{
			Name:             "陶瓷花瓶",
			Category:         "装饰品",
			Price:            20000,
			Desc:             "中式陶瓷花瓶",
			Stock:            50,
			Status:           0, // 未上架
			PicInfo:          "pic2.jpg",
			Dimensions:       "20x20x30",
			Material:         "陶瓷",
			Weight:           "2kg",
			Capacity:         "2L",
			CareInstructions: "防摔",
		},
	}

	testCases := []struct {
		name        string
		query       types.GetProductListQuery
		mockResult  []*model.Product
		mockCount   int
		mockError   error
		expectCount int
		expectLen   int
		expectError bool
	}{
		{
			name: "成功获取商家端全部商品列表",
			query: types.GetProductListQuery{
				Offset:     0,
				Limit:      10,
				IsCustomer: false,
				OrderBy:    0,
			},
			mockResult:  mockProducts,
			mockCount:   2,
			mockError:   nil,
			expectCount: 2,
			expectLen:   2,
			expectError: false,
		},
		{
			name: "成功获取用户端商品列表(只显示已上架)",
			query: types.GetProductListQuery{
				Offset:     0,
				Limit:      10,
				IsCustomer: true,
				OrderBy:    0,
			},
			mockResult:  mockProducts[:1], // 只返回已上架的商品
			mockCount:   1,
			mockError:   nil,
			expectCount: 1,
			expectLen:   1,
			expectError: false,
		},
		{
			name: "按关键词搜索",
			query: types.GetProductListQuery{
				Keyword:    "茶具",
				Offset:     0,
				Limit:      10,
				IsCustomer: true,
				OrderBy:    0,
			},
			mockResult:  mockProducts[:1],
			mockCount:   1,
			mockError:   nil,
			expectCount: 1,
			expectLen:   1,
			expectError: false,
		},
		{
			name: "按分类筛选",
			query: types.GetProductListQuery{
				Category:   "茶具",
				Offset:     0,
				Limit:      10,
				IsCustomer: true,
				OrderBy:    0,
			},
			mockResult:  mockProducts[:1],
			mockCount:   1,
			mockError:   nil,
			expectCount: 1,
			expectLen:   1,
			expectError: false,
		},
		{
			name: "数据库错误",
			query: types.GetProductListQuery{
				Offset:     0,
				Limit:      10,
				IsCustomer: true,
				OrderBy:    0,
			},
			mockResult:  nil,
			mockCount:   0,
			mockError:   errors.New("database error"),
			expectCount: -1,
			expectLen:   0,
			expectError: true,
		},
		{
			name: "空结果",
			query: types.GetProductListQuery{
				Keyword:    "不存在的商品",
				Offset:     0,
				Limit:      10,
				IsCustomer: true,
				OrderBy:    0,
			},
			mockResult:  []*model.Product{},
			mockCount:   0,
			mockError:   nil,
			expectCount: 0,
			expectLen:   0,
			expectError: false,
		},
		{
			name: "按更新时间升序",
			query: types.GetProductListQuery{
				Offset:     0,
				Limit:      10,
				IsCustomer: true,
				OrderBy:    1, // 升序
			},
			mockResult:  mockProducts[:1],
			mockCount:   1,
			mockError:   nil,
			expectCount: 1,
			expectLen:   1,
			expectError: false,
		},
	}

	testProductServiceImpl := &ProductServiceImpl{
		productDao: m,
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 设置Mock期望
			m.EXPECT().ListProduct(gomock.Any(), dao.ListProductQuery{
				Keyword:    tc.query.Keyword,
				Category:   tc.query.Category,
				Offset:     tc.query.Offset,
				Limit:      tc.query.Limit,
				IsCustomer: tc.query.IsCustomer,
				OrderBy:    tc.query.OrderBy,
			}).Return(tc.mockResult, tc.mockCount, tc.mockError)

			// 调用被测试的方法
			products, count, err := testProductServiceImpl.GetProductList(context.Background(), tc.query)

			// 验证结果
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if count != tc.expectCount {
				t.Errorf("Expected count %d but got %d", tc.expectCount, count)
			}

			if len(products) != tc.expectLen {
				t.Errorf("Expected %d products but got %d", tc.expectLen, len(products))
			}

			// 如果有返回结果，验证字段映射是否正确
			if len(products) > 0 {
				for i, p := range products {
					if p.Name != tc.mockResult[i].Name {
						t.Errorf("Product name mismatch at index %d: expected %s but got %s", i, tc.mockResult[i].Name, p.Name)
					}
					if p.Category != tc.mockResult[i].Category {
						t.Errorf("Product category mismatch at index %d: expected %s but got %s", i, tc.mockResult[i].Category, p.Category)
					}
					if p.Price != tc.mockResult[i].Price {
						t.Errorf("Product price mismatch at index %d: expected %d but got %d", i, tc.mockResult[i].Price, p.Price)
					}
					if p.Stock != tc.mockResult[i].Stock {
						t.Errorf("Product stock mismatch at index %d: expected %d but got %d", i, tc.mockResult[i].Stock, p.Stock)
					}
					if p.Status != tc.mockResult[i].Status {
						t.Errorf("Product status mismatch at index %d: expected %d but got %d", i, tc.mockResult[i].Status, p.Status)
					}
				}
			}
		})
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
