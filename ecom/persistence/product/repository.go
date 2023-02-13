package users

import (
	"context"

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
