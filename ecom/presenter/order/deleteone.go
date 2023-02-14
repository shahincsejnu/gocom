package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shahincsejnu/gocom/ecom/usecase/order"
)

type deleteOneResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func DeleteOneHandler(uc *order.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("orderID")

		err := Validate(orderID, Required, is.UUID)
		if err != nil {
			returnDeleteOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		err = uc.DeleteOneOrder(c, orderID)
		if err != nil {
			returnDeleteOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnDeleteOneResponse(c, http.StatusOK, true, "Order deletion is succeeded")
	}
}

func returnDeleteOneResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &deleteOneResponse{
		OK:      ok,
		Message: message,
	})
}
