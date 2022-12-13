package hyp

import (
	"context"
	"time"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type ContainerStatus struct {
	ID    string
	Name  string
	Image string
	//Network   NetworkInfo
	//Volume    VolumeInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (i *ContainerStatus) ConnectToNetwork(ctx context.Context, cli *client.Client, networkID string) {
	cli.NetworkConnect(ctx, networkID, i.ID, &network.EndpointSettings{})
}

/*
func (b *BundleInfo) StopBundle() {

}

func (b *BundleInfo) RestartBundle() {

}

func (b *BundleInfo) DeleteBundle() {

}
*/
