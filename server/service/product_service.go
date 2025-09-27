package service

import (
	"context"
	"fmt"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/dao"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/types"
)

type ProductService interface {
	Create(ctx context.Context, product *types.ProductInfo) (productId int, err error)
	GetProductByID(ctx context.Context, id int) (productInfo *types.ProductInfo, err error)
	PublishProduct(ctx context.Context, id int) error
	UnpublishProduct(ctx context.Context, id int) error
}

type ProductServiceImpl struct {
	productDao dao.ProductDao
}

func GetProductServiceInstance() *ProductServiceImpl {
	return &ProductServiceImpl{
		productDao: dao.GetProductDao(),
	}
}

func (p *ProductServiceImpl) Create(ctx context.Context, product *types.ProductInfo) (productId int, err error) {
	id, err := p.productDao.CreateProduct(ctx, &model.Product{
		Name:             product.Name,
		Category:         product.Category,
		Price:            product.Price,
		Desc:             product.Desc,
		Stock:            product.Stock,
		PicInfo:          product.PicInfo,
		Weight:           product.Weight,
		Material:         product.Material,
		Capacity:         product.Capacity,
		Dimensions:       product.Dimensions,
		CareInstructions: product.CareInstructions,
		Status:           product.Status,
	})
	if err != nil {
		log.Logger.Errorf("ProductService: Failed to create product: %v", err)
		return -1, err
	}
	return id, nil
}

// GetProductByID 根据ID获取产品信息 (商家侧，无论是否上架都可以看到)
func (p *ProductServiceImpl) GetProductByID(ctx context.Context, id int) (productInfo *types.ProductInfo, err error) {
	product, err := p.productDao.GetProductByID(ctx, id)
	if err != nil {
		log.Logger.Errorf("ProductService: Failed to get product by ID: %v", err)
		return nil, err
	}
	if product == nil {
		return nil, nil
	}
	return &types.ProductInfo{
		Name:             product.Name,
		Category:         product.Category,
		Price:            product.Price,
		Desc:             product.Desc,
		Stock:            product.Stock,
		PicInfo:          product.PicInfo,
		Weight:           product.Weight,
		Material:         product.Material,
		Capacity:         product.Capacity,
		Dimensions:       product.Dimensions,
		CareInstructions: product.CareInstructions,
		Status:           product.Status,
	}, nil
}

const (
	ProductStatusUnpublished = 0 // 下架状态
	ProductStatusPublished   = 1 // 上架状态
)

// PublishProduct 上架商品
func (p *ProductServiceImpl) PublishProduct(ctx context.Context, id int) error {
	// 获取商品当前信息
	product, err := p.productDao.GetProductByID(ctx, id)
	if err != nil {
		log.Logger.Errorf("PublishProduct: Failed to get product by ID: %v", err)
		return err
	}
	if product == nil {
		return fmt.Errorf("product not found with ID: %d", id)
	}

	// 检查当前状态
	if product.Status == ProductStatusPublished {
		return fmt.Errorf("product (ID: %d) is already published", id)
	}

	// 更新状态为已上架
	err = p.productDao.UpdateProductStatus(ctx, id, ProductStatusPublished)
	if err != nil {
		log.Logger.Errorf("PublishProduct: Failed to update product status: %v", err)
		return err
	}

	return nil
}

// UnpublishProduct 下架商品
func (p *ProductServiceImpl) UnpublishProduct(ctx context.Context, id int) error {
	// 获取商品当前信息
	product, err := p.productDao.GetProductByID(ctx, id)
	if err != nil {
		log.Logger.Errorf("UnpublishProduct: Failed to get product by ID: %v", err)
		return err
	}
	if product == nil {
		return fmt.Errorf("product not found with ID: %d", id)
	}

	// 检查当前状态
	if product.Status == ProductStatusUnpublished {
		return fmt.Errorf("product (ID: %d) is already unpublished", id)
	}

	// 更新状态为已下架
	err = p.productDao.UpdateProductStatus(ctx, id, ProductStatusUnpublished)
	if err != nil {
		log.Logger.Errorf("UnpublishProduct: Failed to update product status: %v", err)
		return err
	}

	return nil
}
