package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	containertypes "github.com/docker/docker/api/types/container"

	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func ListImages(ctx context.Context, cli *client.Client) {
	images, err := cli.ImageList(ctx, imagetypes.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		imageBytes, _ := json.Marshal(image)
		fmt.Println(string(imageBytes))
	}
}

func ListContainers(ctx context.Context, cli *client.Client) {
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		containerBytes, _ := json.Marshal(container)
		fmt.Println(string(containerBytes))
	}
}

func main() {
	gin.Default()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ListImages(ctx, cli)
	ListContainers(ctx, cli)

}
