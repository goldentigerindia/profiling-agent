package route

import (
	"github.com/gin-gonic/gin"
	"github.com/goldentigerindia/profiling-agent/metrics/os"
	"net/http"
)


type UpTimeStat struct{

}
//https://github.com/mackerelio/go-osstat for metrics collection
func Os(api *gin.RouterGroup)  {
	route := api.Group("/os")
	route.GET("/loadavg", func(c *gin.Context) {
		loadAvgStat := os.GetOSLoadAvg()
		if loadAvgStat==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, loadAvgStat)
		}
	})
	route.GET("/uptime", func(c *gin.Context) {
		upTimeStat := os.GetOsUpTime()
		if upTimeStat==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, upTimeStat)
		}
	})
	route.GET("/swap", func(c *gin.Context) {
		swapStats := os.GetOSSwap()
		if swapStats ==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, swapStats)
		}
	})
	route.GET("/cpu", func(c *gin.Context) {
		cpuStat := os.GetOSCpuStat()
		if cpuStat ==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, cpuStat)
		}
	})
	route.GET("/memory", func(c *gin.Context) {
		memStat := os.GetOSMem()
		if memStat ==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, memStat)
		}
	})
	route.GET("/network", func(c *gin.Context) {
		netStat := os.GetOSNetwork()
		if netStat ==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, netStat)
		}
	})
	route.GET("/disk", func(c *gin.Context) {
		diskStat := os.GetOSDisk()
		if diskStat ==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, diskStat)
		}
	})
	route.GET("/process", func(c *gin.Context) {
		processStat := os.GetOSProcess()
		if processStat ==nil {
			c.JSON(http.StatusBadRequest, "Unable to receive data")
		}else {
			c.JSON(http.StatusOK, processStat)
		}
	})
}

