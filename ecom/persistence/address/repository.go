package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/shahincsejnu/gocom/ecom/domain/address"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type Repository struct {
	DB *sqlcdb.Queries
}

func NewRepository(db *sqlcdb.Queries) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetList(ctx context.Context, userID string) ([]sqlcdb.Address, error) {
	return r.DB.GetAddressesByUser(ctx, userID)
}

func (r *Repository) Create(ctx context.Context, opts *address.CreationOptions) (string, error) {
	arg := sqlcdb.CreateAddressParams{
		ID:            uuid.NewString(),
		UserID:        opts.UserID,
		Country:       opts.Country,
		City:          opts.City,
		StreetAddress: opts.StreetAddress,
	}
	err := r.DB.CreateAddress(ctx, arg)
	if err != nil {
		return "", err
	}

	return arg.ID, nil
}
