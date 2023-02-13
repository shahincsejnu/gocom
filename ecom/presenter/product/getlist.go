package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sqlcdb "github.com/shahincsejnu/gocom/ecom/infra/sqlc"
	"github.com/shahincsejnu/gocom/ecom/usecase/product"
)

type response struct {
	OK       bool             `json:"ok"`
	Message  string           `json:"message"`
	Products []sqlcdb.Product `json:"products"`
}

func GetListHandler(uc *product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := uc.GetAllProducts(c)
		if err != nil {
			returnResponse(c, http.StatusBadRequest, false, err.Error(), nil)
			return
		}

		returnResponse(c, http.StatusOK, true, "Products fetching is succeeded", products)
	}
}

func returnResponse(c *gin.Context, statusCode int, ok bool, message string, products []sqlcdb.Product) {
	c.JSON(statusCode, &response{
		OK:       ok,
		Message:  message,
		Products: products,
	})
}
