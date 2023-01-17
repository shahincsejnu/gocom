package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahincsejnu/gocom/auth/usecase/signup"
	"golang.org/x/crypto/bcrypt"
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
			returnSignupResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		hashPass, err := hashifyPass(r.Password)
		if err != nil {
			returnSignupResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}
		err = uc.Do(c, r.Name, r.Email, hashPass)
		if err != nil {
			returnSignupResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnSignupResponse(c, http.StatusOK, true, "Successfully Signed Up!")
	}
}

func hashifyPass(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func returnSignupResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &response{
		OK:      ok,
		Message: message,
	})
}
