package hyp

import (
	"context"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// network resources in the container engine
type NetworkStatus struct {
	StatusMeta
	Name string
}

/*
	func (n *NetworkRsrc) Network(ctx context.Context, cli *client.Client) (types.NetworkResource, error) {
		res, err := cli.NetworkList(ctx, types.NetworkListOptions{
			Filters: filters.NewArgs(filters.KeyValuePair{Key: "Name", Value: n.Name}, filters.KeyValuePair{Key: "ID", Value: n.UID}),
		})
		if err != nil {
			return types.NetworkResource{}, err
		}

		if len(res) > 0 {
			return types.NetworkResource{}, fmt.Errorf("more than one network with ID: %s", n.ID)
		}

		return res[0], nil
	}
*/
func (n *NetworkStatus) RemoveNetwork(ctx context.Context, cli *client.Client) error {
	err := cli.NetworkRemove(ctx, n.ID)
	if err != nil {
		return err
	}

	return nil
}

func (n *NetworkStatus) Connect(ctx context.Context, cli *client.Client, containerID string) error {
	err := cli.NetworkConnect(ctx, n.ID, containerID, &network.EndpointSettings{})
	if err != nil {
		return err
	}

	return nil
}
