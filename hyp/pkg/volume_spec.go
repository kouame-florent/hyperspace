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
}

func NewVolumeSpec(id, tag string) *VolumeSpec {
	return &VolumeSpec{
		SpecMeta: SpecMeta{
			ID:        id,
			Tag:       tag,
			CreatedAt: time.Now(),
		},
	}
}

func (v *VolumeSpec) CreateVolume(ctx context.Context, cli *client.Client, name string) (VolumeObject, error) {
	vol, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{
		Name: name,
	})
	if err != nil {
		return VolumeObject{}, err
	}

	volRsrc := VolumeObject{
		ObjectMeta: ObjectMeta{
			CreatedAt: time.Now(),
			Name:      vol.Name,
		},

		MountPoint: vol.Mountpoint,
	}

	log.Printf("NAME: %s", volRsrc.Name)
	log.Printf("MOUNT POINT: %s", volRsrc.MountPoint)

	return volRsrc, nil
}
