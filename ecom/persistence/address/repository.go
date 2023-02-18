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

func (r *Repository) GetOne(ctx context.Context, addressID string) (*sqlcdb.Address, error) {
	addr, err := r.DB.GetAddressById(ctx, addressID)
	if err != nil {
		return nil, err
	}

	return &addr, nil
}

func (r *Repository) UpdateOne(ctx context.Context, opts *address.UpdateOptions, addressID string) error {
	arg := sqlcdb.UpdateAddressParams{
		ID:            addressID,
		Country:       opts.Country,
		City:          opts.City,
		StreetAddress: opts.StreetAddress,
	}
	return r.DB.UpdateAddress(ctx, arg)
}

func (r *Repository) DeleteOne(ctx context.Context, addressID string) error {
	return r.DB.DeleteAddress(ctx, addressID)
}
