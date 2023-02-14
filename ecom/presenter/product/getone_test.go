package product_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	productpres "github.com/shahincsejnu/gocom/ecom/presenter/product"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/product"
	productuc "github.com/shahincsejnu/gocom/ecom/usecase/product"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetOneProduct", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		productRepository = NewMockProductRepository(ctrl)

		uc = &productuc.Usecase{
			ProductRepo: productRepository,
		}
		handler = productpres.GetOneHandler(uc)
		response = httptest.NewRecorder()

		r = NewGetOneProductRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling get one product endpoint correctly", func() {
		It("Should return status 200", func() {
			productId := uuid.NewString()

			productRepository.EXPECT().GetOne(gomock.Any(), gomock.Any()).Return(&sqlcdb.Product{}, nil)
			requestString := fmt.Sprintf("/products/%s", productId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling get one product endpoint with incorrect product id", func() {
		It("Should return status 400", func() {
			productId := "demo"

			productRepository.EXPECT().GetOne(gomock.Any(), gomock.Any()).Return(nil, nil)
			requestString := fmt.Sprintf("/products/%s", productId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
