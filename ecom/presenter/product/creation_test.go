package product_test

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	productdom "github.com/shahincsejnu/gocom/ecom/domain/product"
	productpres "github.com/shahincsejnu/gocom/ecom/presenter/product"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/product"
	productuc "github.com/shahincsejnu/gocom/ecom/usecase/product"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("CreateProduct", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		productRepository = NewMockProductRepository(ctrl)

		uc = &productuc.Usecase{
			ProductRepo: productRepository,
		}
		handler = productpres.CreationHandler(uc)
		response = httptest.NewRecorder()

		r = NewCreateProductRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling create product endpoint correctly", func() {
		creationRequest := productdom.CreationOptions{
			Name:        "demo",
			Price:       100,
			Description: "demo product",
			Stock:       1,
		}
		jsonValue, _ := json.Marshal(creationRequest)

		It("Should return status 200", func() {
			productRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return("demo-id", nil)
			requestString := "/products"
			request, _ := http.NewRequest("POST", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling create product endpoint incorrectly", func() {
		It("Should return status 400", func() {
			productRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", nil)
			requestString := "/products"
			request, _ := http.NewRequest("POST", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
