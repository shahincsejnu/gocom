package address_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	addressdom "github.com/shahincsejnu/gocom/ecom/domain/address"
	addresspres "github.com/shahincsejnu/gocom/ecom/presenter/address"
	. "github.com/shahincsejnu/gocom/ecom/tests/usecase/address"
	addressuc "github.com/shahincsejnu/gocom/ecom/usecase/address"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("UpdateOneAddress", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		addressRepository = NewMockAddressRepository(ctrl)

		uc = &addressuc.Usecase{
			AddressRepo: addressRepository,
		}
		handler = addresspres.UpdateOneHandler(uc)
		response = httptest.NewRecorder()

		r = NewUpdateOneAddressRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling update one address endpoint correctly", func() {
		It("Should return status 200", func() {
			addressId := uuid.NewString()

			updateRequest := addressdom.UpdateOptions{
				Country:       "demo",
				City:          "demo",
				StreetAddress: "demo",
			}
			jsonValue, _ := json.Marshal(updateRequest)

			addressRepository.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/addresses/%s", addressId)
			request, _ := http.NewRequest("PATCH", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling update one address endpoint with incorrect order id", func() {
		It("Should return status 400", func() {
			addressId := "demo"

			updateRequest := addressdom.UpdateOptions{
				Country:       "demo",
				City:          "demo",
				StreetAddress: "demo",
			}
			jsonValue, _ := json.Marshal(updateRequest)

			addressRepository.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			requestString := fmt.Sprintf("/addresses/%s", addressId)
			request, _ := http.NewRequest("PATCH", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
