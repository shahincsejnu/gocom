package product_test

import (
	"net/http/httptest"
	"testing"

	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/product"
	usecase "github.com/shahincsejnu/gocom/ecom/usecase/product"
)

var (
	ctrl              *gomock.Controller
	handler           gin.HandlerFunc
	uc                *usecase.Usecase
	productRepository *MockProductRepository
	r                 *gin.Engine
	response          *httptest.ResponseRecorder
)

func Test(t *testing.T) {
	apath, _ := filepath.Abs("../../../")
	_ = os.Chdir(apath)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Presenter/Product Suite")
}

func NewGetProductListRouter(handlerFunc gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET("/products", handlerFunc)
	return router
}
