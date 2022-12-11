package hyp

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// Informations about a network object in the container engine
type NetworkInfo struct {
	ID   string
	Name string
}

func (n *NetworkInfo) Network(ctx context.Context, cli *client.Client) (types.NetworkResource, error) {
	res, err := cli.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "Name", Value: n.Name}, filters.KeyValuePair{Key: "ID", Value: n.ID}),
	})
	if err != nil {
		return types.NetworkResource{}, err
	}

	if len(res) > 0 {
		return types.NetworkResource{}, fmt.Errorf("more than one network with ID: %s", n.ID)
	}

	return res[0], nil
}

func (n *NetworkInfo) RemoveNetwork(ctx context.Context, cli *client.Client) error {
	err := cli.NetworkRemove(ctx, n.ID)
	if err != nil {
		return err
	}

	return nil
}
