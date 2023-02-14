package order_test

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	orderdom "github.com/shahincsejnu/gocom/ecom/domain/order"
	orderpres "github.com/shahincsejnu/gocom/ecom/presenter/order"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/order"
	orderuc "github.com/shahincsejnu/gocom/ecom/usecase/order"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("CreateOrder", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		orderRepository = NewMockOrderRepository(ctrl)

		uc = &orderuc.Usecase{
			OrderRepo: orderRepository,
		}
		handler = orderpres.CreationHandler(uc)
		response = httptest.NewRecorder()

		r = NewCreateOrderRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling create order endpoint correctly", func() {
		creationRequest := orderdom.CreationOptions{
			ProductID: uuid.NewString(),
			UserID:    uuid.NewString(),
			AddressID: uuid.NewString(),
			Quantity:  1,
		}
		jsonValue, _ := json.Marshal(creationRequest)

		It("Should return status 200", func() {
			orderRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return("demo-id", nil)
			requestString := "/orders"
			request, _ := http.NewRequest("POST", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling create order endpoint incorrectly", func() {
		It("Should return status 400", func() {
			orderRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", nil)
			requestString := "/orders"
			request, _ := http.NewRequest("POST", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
