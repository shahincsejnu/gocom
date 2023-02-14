package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	addressdom "github.com/shahincsejnu/gocom/ecom/domain/address"
	"github.com/shahincsejnu/gocom/ecom/usecase/address"
)

type creationResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	AddressID string `json:"addressID"`
}

func CreationHandler(uc *address.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(addressdom.CreationOptions)
		if err := c.ShouldBindJSON(&r); err != nil {
			returnCreationResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		addressID, err := uc.CreateAddress(c, r)
		if err != nil {
			returnCreationResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		returnCreationResponse(c, http.StatusOK, true, "Address is created successfully", addressID)
	}
}

func returnCreationResponse(c *gin.Context, statusCode int, ok bool, message, addressID string) {
	c.JSON(statusCode, &creationResponse{
		OK:        ok,
		Message:   message,
		AddressID: addressID,
	})
}
