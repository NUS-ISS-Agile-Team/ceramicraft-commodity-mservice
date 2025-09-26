package grpc

import (
	"context"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/common/productpb"
)

type ProductService struct {
	productpb.UnimplementedProductServiceServer
}

func (p *ProductService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	product := &productpb.Product{
		Id: req.Id,
		Name: "Sample Product",
		Price: 100,
	}
	return &productpb.GetProductResponse{Product: product}, nil
}
