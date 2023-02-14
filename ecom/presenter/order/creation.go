package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	orderdom "github.com/shahincsejnu/gocom/ecom/domain/order"
	"github.com/shahincsejnu/gocom/ecom/usecase/order"
)

type creationResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	OrderID string `json:"orderID"`
}

func CreationHandler(uc *order.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(orderdom.CreationOptions)
		if err := c.ShouldBindJSON(&r); err != nil {
			returnCreationResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		orderID, err := uc.CreateOrder(c, r)
		if err != nil {
			returnCreationResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		returnCreationResponse(c, http.StatusOK, true, "Order is created successfully", orderID)
	}
}

func returnCreationResponse(c *gin.Context, statusCode int, ok bool, message, orderID string) {
	c.JSON(statusCode, &creationResponse{
		OK:      ok,
		Message: message,
		OrderID: orderID,
	})
}
