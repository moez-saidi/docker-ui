package services

import (
	"context"
	"io"
	"os"

	"docker-ui/internal/app/models"

	"github.com/docker/docker/api/types"
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

func PullImage(ctx context.Context, cli *client.Client, imageInfo models.ImageInfo) error {

	reader, err := cli.ImagePull(ctx, imageInfo.GetFullRepoURL(), imagetypes.PullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	return err
}
