package services

import (
	"context"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"

	"github.com/docker/docker/client"
)

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
