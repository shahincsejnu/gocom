package product_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	productdom "github.com/shahincsejnu/gocom/ecom/domain/product"
	productpres "github.com/shahincsejnu/gocom/ecom/presenter/product"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/product"
	productuc "github.com/shahincsejnu/gocom/ecom/usecase/product"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("UpdateProduct", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		productRepository = NewMockProductRepository(ctrl)

		uc = &productuc.Usecase{
			ProductRepo: productRepository,
		}
		handler = productpres.UpdateHandler(uc)
		response = httptest.NewRecorder()

		r = NewUpdateProductRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling update product endpoint correctly", func() {
		It("Should return status 200", func() {
			productId := uuid.NewString()

			creationRequest := productdom.UpdateOptions{
				Price:       10,
				Description: "demo product",
				Stock:       2,
			}
			jsonValue, _ := json.Marshal(creationRequest)

			productRepository.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/products/%s", productId)
			request, _ := http.NewRequest("PATCH", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling update product endpoint incorrect product id", func() {
		It("Should return status 400", func() {
			productId := "demo"

			creationRequest := productdom.UpdateOptions{
				Price:       10,
				Description: "demo product",
				Stock:       2,
			}
			jsonValue, _ := json.Marshal(creationRequest)

			productRepository.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/products/%s", productId)
			request, _ := http.NewRequest("PATCH", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
