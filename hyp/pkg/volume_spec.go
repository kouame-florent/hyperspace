package hyp

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type VolumeSpec struct {
	SpecMeta

	//name of the volume object
	Name string

	//mount path on the host
	MountPoint string
}

func NewVolumeSpec(id, name string) *VolumeSpec {
	return &VolumeSpec{
		SpecMeta: SpecMeta{
			ID:        id,
			CreatedAt: time.Now(),
		},
		Name: name,
	}
}

func (v *VolumeSpec) CreateVolume(ctx context.Context, cli *client.Client) (VolumeStatus, error) {
	vol, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{
		Name: v.Name,
	})
	if err != nil {
		return VolumeStatus{}, err
	}

	volRsrc := VolumeStatus{
		StatusMeta: StatusMeta{
			CreatedAt: time.Now(),
		},

		Name:       vol.Name,
		MountPoint: vol.Mountpoint,
	}

	log.Printf("NAME: %s", volRsrc.Name)
	log.Printf("MOUNT POINT: %s", volRsrc.MountPoint)

	return volRsrc, nil
}
