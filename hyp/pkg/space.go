package hyp

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// container id returned by docker engine
type ContainerID = string

// user context to manage container
type Space struct {
	ID         string
	OwnerID    string
	Containers []ContainerConfig
}

// values for creating container on the fly
type ContainerConfig struct {
	// container id returned engine
	ID ContainerID

	//container name
	Name string

	//container network name
	Network string

	//container volume
	Image string

	//container environment viriables
	Environment []string

	//container mount volumes
	Volumes []string

	//container exposed ports
	Ports []string
}

// create a new space
func NewSpace(ownerID, image string, volume, ports []string) *Space {
	return &Space{
		ID:      uuid.New().String(),
		OwnerID: ownerID,
		Containers: []ContainerConfig{
			{
				Name:  uuid.New().String(),
				Image: image,
				Environment: []string{
					"PUID=1000",
					"PGID=1000",
					"TZ=Africa/Abidjan",
					"PASSWORD=code",
					"DEFAULT_WORKSPACE=/config/workspace",
				},
				Volumes: volume,
				Ports:   ports,
			},
		},
	}
}

func NewContainerConfig(name, image string, environment, volumes, network, ports []string) *ContainerConfig {
	return &ContainerConfig{
		Name:        uuid.New().String(),
		Image:       image,
		Environment: environment,
		Volumes:     volumes,
		Ports:       ports,
	}
}

// convert ports slice to docker nat.PortSet
func dockerExposedPorts(ports []string) nat.PortSet {
	ps := nat.PortSet{}
	for _, p := range ports {
		ps[nat.Port(p)] = struct{}{}
	}

	return ps

}

// convert ports slice to docker nat.PortMap
func dockerPortsBinding(ports []string) nat.PortMap {
	pm := nat.PortMap{}
	//retrieve available port and bind some of them
	port1 := "8443"
	port2 := "8282"

	for _, p := range ports {
		pm[nat.Port(p)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: port1,
			},
			{
				HostIP:   "0.0.0.0",
				HostPort: port2,
			},
		}
	}

	return pm

}

func (s *Space) CreateContainer(ctx context.Context, cli *client.Client) error {

	for _, ctn := range s.Containers {
		createBody, err := cli.ContainerCreate(ctx, &container.Config{
			Image:        ctn.Image,
			Env:          ctn.Environment,
			ExposedPorts: dockerExposedPorts(ctn.Ports),
		}, &container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeVolume,
					Source: ctn.Volume,
					Target: "/config",
				},
			},
			PortBindings: dockerPortsBinding(ctn.Ports),
		}, &network.NetworkingConfig{}, &v1.Platform{},
			"code-server")
		if err != nil {
			panic(err)
		}

		fmt.Printf("---> created container id: %s", createBody.ID)
		//set space id with created container id
		ctn.ID = createBody.ID

		if err := cli.ContainerStart(ctx, createBody.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

	}

	return nil
}
