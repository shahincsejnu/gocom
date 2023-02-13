package product_test

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	productpres "github.com/shahincsejnu/gocom/ecom/presenter/product"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/product"
	productuc "github.com/shahincsejnu/gocom/ecom/usecase/product"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetProductList", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		productRepository = NewMockProductRepository(ctrl)

		uc = &productuc.Usecase{
			ProductRepo: productRepository,
		}
		handler = productpres.GetListHandler(uc)
		response = httptest.NewRecorder()

		r = NewGetProductListRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling get product list endpoint correctly", func() {
		It("Should return status 200", func() {
			productRepository.EXPECT().GetAll(gomock.Any()).Return([]sqlcdb.Product{}, nil)
			requestString := "/products"
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})
})
