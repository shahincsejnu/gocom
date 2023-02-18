package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/users"
)

type getUserOrdersListResponse struct {
	OK      bool           `json:"ok"`
	Message string         `json:"message"`
	Orders  []sqlcdb.Order `json:"orders"`
}

func GetUserOrdersListHandler(uc *users.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		err := Validate(userID, Required, is.UUID)
		if err != nil {
			returnGetUserOrdersListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		orders, err := uc.GetUserOrdersList(c, userID)
		if err != nil {
			returnGetUserOrdersListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetUserOrdersListResponse(c, http.StatusOK, true, "User Orders fetching is succeeded", orders)
	}
}

func returnGetUserOrdersListResponse(c *gin.Context, statusCode int, ok bool, message string, orders []sqlcdb.Order) {
	c.JSON(statusCode, &getUserOrdersListResponse{
		OK:      ok,
		Message: message,
		Orders:  orders,
	})
}
