package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func main() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	createBody, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "127.0.0.1:5000/code-java-env:v0.0.8",
		Env: []string{
			"PUID=1000",
			"PGID=1000",
			"TZ=Africa/Abidjan",
			"PASSWORD=code",
			"DEFAULT_WORKSPACE=/config/workspace",
		},
		ExposedPorts: nat.PortSet{
			"8443/tcp": struct{}{},
			"8080/tcp": struct{}{},
		},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: "config",
				Target: "/config",
			},
		},
		PortBindings: nat.PortMap{
			"8443/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8443",
				},
			},
			"8080/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8282",
				},
			},
		},
	}, &network.NetworkingConfig{}, &v1.Platform{}, "code-server")
	if err != nil {
		panic(err)
	}

	fmt.Printf("---> created container id: %s", createBody.ID)

	if err := cli.ContainerStart(ctx, createBody.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
