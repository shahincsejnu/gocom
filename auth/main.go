package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"github.com/shahincsejnu/gocom/auth/presenter/consoleauth"
	"github.com/shahincsejnu/gocom/auth/presenter/healthcheck"
	"github.com/shahincsejnu/gocom/auth/presenter/login"
	"github.com/shahincsejnu/gocom/auth/presenter/signup"
	consoleauthuc "github.com/shahincsejnu/gocom/auth/usecase/consoleauth"
	loginuc "github.com/shahincsejnu/gocom/auth/usecase/login"
	signupuc "github.com/shahincsejnu/gocom/auth/usecase/signup"
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
	c, err := newDIContainer()
	if err != nil {
		return nil, err
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	err = c.Invoke(func(
		signupuc *signupuc.Usecase,
		loginuc *loginuc.Usecase,
		consoleAuth *consoleauthuc.Usecase,
	) {
		r.GET("/", healthcheck.Handler())
		r.POST("/signup", signup.Handler(signupuc))
		r.POST("/login", login.Handler(loginuc))
		r.POST("/console", consoleauth.Handler(consoleAuth))
	})
	if err != nil {
		return nil, err
	}

	return &http.Server{Addr: ":8080", Handler: r}, nil
}
