package users_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	userspres "github.com/shahincsejnu/gocom/ecom/presenter/users"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/users"
	usersuc "github.com/shahincsejnu/gocom/ecom/usecase/users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetUserOrdersList", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		userRepository = NewMockUserRepository(ctrl)

		uc = &usersuc.Usecase{
			UserRepository: userRepository,
		}
		handler = userspres.GetUserOrdersListHandler(uc)
		response = httptest.NewRecorder()

		r = NewGetUserOrdersListRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling get user orders list endpoint correctly", func() {
		It("Should return status 200", func() {
			userId := uuid.NewString()

			userRepository.EXPECT().GetOrdersList(gomock.Any(), gomock.Any()).Return([]sqlcdb.Order{}, nil)
			requestString := fmt.Sprintf("/users/%s/orders", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling get user orders list endpoint incorrectly", func() {
		It("Should return status 400", func() {
			userId := "demo"

			userRepository.EXPECT().GetOrdersList(gomock.Any(), gomock.Any()).Return([]sqlcdb.Order{}, nil)
			requestString := fmt.Sprintf("/users/%s/orders", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
