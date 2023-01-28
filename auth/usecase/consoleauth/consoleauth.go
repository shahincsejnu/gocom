package consoleauth

import (
	"context"

	dtoken "github.com/shahincsejnu/gocom/auth/domain/token"
	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"
)

type UserRepository interface {
	GetByID(ctx context.Context, authID string) (*sqlcdb.User, error)
}

type Usecase struct {
	UserRepository UserRepository
}

func (uc *Usecase) Do(ctx context.Context, token string) (*sqlcdb.User, error) {
	userId, err := dtoken.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	usr, err := uc.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
