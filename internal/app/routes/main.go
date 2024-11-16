package routes

import (
	"context"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"

	svcContainer "docker-ui/internal/app/services/containers"
	svcImage "docker-ui/internal/app/services/images"
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
		containers, err := svcContainer.ListContainers(ctx, cli)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, containers)
	})

	api.GET("/containers/:id", func(c *gin.Context) {
		containerId := c.Param("id")

		container, err := svcContainer.GetContainerById(ctx, cli, containerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, container)
	})

	api.GET("/images", func(c *gin.Context) {
		images, err := svcImage.ListImages(ctx, cli)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, images)
	})

	api.GET("/images/:id", func(c *gin.Context) {
		imageId := c.Param("id")
		image, err := svcImage.GetImageById(ctx, cli, imageId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, image)
	})
}
