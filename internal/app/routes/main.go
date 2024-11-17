package routes

import (
	"context"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"

	"docker-ui/internal/app/models"
	svc "docker-ui/internal/app/services"
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
		containers, err := svc.ListContainers(ctx, cli)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, containers)
	})

	api.GET("/containers/:id", func(c *gin.Context) {
		containerId := c.Param("id")

		container, err := svc.GetContainerById(ctx, cli, containerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, container)
	})

	api.GET("/images", func(c *gin.Context) {
		images, err := svc.ListImages(ctx, cli)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, images)
	})

	api.GET("/images/:id", func(c *gin.Context) {
		imageId := c.Param("id")
		image, err := svc.GetImageById(ctx, cli, imageId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, image)
	})

	api.POST("/pull-image", func(c *gin.Context) {
		var imageInfo models.ImageInfo
		if err := c.ShouldBindJSON(&imageInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
			return
		}

		if err := svc.PullImage(ctx, cli, imageInfo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image pulled successfully"})
	})
}
