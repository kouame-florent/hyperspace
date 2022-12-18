package hyp

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// network specification to create a network resource in engine
type NetworkSpec struct {
	SpecMeta
}

func NewNetworkSpec(id, title string) *NetworkSpec {
	return &NetworkSpec{
		SpecMeta: SpecMeta{
			ID:        id,
			Tag:       title,
			CreatedAt: time.Now(),
		},
	}
}

func (n *NetworkSpec) CreateNetwork(ctx context.Context, cli *client.Client, name string) (*NetworkObject, error) {

	resp, err := cli.NetworkCreate(ctx, name, types.NetworkCreate{
		Attachable:     true,
		CheckDuplicate: true,
	})
	if err != nil {
		return &NetworkObject{}, err
	}

	return &NetworkObject{
		ObjectMeta: ObjectMeta{
			UID:    resp.ID,
			Name:   name,
			SpecID: n.ID,
		},
	}, nil

}

// get network by name
func Network(ctx context.Context, cli *client.Client, name string) (*NetworkObject, error) {
	res, err := cli.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: string(NetFilterKey), Value: name}),
	})
	if err != nil {
		return &NetworkObject{}, nil
	}

	if len(res) > 0 {
		return &NetworkObject{}, fmt.Errorf("more than one network with name: %s", name)
	}

	return &NetworkObject{
		ObjectMeta: ObjectMeta{
			UID:       res[0].ID,
			Name:      res[0].Name,
			CreatedAt: time.Now(),
		},
	}, nil
}
