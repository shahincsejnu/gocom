package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/address"
)

type getListResponse struct {
	OK        bool             `json:"ok"`
	Message   string           `json:"message"`
	Addresses []sqlcdb.Address `json:"addresses"`
}

func GetListHandler(uc *address.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		err := Validate(userID, Required, is.UUID)
		if err != nil {
			returnGetListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		addrs, err := uc.GetAddressesList(c, userID)
		if err != nil {
			returnGetListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetListResponse(c, http.StatusOK, true, "Addresses fetching is succeeded", addrs)
	}
}

func returnGetListResponse(c *gin.Context, statusCode int, ok bool, message string, addrs []sqlcdb.Address) {
	c.JSON(statusCode, &getListResponse{
		OK:        ok,
		Message:   message,
		Addresses: addrs,
	})
}
