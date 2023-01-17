package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/shahincsejnu/gocom/auth/presenter/healthcheck"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	server, err := newServer()
	if err != nil {
		return err
	}

	log.Println("server is running")
	return server.ListenAndServe()
}

func newServer() (*http.Server, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", healthcheck.Handler())

	return &http.Server{Addr: ":8080", Handler: r}, nil
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
