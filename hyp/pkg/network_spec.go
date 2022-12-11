package hyp

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type NetworkSpec struct {
	//ID   string
	Name string
}

func NewNetworSpec(name string) *NetworkSpec {
	return &NetworkSpec{
		Name: name,
	}
}

func (n *NetworkSpec) CreateNetwork(ctx context.Context, cli *client.Client) (*NetworkInfo, error) {

	resp, err := cli.NetworkCreate(ctx, n.Name, types.NetworkCreate{
		Attachable:     true,
		CheckDuplicate: true,
	})
	if err != nil {
		return &NetworkInfo{}, err
	}
	return &NetworkInfo{ID: resp.ID, Name: n.Name}, nil

}
