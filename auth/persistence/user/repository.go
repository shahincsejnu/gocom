package user

import (
	"context"

	"github.com/google/uuid"

	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"
)

type Repository struct {
	DB *sqlcdb.Queries
}

func (r *Repository) Save(ctx context.Context, name, email, password string) error {
	arg := sqlcdb.CreateUserParams{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	return r.DB.CreateUser(ctx, arg)
}

func (r *Repository) GetByName(ctx context.Context, name string) (sqlcdb.User, error) {
	return r.DB.GetUserByName(ctx, name)
}
