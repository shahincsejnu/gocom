package product

import (
	"context"

	"github.com/shahincsejnu/gocom/ecom/domain/product"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]sqlcdb.Product, error)
	Create(ctx context.Context, opts *product.CreationOptions) (string, error)
	Update(ctx context.Context, opts *product.UpdateOptions, productID string) error
}

type Usecase struct {
	ProductRepo ProductRepository
}

func (uc *Usecase) GetAllProducts(ctx context.Context) ([]sqlcdb.Product, error) {
	return uc.ProductRepo.GetAll(ctx)
}

func (uc *Usecase) CreateProduct(ctx context.Context, opts *product.CreationOptions) (string, error) {
	return uc.ProductRepo.Create(ctx, opts)
}

func (uc *Usecase) UpdateProduct(ctx context.Context, opts *product.UpdateOptions, productID string) error {
	return uc.ProductRepo.Update(ctx, opts, productID)
}
