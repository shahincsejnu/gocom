package address_test

import (
	"net/http/httptest"
	"testing"

	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/address"
	usecase "github.com/shahincsejnu/gocom/ecom/usecase/address"
)

var (
	ctrl              *gomock.Controller
	handler           gin.HandlerFunc
	uc                *usecase.Usecase
	addressRepository *MockAddressRepository
	r                 *gin.Engine
	response          *httptest.ResponseRecorder
)

func Test(t *testing.T) {
	apath, _ := filepath.Abs("../../../")
	_ = os.Chdir(apath)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Presenter/Address Suite")
}

func NewGetAddressListRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/address/:userID", handlerFunc)
	return router
}

func NewCreateAddressRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.POST("/addresses", handlerFunc)
	return router
}