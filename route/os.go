package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//https://github.com/akhenakh/statgo for metrics collection
func Os(api *gin.RouterGroup)  {
	route := api.Group("/os")
	route.GET("/list", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello ")
	})
}
