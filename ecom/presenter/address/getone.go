package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/address"
)

type getOneResponse struct {
	OK      bool            `json:"ok"`
	Message string          `json:"message"`
	Address *sqlcdb.Address `json:"address"`
}

func GetOneHandler(uc *address.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		addressID := c.Param("addressID")

		err := Validate(addressID, Required, is.UUID)
		if err != nil {
			returnGetOneResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		addr, err := uc.GetOneAddress(c, addressID)
		if err != nil {
			returnGetOneResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetOneResponse(c, http.StatusOK, true, "Address fetching is succeeded", addr)
	}
}

func returnGetOneResponse(c *gin.Context, statusCode int, ok bool, message string, addr *sqlcdb.Address) {
	c.JSON(statusCode, &getOneResponse{
		OK:      ok,
		Message: message,
		Address: addr,
	})
}
