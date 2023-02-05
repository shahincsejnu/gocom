package users

import (
	"context"

	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*sqlcdb.User, error)
}

type Usecase struct {
	UserRepository UserRepository
}

func (uc *Usecase) GetUserById(ctx context.Context, id string) (*sqlcdb.User, error) {
	return uc.UserRepository.GetById(ctx, id)
}
