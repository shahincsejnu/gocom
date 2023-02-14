package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shahincsejnu/gocom/ecom/usecase/product"
)

type deleteOneResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func DeleteOneHandler(uc *product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("productID")

		err := Validate(productID, Required, is.UUID)
		if err != nil {
			returnDeleteOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		err = uc.DeleteOneProduct(c, productID)
		if err != nil {
			returnDeleteOneResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnDeleteOneResponse(c, http.StatusOK, true, "Product deletion is succeeded")
	}
}

func returnDeleteOneResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &deleteOneResponse{
		OK:      ok,
		Message: message,
	})
}
