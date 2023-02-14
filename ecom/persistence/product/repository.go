package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/shahincsejnu/gocom/ecom/domain/product"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type Repository struct {
	DB *sqlcdb.Queries
}

func NewRepository(db *sqlcdb.Queries) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]sqlcdb.Product, error) {
	products, err := r.DB.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) GetOne(ctx context.Context, productID string) (*sqlcdb.Product, error) {
	prd, err := r.DB.GetProductById(ctx, productID)
	if err != nil {
		return nil, err
	}

	return &prd, nil
}

func (r *Repository) Create(ctx context.Context, opts *product.CreationOptions) (string, error) {
	arg := sqlcdb.CreateProductParams{
		ID:          uuid.NewString(),
		Name:        opts.Name,
		Price:       int32(opts.Price),
		Description: opts.Description,
		Stock:       int32(opts.Stock),
	}
	err := r.DB.CreateProduct(ctx, arg)
	if err != nil {
		return "", err
	}

	return arg.ID, nil
}

func (r *Repository) Update(ctx context.Context, opts *product.UpdateOptions, productID string) error {
	arg := sqlcdb.UpdateProductParams{
		ID:          productID,
		Price:       int32(opts.Price),
		Description: opts.Description,
		Stock:       int32(opts.Stock),
	}
	return r.DB.UpdateProduct(ctx, arg)
}
