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

func (r *Repository) GetList(ctx context.Context, userID string) ([]sqlcdb.Order, error) {
	return r.DB.GetOrdersByUser(ctx, userID)
}
