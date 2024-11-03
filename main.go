package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"

	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

const Port = ":10000"

func ListImages(ctx context.Context, cli *client.Client) ([]imagetypes.Summary, error) {
	images, err := cli.ImageList(ctx, imagetypes.ListOptions{})
	if err != nil {
		return nil, err
	}

	return images, nil
}

func ListContainers(ctx context.Context, cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func main() {
	router := gin.Default()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	router.GET("/containers", func(c *gin.Context) {
		containers, err := ListContainers(ctx, cli)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, containers)
	})

	router.GET("/images", func(c *gin.Context) {
		images, err := ListImages(ctx, cli)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, images)
	})
	router.Run(Port)

}
