package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(api *gin.RouterGroup)  {
	route := api.Group("/user")
	// This handler will match /route/john but will not match /route/ or /route
	route.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /route/john/ and also /route/john/send
	// If no other route match /route/john, it will redirect to /route/john/
	route.GET("/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	
}
