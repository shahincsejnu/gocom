package order_test

import (
	"net/http/httptest"
	"testing"

	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/order"
	usecase "github.com/shahincsejnu/gocom/ecom/usecase/order"
)

var (
	ctrl            *gomock.Controller
	handler         gin.HandlerFunc
	uc              *usecase.Usecase
	orderRepository *MockOrderRepository
	r               *gin.Engine
	response        *httptest.ResponseRecorder
)

func Test(t *testing.T) {
	apath, _ := filepath.Abs("../../../")
	_ = os.Chdir(apath)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Presenter/Order Suite")
}

func NewGetOrdersListRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/orders/:userID", handlerFunc)
	return router
}

func NewCreateOrderRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.POST("/orders", handlerFunc)
	return router
}

func NewGetOneOrderRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/orders/:orderID", handlerFunc)
	return router
}

func NewUpdateOneOrderRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.PATCH("/orders/:orderID", handlerFunc)
	return router
}
