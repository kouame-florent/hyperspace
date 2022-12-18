package hyp

import (
	"strconv"

	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
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

// keys are mount points (target) in engine and values are volumes objects (source)
func dockerVolumesMounts(volumesBindings map[string]VolumeObject) []mount.Mount {
	mnt := []mount.Mount{}
	for k, v := range volumesBindings {
		m := mount.Mount{
			Type:   mount.TypeVolume,
			Source: v.Name,
			Target: k,
		}

		mnt = append(mnt, m)
	}

	return mnt
}

func dockerNetworkEndPoints(networks []NetworkObject) map[string]*network.EndpointSettings {
	endPoints := map[string]*network.EndpointSettings{}

	for _, o := range networks {
		endPoints[o.Name] = &network.EndpointSettings{NetworkID: o.Name}
	}

	return endPoints
}
