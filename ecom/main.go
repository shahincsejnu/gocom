package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/ecom/presenter/healthcheck"
	"github.com/shahincsejnu/gocom/ecom/presenter/product"
	"github.com/shahincsejnu/gocom/ecom/presenter/users"
	productuc "github.com/shahincsejnu/gocom/ecom/usecase/product"
	useruc "github.com/shahincsejnu/gocom/ecom/usecase/users"
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

	log.Println("server is running...")
	return server.ListenAndServe()
}

func newServer() (*http.Server, error) {
	c, err := newDIContainer()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	err = c.Invoke(func(
		userUC *useruc.Usecase,
		productUC *productuc.Usecase,
	) {
		r.GET("/", healthcheck.Handler())
		r.GET("/users/:userID", users.GetOneHandler(userUC))
		r.GET("/products", product.GetListHandler(productUC))
		r.POST("/products", product.CreationHandler(productUC))
		r.PATCH("/products/:productID", product.UpdateHandler(productUC))
		r.GET("/products/:productID", product.GetOneHandler(productUC))
	})
	if err != nil {
		return nil, err
	}

	return &http.Server{Addr: ":8080", Handler: r}, nil
}
