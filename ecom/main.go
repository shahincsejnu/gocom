package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/ecom/presenter/address"
	"github.com/shahincsejnu/gocom/ecom/presenter/healthcheck"
	"github.com/shahincsejnu/gocom/ecom/presenter/order"
	"github.com/shahincsejnu/gocom/ecom/presenter/product"
	"github.com/shahincsejnu/gocom/ecom/presenter/users"
	addressuc "github.com/shahincsejnu/gocom/ecom/usecase/address"
	orderuc "github.com/shahincsejnu/gocom/ecom/usecase/order"
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
		orderUC *orderuc.Usecase,
		addressUC *addressuc.Usecase,
	) {
		r.GET("/", healthcheck.Handler())
		r.GET("/users/:userID", users.GetOneHandler(userUC))
		r.GET("/users/:userID/orders", users.GetUserOrdersListHandler(userUC))
		r.GET("/users/:userID/addresses", users.GetUserAddressesListHandler(userUC))

		r.GET("/products", product.GetListHandler(productUC))
		r.POST("/products", product.CreationHandler(productUC))
		r.PATCH("/products/:productID", product.UpdateHandler(productUC))
		r.GET("/products/:productID", product.GetOneHandler(productUC))
		r.DELETE("/products/:productID", product.DeleteOneHandler(productUC))

		r.POST("/orders", order.CreationHandler(orderUC))
		r.GET("/orders/:orderID", order.GetOneHandler(orderUC))
		r.PATCH("/orders/:orderID", order.UpdateOneHandler(orderUC))
		r.DELETE("/orders/:orderID", order.DeleteOneHandler(orderUC))

		r.POST("/addresses", address.CreationHandler(addressUC))
		r.GET("/addresses/:addressID", address.GetOneHandler(addressUC))
		r.PATCH("/addresses/:addressID", address.UpdateOneHandler(addressUC))
		r.DELETE("/addresses/:addressID", address.DeleteOneHandler(addressUC))
	})
	if err != nil {
		return nil, err
	}

	return &http.Server{Addr: ":8080", Handler: r}, nil
}
