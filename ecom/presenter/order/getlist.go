package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/order"
)

type getListResponse struct {
	OK      bool           `json:"ok"`
	Message string         `json:"message"`
	Orders  []sqlcdb.Order `json:"orders"`
}

func GetListHandler(uc *order.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		err := Validate(userID, Required, is.UUID)
		if err != nil {
			returnGetListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		orders, err := uc.GetOrdersList(c, userID)
		if err != nil {
			returnGetListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetListResponse(c, http.StatusOK, true, "Orders fetching is succeeded", orders)
	}
}

func returnGetListResponse(c *gin.Context, statusCode int, ok bool, message string, orders []sqlcdb.Order) {
	c.JSON(statusCode, &getListResponse{
		OK:      ok,
		Message: message,
		Orders:  orders,
	})
}
