package hyp

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/namesgenerator"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// values for creating container on the fly
type ContainerSpec struct {
	ID string

	//container name
	Name string

	//container image uri volume
	Image string

	//container network name
	Networks []string

	//container environment viriables
	Environment []string

	//container mount volumes
	Volumes []string

	//container exposed ports
	Ports []string
}

func NewContainerSpec(image string, environment, volumes, network, ports []string) *ContainerSpec {
	//init the generator or you will always get the same value from GetRandomName
	rand.Seed(time.Now().UnixNano())
	return &ContainerSpec{
		Name:        namesgenerator.GetRandomName(10),
		Image:       image,
		Environment: environment,
		Volumes:     volumes,
		Networks:    network,
		Ports:       ports,
	}
}

func (i *ContainerSpec) CreateContainer(ctx context.Context, cli *client.Client) (string, error) {

	createBody, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        i.Image,
		Env:          i.Environment,
		ExposedPorts: dockerExposedPorts(i.Ports),
	}, &container.HostConfig{
		Mounts:       dockerVolumes(i.Volumes),
		PortBindings: dockerPortsBinding(i.Ports),
	}, &network.NetworkingConfig{
		EndpointsConfig: dockerNetwork(i.Networks),
	}, &v1.Platform{}, i.Name)

	if err != nil {
		return "", err
	}

	fmt.Printf("---> created container id: %s\n", createBody.ID)

	return createBody.ID, nil

}

func (i *ContainerSpec) StartContainer(ctx context.Context, cli *client.Client, id string) (ContainerInfo, error) {
	fmt.Printf("---> stating container: %s\n", id)
	if err := cli.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
		return ContainerInfo{}, err
	}

	return ContainerInfo{
		ID:        id,
		Name:      i.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
