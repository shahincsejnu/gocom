package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	addressdom "github.com/shahincsejnu/gocom/ecom/domain/address"
	"github.com/shahincsejnu/gocom/ecom/usecase/address"
)

type updateOneResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func UpdateOneHandler(uc *address.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		addressID := c.Param("addressID")

		err := Validate(addressID, Required, is.UUID)
		if err != nil {
			returnUpdateOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		r := new(addressdom.UpdateOptions)
		if err := c.ShouldBindJSON(&r); err != nil {
			returnUpdateOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		err = uc.UpdateOneAddress(c, r, addressID)
		if err != nil {
			returnUpdateOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnUpdateOneResponse(c, http.StatusOK, true, "Address is updated successfully")
	}
}

func returnUpdateOneResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &updateOneResponse{
		OK:      ok,
		Message: message,
	})
}
