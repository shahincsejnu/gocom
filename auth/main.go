package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello from auth!")

	db, err := newSQLC()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	users, err := db.GetUsers(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Numbers of users are:", len(users))

	arg := sqlcdb.CreateUserParams{
		ID:       "4e600115-a406-4cab-ad08-9d6b2364a481",
		Name:     "Sha",
		Email:    "sha@gmail.com",
		Password: "123",
	}
	err = db.CreateUser(context.TODO(), arg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Create success!")

	users, err = db.GetUsers(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Numbers of users are:", len(users))
	fmt.Println(users)
}

func newSQLC() (*sqlcdb.Queries, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, errors.New("databse url env is not set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, errors.New("couldn't open the DB, because:" + err.Error())
	}

	var counter int
	for {
		if counter == 30 {
			return nil, errors.New("reached maximum number of attempt to connect the DB")
		}

		fmt.Println("attempt to connect the DB, counter:", counter)
		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(time.Second)
		counter++
	}

	queries := sqlcdb.New(db)

	return queries, nil
}
