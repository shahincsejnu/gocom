package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productdom "github.com/shahincsejnu/gocom/ecom/domain/product"
	"github.com/shahincsejnu/gocom/ecom/usecase/product"
)

type creationResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	ProductID string `json:"productID"`
}

func CreationHandler(uc *product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := new(productdom.CreationOptions)
		if err := c.ShouldBindJSON(&r); err != nil {
			returnCreationResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		productID, err := uc.CreateProduct(c, r)
		if err != nil {
			returnCreationResponse(c, http.StatusBadRequest, false, err.Error(), "")
			return
		}

		returnCreationResponse(c, http.StatusOK, true, "Product is created successfully", productID)
	}
}

func returnCreationResponse(c *gin.Context, statusCode int, ok bool, message, productID string) {
	c.JSON(statusCode, &creationResponse{
		OK:        ok,
		Message:   message,
		ProductID: productID,
	})
}
