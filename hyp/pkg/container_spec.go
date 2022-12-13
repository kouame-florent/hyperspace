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
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// values for creating container on the fly
type ContainerSpec struct {
	ID string

	//container name
	Name string

	//container image uri volume
	Image string

	//container network ids
	Networks []string

	//container environment viriables
	Environment []string

	//bindings between volumes and mount points.
	// keys are volume name and values are mount point in container e.g. test_volume => /etc
	VolumeBinding map[string]string

	//container exposed ports
	Ports []string
}

func NewContainerSpec(image, name string, environment, networks, ports []string, volumeBinding map[string]string) *ContainerSpec {
	//init the generator or you will always get the same value from GetRandomName
	rand.Seed(time.Now().UnixNano())
	return &ContainerSpec{
		Name:          name, // namesgenerator.GetRandomName(10),
		Image:         image,
		Environment:   environment,
		VolumeBinding: volumeBinding,
		Networks:      networks,
		Ports:         ports,
	}
}

func (i *ContainerSpec) CreateContainer(ctx context.Context, cli *client.Client) (ContainerStatus, error) {

	createBody, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        i.Image,
		Env:          i.Environment,
		ExposedPorts: dockerExposedPorts(i.Ports),
	}, &container.HostConfig{
		Mounts:       dockerVolumes(i.VolumeBinding),
		PortBindings: dockerPortsBinding(i.Ports),
	}, &network.NetworkingConfig{
		EndpointsConfig: dockerNetwork(i.Networks),
	}, &v1.Platform{}, i.Name)

	if err != nil {
		return ContainerStatus{}, err
	}

	fmt.Printf("---> created container id: %s\n", createBody.ID)

	return ContainerStatus{
		ID:    createBody.ID,
		Name:  i.Name,
		Image: i.Image,
	}, nil

	//return createBody.ID, nil

}

func (i *ContainerSpec) StartContainer(ctx context.Context, cli *client.Client, id string) (ContainerStatus, error) {
	fmt.Printf("---> stating container: %s\n", id)
	if err := cli.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
		return ContainerStatus{}, err
	}

	return ContainerStatus{
		ID:        id,
		Name:      i.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
