package hyp

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// network resources in the container engine
type NetworkObject struct {
	ObjectMeta
}

type NetworkFilterKey string
type NetworkFilterValue string

const (
	NetFilterKey   NetworkFilterKey   = "Name"
	NetFilterValue NetworkFilterValue = "UID"
)

func (n *NetworkObject) RemoveNetwork(ctx context.Context, cli *client.Client) error {
	err := cli.NetworkRemove(ctx, n.UID)
	if err != nil {
		return err
	}

	return nil
}

func (n *NetworkObject) Connect(ctx context.Context, cli *client.Client, containerID string) error {
	err := cli.NetworkConnect(ctx, n.UID, containerID, &network.EndpointSettings{})
	if err != nil {
		return err
	}

	return nil
}

func (n *NetworkObject) Exist(ctx context.Context, cli *client.Client) (bool, error) {

	res, err := cli.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(
			filters.KeyValuePair{
				Key:   string(NetFilterKey),
				Value: n.Name,
			},
		),
	})
	if err != nil {
		return false, err
	}

	if len(res) > 0 {
		return true, nil
	}

	return false, nil

}
