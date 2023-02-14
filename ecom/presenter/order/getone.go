package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/order"
)

type getOneResponse struct {
	OK      bool          `json:"ok"`
	Message string        `json:"message"`
	Order   *sqlcdb.Order `json:"order"`
}

func GetOneHandler(uc *order.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("orderID")

		err := Validate(orderID, Required, is.UUID)
		if err != nil {
			returnGetOneResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		ord, err := uc.GetOneOrder(c, orderID)
		if err != nil {
			returnGetOneResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetOneResponse(c, http.StatusOK, true, "Order fetching is succeeded", ord)
	}
}

func returnGetOneResponse(c *gin.Context, statusCode int, ok bool, message string, ord *sqlcdb.Order) {
	c.JSON(statusCode, &getOneResponse{
		OK:      ok,
		Message: message,
		Order:   ord,
	})
}
