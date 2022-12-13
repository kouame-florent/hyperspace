package hyp_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

func TestContainer(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	env := []string{
		"NGINX_PORT=80",
	}

	ports := []string{
		"80/tcp",
	}

	image := "nginx:1.23.2-alpine"
	name := "alpine_container"

	contSpec := hyp.NewContainerSpec(image, name, env, []string{}, ports, map[string]string{})
	cinf, err := contSpec.CreateContainer(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if cinf.ID == "" {
		t.Fatal(err)
	}

	cmdContainerRemove(t, name)
}

func TestCreateContainerWithNetwork(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	netName := "alpine_network"

	netSpec := hyp.NewNetworkSpec(netName)

	netInfo, err := netSpec.CreateNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	networks := []string{netInfo.ID}

	volBindings := map[string]string{}

	env := []string{
		"NGINX_PORT=80",
	}

	ports := []string{
		"80/tcp",
	}

	image := "nginx:1.23.2-alpine"
	name := "alpine_container"

	contSpec := hyp.NewContainerSpec(image, name, env, networks, ports, volBindings)
	cinf, err := contSpec.CreateContainer(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if cinf.ID == "" {
		t.Fatal(err)
	}

	cmdNetworkRemove(t, "alpine_network")
	cmdContainerRemove(t, name)

}

func TestCreateContainerWithVolume(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	volName := "alpine_volume"

	volSpec := hyp.NewVolumeSpec(uuid.NewString(), volName)
	_, err = volSpec.CreateVolume(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	networks := []string{}

	env := []string{
		"NGINX_PORT=80",
	}

	ports := []string{
		"80/tcp",
	}

	image := "nginx:1.23.2-alpine"
	name := "alpine_container"

	volBindings := map[string]string{
		volSpec.Name: "/usr/share/nginx/html",
	}

	contSpec := hyp.NewContainerSpec(image, name, env, networks, ports, volBindings)
	cRsrc, err := contSpec.CreateContainer(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if cRsrc.ID == "" {
		t.Fatal(err)
	}

	cmdContainerRemove(t, name)
	cmdVolumeRemove(t, name)
}
