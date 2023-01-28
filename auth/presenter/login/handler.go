package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/auth/usecase/login"
)

type request struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Handler(uc *login.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(request)

		if err := c.ShouldBindJSON(&r); err != nil {
			returnLoginResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		token, err := uc.Do(c, r.Name, r.Password)
		if err != nil {
			returnLoginResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		returnLoginResponse(c, http.StatusOK, true, "Successfully Logged In!", token)
	}
}

func returnLoginResponse(c *gin.Context, statusCode int, ok bool, message, token string) {
	c.JSON(statusCode, &response{
		OK:      ok,
		Message: message,
		Token:   token,
	})
}
