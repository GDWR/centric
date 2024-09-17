package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type DockerService struct {
	cli *client.Client
}

func NewDockerService() (*DockerService, error) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.46"))
	if err != nil {
		return nil, err
	}

	return &DockerService{cli: cli}, nil
}

func (d *DockerService) GetContainers(ctx context.Context, uri string) ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli.ContainerList(ctx, container.ListOptions{
		All: true,
	})
}

func (d *DockerService) GetImages(ctx context.Context, uri string) ([]image.Summary, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli.ImageList(ctx, image.ListOptions{
		All: true,
	})
}

func (d *DockerService) GetVolumes(ctx context.Context, uri string) ([]*volume.Volume, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	volumeList, err := cli.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return nil, err
	}

	return volumeList.Volumes, nil
}

func (d *DockerService) GetSystemInformation(ctx context.Context, uri string) (*system.Info, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	info, err := cli.Info(ctx)
	return &info, err
}

func (d *DockerService) GetNetworks(ctx context.Context, uri string) ([]network.Inspect, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli.NetworkList(ctx, network.ListOptions{})
}
