package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/ecom/presenter/healthcheck"
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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", healthcheck.Handler())

	return &http.Server{Addr: ":8080", Handler: r}, nil
}
