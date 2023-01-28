package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"
	userpersistence "github.com/shahincsejnu/gocom/auth/persistence/user"
	"github.com/shahincsejnu/gocom/auth/usecase/consoleauth"
	"github.com/shahincsejnu/gocom/auth/usecase/login"
	"github.com/shahincsejnu/gocom/auth/usecase/signup"
	"go.uber.org/dig"
)

func newDIContainer() (*dig.Container, error) {
	c := dig.New()

	pp := []interface{}{
		newSQLC,
		newUserRepository,
		newSignupUsecase,
		newLoginUsecase,
		newConsoleAuthUsecse,
	}
	for _, p := range pp {
		if err := c.Provide(p); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func newUserRepository(db *sqlcdb.Queries) *userpersistence.Repository {
	return &userpersistence.Repository{DB: db}
}

func newSignupUsecase(ur *userpersistence.Repository) *signup.Usecase {
	return &signup.Usecase{
		UserRepository: ur,
	}
}

func newLoginUsecase(ur *userpersistence.Repository) *login.Usecase {
	return &login.Usecase{
		UserRepository: ur,
	}
}

func newConsoleAuthUsecse(ur *userpersistence.Repository) *consoleauth.Usecase {
	return &consoleauth.Usecase{
		UserRepository: ur,
	}
}

func newSQLC() (*sqlcdb.Queries, error) {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		return nil, errors.New("database url env is not set")
	}

	db, err := sql.Open("postgres", dburl)
	if err != nil {
		return nil, errors.New("couldn't open the DB, because:" + err.Error())
	}

	var counter int
	for {
		if counter == 30 {
			log.Fatal("reached maximum number of attempt connecting to database")
		}

		fmt.Println("attempt to connect to database", "counter", counter)
		err := db.Ping()
		if err == nil {
			break
		}

		log.Println("attempt connecting to database failed, will be repeated in one second", "err", err)
		time.Sleep(time.Second)
		counter++
	}

	queries := sqlcdb.New(db)

	return queries, nil
}
