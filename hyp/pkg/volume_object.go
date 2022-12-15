package hyp

import (
	"context"

	"github.com/docker/docker/client"
)

type VolumeObject struct {
	ObjectMeta

	//mount path on the host
	MountPoint string
}

func (v *VolumeObject) RemoveVolume(ctx context.Context, cli *client.Client) error {
	err := cli.VolumeRemove(ctx, v.Name, true)
	if err != nil {
		return err
	}

	return nil
}
