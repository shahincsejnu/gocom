package users_test

import (
	"net/http/httptest"
	"testing"

	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/users"
	usecase "github.com/shahincsejnu/gocom/ecom/usecase/users"
)

var (
	ctrl           *gomock.Controller
	handler        gin.HandlerFunc
	uc             *usecase.Usecase
	userRepository *MockUserRepository
	r              *gin.Engine
	response       *httptest.ResponseRecorder
)

func Test(t *testing.T) {
	apath, _ := filepath.Abs("../../../")
	_ = os.Chdir(apath)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Presenter/Users Suite")
}

func NewGetUserOneRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/users/:userID", handlerFunc)
	return router
}

func NewGetUserAddressesListRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/users/:userID/addresses", handlerFunc)
	return router
}

func NewGetUserOrdersListRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/users/:userID/orders", handlerFunc)
	return router
}
