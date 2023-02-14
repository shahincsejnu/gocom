package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	productdom "github.com/shahincsejnu/gocom/ecom/domain/product"
	"github.com/shahincsejnu/gocom/ecom/usecase/product"
)

type updateResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func UpdateHandler(uc *product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("productID")

		err := Validate(productID, Required, is.UUID)
		if err != nil {
			returnUpdateResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		r := new(productdom.UpdateOptions)
		if err := c.ShouldBindJSON(&r); err != nil {
			returnUpdateResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		err = uc.UpdateProduct(c, r, productID)
		if err != nil {
			returnUpdateResponse(c, http.StatusBadRequest, false, err.Error())
			return
		}

		returnUpdateResponse(c, http.StatusOK, true, "Product is updated successfully")
	}
}

func returnUpdateResponse(c *gin.Context, statusCode int, ok bool, message string) {
	c.JSON(statusCode, &updateResponse{
		OK:      ok,
		Message: message,
	})
}
