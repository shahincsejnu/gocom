package address

import (
	"context"

	"github.com/shahincsejnu/gocom/ecom/domain/address"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type AddressRepository interface {
	GetList(ctx context.Context, userID string) ([]sqlcdb.Address, error)
	GetOne(ctx context.Context, addressID string) (*sqlcdb.Address, error)
	Create(ctx context.Context, opts *address.CreationOptions) (string, error)
	UpdateOne(ctx context.Context, opts *address.UpdateOptions, addressID string) error
}

type Usecase struct {
	AddressRepo AddressRepository
}

func (uc *Usecase) GetAddressesList(ctx context.Context, userID string) ([]sqlcdb.Address, error) {
	return uc.AddressRepo.GetList(ctx, userID)
}

func (uc *Usecase) CreateAddress(ctx context.Context, opts *address.CreationOptions) (string, error) {
	return uc.AddressRepo.Create(ctx, opts)
}

func (uc *Usecase) GetOneAddress(ctx context.Context, addressID string) (*sqlcdb.Address, error) {
	return uc.AddressRepo.GetOne(ctx, addressID)
}

func (uc *Usecase) UpdateOneAddress(ctx context.Context, opts *address.UpdateOptions, addressID string) error {
	return uc.AddressRepo.UpdateOne(ctx, opts, addressID)
}
