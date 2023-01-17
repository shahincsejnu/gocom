package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"
	userpersistence "github.com/shahincsejnu/gocom/auth/persistence/user"
	"github.com/shahincsejnu/gocom/auth/usecase/signup"
	"go.uber.org/dig"
)

func newDIContainer() (*dig.Container, error) {
	c := dig.New()

	pp := []interface{}{
		newSQLC,
		newUserRepository,
		newSignupUsecase,
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

func newSQLC() (*sqlcdb.Queries, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

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

		log.Fatal("attempt connecting to database failed, will be repeated in one second", "err", err)
		time.Sleep(time.Second)
		counter++
	}

	queries := sqlcdb.New(db)

	return queries, nil
}
