package services

import (
	"context"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"

	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

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
