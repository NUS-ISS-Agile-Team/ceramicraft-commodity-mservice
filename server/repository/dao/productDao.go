package dao

import (
	"context"
	"sync"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"
	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(ctx context.Context, product *model.Product) (int, error)
	UpdateProduct(ctx context.Context, product *model.Product, tx *gorm.DB) error
	GetProductByID(ctx context.Context, id int) (*model.Product, error)
}

type ProductDaoImpl struct {
	db *gorm.DB
}

var (
	productOnce sync.Once
	productDao  *ProductDaoImpl
)

func GetProductDao() *ProductDaoImpl {
	productOnce.Do(func() {
		if productDao == nil {
			productDao = &ProductDaoImpl{db: repository.DB}
		}
	})
	return productDao
}

// CreateProduct 创建产品并返回ID
func (p *ProductDaoImpl) CreateProduct(ctx context.Context, product *model.Product) (int, error) {
	result := p.db.WithContext(ctx).Create(product)
	if result.Error != nil {
		log.Logger.Errorf("Failed to create product: %v", result.Error)
		return 0, result.Error
	}
	return int(product.ID), nil
}

// UpdateProduct 更新产品信息
func (p *ProductDaoImpl) UpdateProduct(ctx context.Context, product *model.Product, tx *gorm.DB) error {
	ret := tx.WithContext(ctx).Model(&model.Product{}).Where("id = ?", product.ID).Updates(product)
	if ret.Error != nil {
		log.Logger.Errorf("Failed to update product ID %d: %v", product.ID, ret.Error)
		return ret.Error
	}
	return nil
}

// GetProductByID 根据ID获取产品信息
func (p *ProductDaoImpl) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	var product model.Product
	result := p.db.WithContext(ctx).Where("id = ?", id).First(&product)
	if result.Error != nil {
		log.Logger.Errorf("Failed to get product by ID %d: %v", id, result.Error)
		return nil, result.Error
	}
	return &product, nil
}
