package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
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

func (d *DockerService) GetContainers() ([]types.Container, error) {
	return d.cli.ContainerList(context.Background(), container.ListOptions{})
}

func (d *DockerService) GetSystemInformation(ctx context.Context, uri string) (*system.Info, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	info, err := cli.Info(ctx)
	return &info, err
}
