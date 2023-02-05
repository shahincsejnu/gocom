package users_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	userpres "github.com/shahincsejnu/gocom/ecom/presenter/users"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/users"
	useruc "github.com/shahincsejnu/gocom/ecom/usecase/users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetUserOneById", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		userRepository = NewMockUserRepository(ctrl)

		uc = &useruc.Usecase{
			UserRepository: userRepository,
		}
		handler = userpres.GetOneHandler(uc)
		response = httptest.NewRecorder()

		r = NewGetUserOneRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling get one user endpoint with incorrect user id", func() {
		It("Should return status 400", func() {
			userId := "demo"

			userRepository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&sqlcdb.User{}, nil)
			requestString := fmt.Sprintf("/users/%s", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})

	When("Calling get one user endpoint with correct user id", func() {
		It("Should return status 200", func() {
			userId := uuid.NewString()

			userRepository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&sqlcdb.User{}, nil)
			requestString := fmt.Sprintf("/users/%s", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})
})
