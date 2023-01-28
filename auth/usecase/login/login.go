package login

import (
	"context"

	"github.com/shahincsejnu/gocom/auth/domain/login"
	"github.com/shahincsejnu/gocom/auth/domain/token"
	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"
)

type UserRepository interface {
	GetByName(ctx context.Context, name string) (sqlcdb.User, error)
}

type Usecase struct {
	UserRepository UserRepository
}

func (uc *Usecase) Do(ctx context.Context, name, password string) (string, error) {
	user, err := uc.UserRepository.GetByName(ctx, name)
	if err != nil {
		return "", err
	}

	if err := login.LoginCheck(user, password); err != nil {
		return "", err
	}

	return token.GenerateToken(user.ID)
}
