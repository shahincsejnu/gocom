package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	addressper "github.com/shahincsejnu/gocom/ecom/persistence/address"
	orderper "github.com/shahincsejnu/gocom/ecom/persistence/order"
	productper "github.com/shahincsejnu/gocom/ecom/persistence/product"
	userper "github.com/shahincsejnu/gocom/ecom/persistence/users"
	"github.com/shahincsejnu/gocom/ecom/usecase/address"
	"github.com/shahincsejnu/gocom/ecom/usecase/order"
	"github.com/shahincsejnu/gocom/ecom/usecase/product"
	"github.com/shahincsejnu/gocom/ecom/usecase/users"
	"go.uber.org/dig"
)

func newDIContainer() (*dig.Container, error) {
	c := dig.New()

	pp := []interface{}{
		newSQLC,
		userper.NewRepository,
		productper.NewRepository,
		orderper.NewRepository,
		addressper.NewRepository,
		newUserUsecase,
		newProductUsecase,
		newOrderUsecase,
		newAddressUsecase,
	}
	for _, p := range pp {
		if err := c.Provide(p); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func newAddressUsecase(ur *addressper.Repository) *address.Usecase {
	return &address.Usecase{
		AddressRepo: ur,
	}
}

func newOrderUsecase(ur *orderper.Repository) *order.Usecase {
	return &order.Usecase{
		OrderRepo: ur,
	}
}

func newProductUsecase(ur *productper.Repository) *product.Usecase {
	return &product.Usecase{
		ProductRepo: ur,
	}
}

func newUserUsecase(ur *userper.Repository) *users.Usecase {
	return &users.Usecase{
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
