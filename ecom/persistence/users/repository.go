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

func (r *Repository) GetById(ctx context.Context, id string) (*sqlcdb.User, error) {
	user, err := r.DB.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
