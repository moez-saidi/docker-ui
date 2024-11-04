package routes

import (
	"context"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"

	"docker-ui/internal/app/services"
)

const apiVersion string = "/api/v1"

func RegisterRoutes(router *gin.Engine) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	api := router.Group(apiVersion)
	api.GET("/containers", func(c *gin.Context) {
		containers, err := services.ListContainers(ctx, cli)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, containers)
	})

	api.GET("/images", func(c *gin.Context) {
		images, err := services.ListImages(ctx, cli)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, images)
	})
}
