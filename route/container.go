package route

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Container(api *gin.RouterGroup)  {
	route := api.Group("/container")
	// This handler will match /route/john but will not match /route/ or /route
	route.GET("/list", func(c *gin.Context) {
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		for _, container := range containers {
			fmt.Println(container.ID)
		}
		c.JSON(http.StatusOK,containers)
	})
	route.GET("/logs/:containerId", func(c *gin.Context) {
		ctx := context.Background()
		containerId := c.Param("containerId")
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		options := types.ContainerLogsOptions{ShowStdout: true}
		// Replace this ID with a container that really exists
		out, err := cli.ContainerLogs(ctx, containerId, options)
		if err != nil {
			panic(err)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(out)
		log := buf.String()
		c.String(http.StatusOK, log)
	})
}