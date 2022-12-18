package hyp

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// values for creating container on the fly
type ContainerSpec struct {
	SpecMeta

	//container image uri volume
	Image string

	//container environment variables
	Environment []string

	//container network
	//Networks []NetworkSpec

	//bindings between mount points and engine volumes
	//keys are container mount point e.g. /config, /var/lib/html and values are volume spec
	MountPoints []string

	//container exposed ports
	Ports []string
}

// id must be generate by docker namesgenerator.GetRandomName
func NewContainerSpec(id, image, title string, environment, ports []string, mountPoints []string, networks []NetworkSpec) *ContainerSpec {

	//init the generator or you will always get the same value from GetRandomName
	rand.Seed(time.Now().UnixNano())

	return &ContainerSpec{
		SpecMeta: SpecMeta{
			ID:        id,
			Tag:       title,
			CreatedAt: time.Now(),
		},
		Image:       image,
		Environment: environment,
		MountPoints: mountPoints,
		//Networks:        networks,
		Ports: ports,
	}
}

func (c *ContainerSpec) CreateContainer(ctx context.Context, cli *client.Client, name string) (ContainerObject, error) {

	/*
		//create networks before creating container
		netObjs := []NetworkObject{}
		for _, net := range c.Networks {
			netName := uuid.NewString()
			netObj, err := net.CreateNetwork(ctx, cli, netName)
			if err != nil {
				return ContainerObject{}, err
			}
			netObjs = append(netObjs, *netObj)

		}
	*/

	//create volumes  before creating container
	volObjs := map[string]VolumeObject{}
	for k, v := range c.MountPoints {
		volName := uuid.NewString()
		volObj, err := v.CreateVolume(ctx, cli, volName)
		if err != nil {
			return ContainerObject{}, err
		}

		volObjs[k] = volObj
	}

	createBody, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        c.Image,
		Env:          c.Environment,
		ExposedPorts: dockerExposedPorts(c.Ports),
	}, &container.HostConfig{
		Mounts:       dockerVolumesMounts(volObjs),
		PortBindings: dockerPortsBinding(c.Ports),
	}, &network.NetworkingConfig{
		//EndpointsConfig: dockerNetworkEndPoints(netObjs),
	}, &v1.Platform{}, name)

	if err != nil {
		return ContainerObject{}, err
	}

	fmt.Printf("---> created container id: %s\n", createBody.ID)

	return ContainerObject{
		ObjectMeta: ObjectMeta{
			UID:    createBody.ID,
			Name:   name,
			SpecID: c.ID,
		},

		Image: c.Image,
		//Networks: netObjs,
		//Volumes: volObjs,
	}, nil

}

type CreateContainerOpts struct {
	ContainerName string
}
