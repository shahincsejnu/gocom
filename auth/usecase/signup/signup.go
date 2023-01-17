package signup

import (
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, name, email, password string) error
}

type Usecase struct {
	UserRepository UserRepository
}

func (uc *Usecase) Do(ctx context.Context, name, email, password string) error {
	if err := uc.UserRepository.Save(ctx, name, email, password); err != nil {
		return err
	}
	return nil
}
