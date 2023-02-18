package users

import (
	"context"

	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*sqlcdb.User, error)
	GetOrdersList(ctx context.Context, userID string) ([]sqlcdb.Order, error)
	GetAddressesList(ctx context.Context, userID string) ([]sqlcdb.Address, error)
}

type Usecase struct {
	UserRepository UserRepository
}

func (uc *Usecase) GetUserById(ctx context.Context, id string) (*sqlcdb.User, error) {
	return uc.UserRepository.GetById(ctx, id)
}

func (uc *Usecase) GetUserOrdersList(ctx context.Context, userID string) ([]sqlcdb.Order, error) {
	return uc.UserRepository.GetOrdersList(ctx, userID)
}

func (uc *Usecase) GetUserAddressesList(ctx context.Context, userID string) ([]sqlcdb.Address, error) {
	return uc.UserRepository.GetAddressesList(ctx, userID)
}
