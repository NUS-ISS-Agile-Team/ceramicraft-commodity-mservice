package dao

import (
	"context"
	"fmt"
	"sync"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"

	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(ctx context.Context, product *model.Product) (productId int, err error)
	UpdateStockWithCAS(ctx context.Context, id int, version int, newStock int) error
	GetProductByID(ctx context.Context, id int) (*model.Product, error)
	GetProductByIDs(ctx context.Context, ids []int) ([]*model.Product, error)
	UpdateProductStatus(ctx context.Context, id int, status int) error
	UpdateProductStock(ctx context.Context, id int, stock int) error
	ListProduct(ctx context.Context, q ListProductQuery) ([]*model.Product, int, error)
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

// UpdateStockWithCAS
func (p *ProductDaoImpl) UpdateStockWithCAS(ctx context.Context, id, version, newStock int) error {
	ret := p.db.WithContext(ctx).Model(&model.Product{}).Where("id = ? AND version = ?", id, version).Update("stock", newStock)
	if ret.Error != nil {
		log.Logger.Errorf("Failed to update product ID %d: %v", id, ret.Error)
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

func (p *ProductDaoImpl) GetProductByIDs(ctx context.Context, ids []int) ([]*model.Product, error) {
	var products []*model.Product
	result := p.db.WithContext(ctx).Where("id IN ?", ids).Find(&products)
	if result.Error != nil {
		log.Logger.Errorf("Failed to get products by IDs %v: %v", ids, result.Error)
		return nil, result.Error
	}
	return products, nil
}

// UpdateProductStock 更新商品库存
func (p *ProductDaoImpl) UpdateProductStock(ctx context.Context, id int, stock int) error {
	result := p.db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id).Update("stock", stock)
	if result.Error != nil {
		log.Logger.Errorf("Failed to update product stock, ID: %d, stock: %d, error: %v", id, stock, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		err := fmt.Errorf("product not found with ID: %d", id)
		log.Logger.Error(err)
		return err
	}
	return nil
}

// ListProduct 查询商品列表
func (p *ProductDaoImpl) ListProduct(ctx context.Context, q ListProductQuery) ([]*model.Product, int, error) {
	var products []*model.Product
	var total int64

	query := p.db.WithContext(ctx).Model(&model.Product{})

	if q.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+q.Keyword+"%")
	}

	if q.Category != "" {
		query = query.Where("category = ?", q.Category)
	}

	// 用户侧只能看到上架的商品
	if q.IsCustomer {
		query = query.Where("status = ?", 1)
	}

	if q.OrderBy == 0 {
		query = query.Order("updated_at DESC")
	} else {
		query = query.Order("updated_at")
	}

	if q.Limit == 0 {
		q.Limit = 10
	}

	// 获取总数
	err := query.Count(&total).Error
	if err != nil {
		log.Logger.Errorf("Failed to count products: %v", err)
		return nil, 0, err
	}

	err = query.Offset(q.Offset).Limit(q.Limit).Find(&products).Error
	if err != nil {
		log.Logger.Errorf("Failed to get products ordered by time: %v", err)
		return nil, 0, err
	}

	return products, int(total), nil
}

// UpdateProductStatus 更新商品状态
func (p *ProductDaoImpl) UpdateProductStatus(ctx context.Context, id int, status int) error {
	result := p.db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		log.Logger.Errorf("Failed to update product status, ID: %d, status: %d, error: %v", id, status, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		err := fmt.Errorf("product not found with ID: %d", id)
		log.Logger.Error(err)
		return err
	}
	return nil
}
