package address_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	addresspres "github.com/shahincsejnu/gocom/ecom/presenter/address"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/address"
	addressuc "github.com/shahincsejnu/gocom/ecom/usecase/address"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("DeleteOneAddress", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		addressRepository = NewMockAddressRepository(ctrl)

		uc = &addressuc.Usecase{
			AddressRepo: addressRepository,
		}
		handler = addresspres.DeleteOneHandler(uc)
		response = httptest.NewRecorder()

		r = NewDeleteOneAddressRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling delete one address endpoint correctly", func() {
		It("Should return status 200", func() {
			addressId := uuid.NewString()

			addressRepository.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/addresses/%s", addressId)
			request, _ := http.NewRequest("DELETE", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling delete one address endpoint with incorrect product id", func() {
		It("Should return status 400", func() {
			addressId := "demo"

			addressRepository.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/addresses/%s", addressId)
			request, _ := http.NewRequest("DELETE", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
