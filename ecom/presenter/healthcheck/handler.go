package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	OK          bool   `json:"ok"`
	ServiceName string `json:"serviceName"`
}

func Handler() gin.HandlerFunc {
	var staticResponse response
	staticResponse.OK = true
	staticResponse.ServiceName = "EcomService"

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &staticResponse)
	}
}
