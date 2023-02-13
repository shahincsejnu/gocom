package product

import (
	"context"

	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]sqlcdb.Product, error)
}

type Usecase struct {
	ProductRepo ProductRepository
}

func (uc *Usecase) GetAllProducts(ctx context.Context) ([]sqlcdb.Product, error) {
	return uc.ProductRepo.GetAll(ctx)
}
