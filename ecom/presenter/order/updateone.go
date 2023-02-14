package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	orderdom "github.com/shahincsejnu/gocom/ecom/domain/order"
	"github.com/shahincsejnu/gocom/ecom/usecase/order"
)

type updateOneResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func UpdateOneHandler(uc *order.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("orderID")

		err := Validate(orderID, Required, is.UUID)
		if err != nil {
			returnUpdateOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		r := new(orderdom.UpdateOptions)
		if err := c.ShouldBindJSON(&r); err != nil {
			returnUpdateOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		err = uc.UpdateOneOrder(c, r, orderID)
		if err != nil {
			returnUpdateOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnUpdateOneResponse(c, http.StatusOK, true, "Order is updated successfully")
	}
}

func returnUpdateOneResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &updateOneResponse{
		OK:      ok,
		Message: message,
	})
}
