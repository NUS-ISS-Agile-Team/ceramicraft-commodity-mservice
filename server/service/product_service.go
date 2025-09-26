package service

import (
	"context"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/dao"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/types"
)

type ProductService interface {
	Create(ctx context.Context, product *types.ProductInfo) (productId int, err error)
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
