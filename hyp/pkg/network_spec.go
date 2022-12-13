package hyp

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// network specification to create a network resource in engine
type NetworkSpec struct {
	SpecMeta
	Name string
}

func NewNetworkSpec(name string) *NetworkSpec {
	return &NetworkSpec{
		Name: name,
	}
}

func (n *NetworkSpec) CreateNetwork(ctx context.Context, cli *client.Client) (*NetworkStatus, error) {

	resp, err := cli.NetworkCreate(ctx, n.Name, types.NetworkCreate{
		Attachable:     true,
		CheckDuplicate: true,
	})
	if err != nil {
		return &NetworkStatus{}, err
	}

	return &NetworkStatus{
		StatusMeta: StatusMeta{
			ID: resp.ID,
		},
		Name: n.Name}, nil

}
