package service

import (
	"context"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/dao"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/repository/model"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/types"
)

type ProductService interface {
	Create(ctx context.Context, product *types.ProductInfo) error
}

type ProductServiceImpl struct {
	productDao dao.ProductDao
}

var (
	ProductServiceInstance = &ProductServiceImpl{
		productDao: dao.GetProductDao(),
	}
)

func (p *ProductServiceImpl) Create(ctx context.Context, product *types.ProductInfo) error {
	_, err := p.productDao.CreateProduct(ctx, &model.Product{
		Name:  product.Name,
		Price: product.Price,
		Desc:  product.Desc,
		Stock: product.Stock,
		PicInfo: product.PicInfo,
		Status: product.Status,
	})
	if err != nil {
		log.Logger.Errorf("ProductService: Failed to create product: %v", err)
		return err
	}
	return nil
}

