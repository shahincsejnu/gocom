package address_test

import (
	"bytes"
	"encoding/json"

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

var _ = Describe("CreateAddress", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		ctrl = gomock.NewController(GinkgoT())
		addressRepository = NewMockAddressRepository(ctrl)

		uc = &addressuc.Usecase{
			AddressRepo: addressRepository,
		}
		handler = addresspres.CreationHandler(uc)
		response = httptest.NewRecorder()

		r = NewCreateAddressRouter(handler)

		defer ctrl.Finish()
	})

	When("Calling create address endpoint correctly", func() {
		creationRequest := addressdom.CreationOptions{
			UserID:        uuid.NewString(),
			Country:       "demo",
			City:          "demo",
			StreetAddress: "demo",
		}
		jsonValue, _ := json.Marshal(creationRequest)

		It("Should return status 200", func() {
			addressRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return("demo-id", nil)
			requestString := "/addresses"
			request, _ := http.NewRequest("POST", requestString, bytes.NewBuffer(jsonValue))
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(200))
		})
	})

	When("Calling create address endpoint incorrectly", func() {
		It("Should return status 400", func() {
			addressRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return("", nil)
			requestString := "/addresses"
			request, _ := http.NewRequest("POST", requestString, nil)
			r.ServeHTTP(response, request)
			Expect(response.Code).Should(Equal(400))
		})
	})
})
