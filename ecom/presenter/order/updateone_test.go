package order_test

import (
	"bytes"
	"encoding/json"
	"fmt"

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

var _ = Describe("UpdateOneOrder", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		orderRepository = NewMockOrderRepository(ctrl)

		uc = &orderuc.Usecase{
			OrderRepo: orderRepository,
		}
		handler = orderpres.UpdateOneHandler(uc)
		response = httptest.NewRecorder()

		r = NewUpdateOneOrderRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling update one order endpoint correctly", func() {
		It("Should return status 200", func() {
			orderId := uuid.NewString()

			updateRequest := orderdom.UpdateOptions{
				Quantity: 2,
			}
			jsonValue, _ := json.Marshal(updateRequest)

			orderRepository.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/orders/%s", orderId)
			request, _ := http.NewRequest("PATCH", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling update one order endpoint with incorrect order id", func() {
		It("Should return status 400", func() {
			orderId := "demo"

			updateRequest := orderdom.UpdateOptions{
				Quantity: 3,
			}
			jsonValue, _ := json.Marshal(updateRequest)

			orderRepository.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/orders/%s", orderId)
			request, _ := http.NewRequest("PATCH", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
