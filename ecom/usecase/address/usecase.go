package address

import (
	"context"

	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type AddressRepository interface {
	GetList(ctx context.Context, userID string) ([]sqlcdb.Address, error)
}

type Usecase struct {
	AddressRepo AddressRepository
}

func (uc *Usecase) GetAddressesList(ctx context.Context, userID string) ([]sqlcdb.Address, error) {
	return uc.AddressRepo.GetList(ctx, userID)
}
