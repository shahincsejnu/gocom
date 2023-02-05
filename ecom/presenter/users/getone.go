package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/users"
)

type response struct {
	OK      bool         `json:"ok"`
	Message string       `json:"message"`
	User    *sqlcdb.User `json:"user"`
}

func GetOneHandler(uc *users.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		err := Validate(userID, Required, is.UUID)
		if err != nil {
			returnResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		user, err := uc.GetUserById(c, userID)
		if err != nil {
			returnResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnResponse(c, http.StatusOK, true, "User is fetched", user)
	}
}

func returnResponse(c *gin.Context, statusCode int, ok bool, message string, user *sqlcdb.User) {
	c.JSON(statusCode, &response{
		OK:      ok,
		Message: message,
		User:    user,
	})
}
