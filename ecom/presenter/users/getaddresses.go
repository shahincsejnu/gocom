package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/users"
)

type getUserAddressesListResponse struct {
	OK        bool             `json:"ok"`
	Message   string           `json:"message"`
	Addresses []sqlcdb.Address `json:"addresses"`
}

func GetUserAddressesListHandler(uc *users.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		err := Validate(userID, Required, is.UUID)
		if err != nil {
			returnGetUserAddressesListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		addrs, err := uc.GetUserAddressesList(c, userID)
		if err != nil {
			returnGetUserAddressesListResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetUserAddressesListResponse(c, http.StatusOK, true, "User Addresses fetching is succeeded", addrs)
	}
}

func returnGetUserAddressesListResponse(c *gin.Context, statusCode int, ok bool, message string, addrs []sqlcdb.Address) {
	c.JSON(statusCode, &getUserAddressesListResponse{
		OK:        ok,
		Message:   message,
		Addresses: addrs,
	})
}
