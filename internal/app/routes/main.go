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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, containers)
	})

	api.GET("/containers/:id", func(c *gin.Context) {
		containerId := c.Param("id")

		container, err := services.GetContainerById(ctx, cli, containerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, container)
	})

	api.GET("/images", func(c *gin.Context) {
		images, err := services.ListImages(ctx, cli)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, images)
	})

	api.GET("/images/:id", func(c *gin.Context) {
		imageId := c.Param("id")
		image, err := services.GetImageById(ctx, cli, imageId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, image)
	})
}
