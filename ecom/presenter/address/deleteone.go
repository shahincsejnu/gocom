package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shahincsejnu/gocom/ecom/usecase/address"
)

type deleteOneResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func DeleteOneHandler(uc *address.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		addressID := c.Param("addressID")

		err := Validate(addressID, Required, is.UUID)
		if err != nil {
			returnDeleteOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		err = uc.DeleteOneAddress(c, addressID)
		if err != nil {
			returnDeleteOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnDeleteOneResponse(c, http.StatusOK, true, "Address deletion is succeeded")
	}
}

func returnDeleteOneResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &deleteOneResponse{
		OK:      ok,
		Message: message,
	})
}
