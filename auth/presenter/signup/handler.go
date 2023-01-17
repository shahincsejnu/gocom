package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/auth/usecase/signup"
)

type request struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func Handler(uc *signup.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(request)

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, &response{
				OK:      false,
				Message: err.Error(),
			})
			return
		}

		err := uc.Do(c, r.Name, r.Email, r.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, &response{
				OK:      false,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, &response{
			OK:      true,
			Message: "Successfully Signedup",
		})
	}
}
