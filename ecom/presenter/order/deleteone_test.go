package order_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	orderpres "github.com/shahincsejnu/gocom/ecom/presenter/order"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/order"
	orderuc "github.com/shahincsejnu/gocom/ecom/usecase/order"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("DeleteOneOrder", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		orderRepository = NewMockOrderRepository(ctrl)

		uc = &orderuc.Usecase{
			OrderRepo: orderRepository,
		}
		handler = orderpres.DeleteOneHandler(uc)
		response = httptest.NewRecorder()

		r = NewDeleteOneOrderRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling delete one order endpoint correctly", func() {
		It("Should return status 200", func() {
			orderId := uuid.NewString()

			orderRepository.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/orders/%s", orderId)
			request, _ := http.NewRequest("DELETE", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling delete one order endpoint with incorrect product id", func() {
		It("Should return status 400", func() {
			orderId := "demo"

			orderRepository.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/orders/%s", orderId)
			request, _ := http.NewRequest("DELETE", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
