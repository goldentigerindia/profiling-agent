package route

import (
	"github.com/gin-gonic/gin"
	"github.com/goldentigerindia/profiling-agent/route/profile"
)

func Profile(api *gin.RouterGroup) {
	route := api.Group("/profile")
	profile.Java(route)
	profile.NodeJS(route)
	profile.GoLang(route)
}