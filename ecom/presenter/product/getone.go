package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/is"
	. "github.com/go-ozzo/ozzo-validation/v4"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/product"
)

type getOneResponse struct {
	OK      bool            `json:"ok"`
	Message string          `json:"message"`
	Product *sqlcdb.Product `json:"product"`
}

func GetOneHandler(uc *product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("productID")

		err := Validate(productID, Required, is.UUID)
		if err != nil {
			returnGetOneResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		prd, err := uc.GetOneProduct(c, productID)
		if err != nil {
			returnGetOneResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnGetOneResponse(c, http.StatusOK, true, "Product fetching is succeeded", prd)
	}
}

func returnGetOneResponse(c *gin.Context, statusCode int, ok bool, message string, prd *sqlcdb.Product) {
	c.JSON(statusCode, &getOneResponse{
		OK:      ok,
		Message: message,
		Product: prd,
	})
}
