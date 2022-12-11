package hyp

import (
	"strconv"

	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// convert ports slice to docker nat.PortSet
func dockerExposedPorts(ports []string) nat.PortSet {
	ps := nat.PortSet{}
	for _, p := range ports {
		ps[nat.Port(p)] = struct{}{}
	}

	return ps

}

// convert ports slice to docker nat.PortMap
func dockerPortsBinding(ports []string) nat.PortMap {
	pm := nat.PortMap{}
	//retrieve available ports from host and bind some of them

	for _, p := range ports {
		free, err := GetFreePort()
		if err != nil {
			panic(err)
		}
		pm[nat.Port(p)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: strconv.Itoa(free),
			},
		}

	}

	return pm

}

// build docker volumes mount with targets names
func dockerVolumes(volumes []string) []mount.Mount {
	mnt := []mount.Mount{}
	for _, v := range volumes {
		m := mount.Mount{
			Type:   mount.TypeVolume,
			Source: dockerVolumeSource(),
			Target: v,
		}
		mnt = append(mnt, m)
	}

	return mnt
}

// produce docker volume source name
func dockerVolumeSource() string {
	return uuid.NewString()
}

func dockerNetwork(networks []string) map[string]*network.EndpointSettings {
	endPoints := map[string]*network.EndpointSettings{}

	for _, n := range networks {
		endPoints[n] = &network.EndpointSettings{NetworkID: n}
	}

	return endPoints
}
