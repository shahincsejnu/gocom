package login

import (
	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"
	"golang.org/x/crypto/bcrypt"
)

func LoginCheck(user sqlcdb.User, password string) error {
	if err := verfiyPassword(password, user.Password); err != nil {
		return err
	}
	return nil
}

func verfiyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
