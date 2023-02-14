package address_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	addresspres "github.com/shahincsejnu/gocom/ecom/presenter/address"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/address"
	addressuc "github.com/shahincsejnu/gocom/ecom/usecase/address"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetOrdersList", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		addressRepository = NewMockAddressRepository(ctrl)

		uc = &addressuc.Usecase{
			AddressRepo: addressRepository,
		}
		handler = addresspres.GetListHandler(uc)
		response = httptest.NewRecorder()

		r = NewGetAddressListRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling get addresses list endpoint correctly", func() {
		It("Should return status 200", func() {
			userId := uuid.NewString()

			addressRepository.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]sqlcdb.Address{}, nil)
			requestString := fmt.Sprintf("/address/%s", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling get orders list endpoint incorrectly", func() {
		It("Should return status 400", func() {
			userId := "demo"

			addressRepository.EXPECT().GetList(gomock.Any(), gomock.Any()).Return([]sqlcdb.Address{}, nil)
			requestString := fmt.Sprintf("/address/%s", userId)
			request, _ := http.NewRequest("GET", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
