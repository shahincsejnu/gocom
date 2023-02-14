package order_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	orderpres "github.com/shahincsejnu/gocom/ecom/presenter/order"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/order"
	orderuc "github.com/shahincsejnu/gocom/ecom/usecase/order"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetOrdersList", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		orderRepository = NewMockOrderRepository(ctrl)

		uc = &orderuc.Usecase{
			OrderRepo: orderRepository,
		}
		handler = orderpres.GetListHandler(uc)
		response = httptest.NewRecorder()

		r = NewGetOrdersListRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling get orders list endpoint correctly", func() {
		It("Should return status 200", func() {
			userId := uuid.NewString()

			orderRepository.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]sqlcdb.Order{}, nil)
			requestString := fmt.Sprintf("/orders/%s", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling get orders list endpoint incorrectly", func() {
		It("Should return status 400", func() {
			userId := "demo"

			orderRepository.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]sqlcdb.Order{}, nil)
			requestString := fmt.Sprintf("/orders/%s", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
