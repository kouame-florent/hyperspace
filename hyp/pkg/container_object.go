package hyp

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerObject struct {
	ObjectMeta
	Image string
	//Network   NetworkInfo
	//Volume    VolumeInfo
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

/*
func (i *ContainerObject) ConnectToNetwork(ctx context.Context, cli *client.Client, networkID string) {
	cli.NetworkConnect(ctx, networkID, i.ID, &network.EndpointSettings{})
}
*/

/*
func (b *BundleInfo) StopBundle() {

}

func (b *BundleInfo) RestartBundle() {

}

func (b *BundleInfo) DeleteBundle() {

}
*/
