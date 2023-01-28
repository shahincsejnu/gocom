package consoleauth

import (
	"net/http"

	sqlcdb "github.com/shahincsejnu/gocom/auth/infra/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/auth/usecase/consoleauth"
)

type request struct {
	Token string `json:"token" binding:"required"`
}

type response struct {
	OK      bool         `json:"ok"`
	Message string       `json:"message"`
	User    *sqlcdb.User `json:"user"`
}

func Handler(uc *consoleauth.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(request)

		if err := c.ShouldBindJSON(&r); err != nil {
			returnConsoleAuthResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		user, err := uc.Do(c, r.Token)
		if err != nil {
			returnConsoleAuthResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnConsoleAuthResponse(c, http.StatusOK, true, "Successfully verified", user)
	}
}

func returnConsoleAuthResponse(c *gin.Context, statusCode int, ok bool, message string, user *sqlcdb.User) {
	c.JSON(statusCode, &response{
		OK:      ok,
		Message: message,
		User:    user,
	})
}
