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

func GetImageById(ctx context.Context, cli *client.Client, imageId string) (types.ImageInspect, error) {
	image, _, err := cli.ImageInspectWithRaw(ctx, imageId)
	if err != nil {
		return types.ImageInspect{}, err
	}

	return image, nil
}

func ListContainers(ctx context.Context, cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func GetContainerById(ctx context.Context, cli *client.Client, containerId string) (types.ContainerJSON, error) {
	container, err := cli.ContainerInspect(ctx, containerId)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return container, nil
}
