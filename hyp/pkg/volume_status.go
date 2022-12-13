package hyp

import (
	"context"

	"github.com/docker/docker/client"
)

type VolumeStatus struct {
	StatusMeta

	//name of the volume object
	Name string

	//mount path on the host
	MountPoint string
}

func (v *VolumeStatus) RemoveVolume(ctx context.Context, cli *client.Client) error {
	err := cli.VolumeRemove(ctx, v.Name, true)
	if err != nil {
		return err
	}

	return nil
}
