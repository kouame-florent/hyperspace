package hyp

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type ContainerObject struct {
	ObjectMeta
	Image    string
	Networks []NetworkObject
	//bindings between mount points and engine volumes
	//keys are container mount point e.g. /config, /var/lib/html and values are volume objects
	Volumes   map[string]VolumeObject
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *ContainerObject) StartContainer(ctx context.Context, cli *client.Client) (*ContainerObject, error) {
	fmt.Printf("---> stating container: %s\n", c.UID)
	if err := cli.ContainerStart(ctx, c.UID, types.ContainerStartOptions{}); err != nil {
		return &ContainerObject{}, err
	}

	//change container object status
	c.Status = Running

	//set update time
	c.UpdatedAt = time.Now()

	return c, nil
}

func (c *ContainerObject) ConnectToNetwork(ctx context.Context, cli *client.Client, networkID string) error {
	err := cli.NetworkConnect(ctx, networkID, c.UID, &network.EndpointSettings{})

	if err != nil {
		return err
	}

	return nil
}

/*
func (b *BundleInfo) StopBundle() {

}

func (b *BundleInfo) RestartBundle() {

}

func (b *BundleInfo) DeleteBundle() {

}
*/
