package hyp_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

func TestCreateContainer(t *testing.T) {
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

	contSpec := hyp.NewContainerSpec(uuid.NewString(), image, name, env, ports, map[string]hyp.VolumeSpec{}, []hyp.NetworkSpec{})
	cinf, err := contSpec.CreateContainer(ctx, cli, name)
	if err != nil {
		t.Fatal(err)
	}

	if cinf.UID == "" {
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

	nsp := hyp.NewNetworkSpec(uuid.NewString(), netName)

	/*
		_, err = nsp.CreateNetwork(ctx, cli, netName)
		if err != nil {
			t.Fatal(err)
		}
	*/

	netSpecs := []hyp.NetworkSpec{*nsp}

	volBindings := map[string]hyp.VolumeSpec{}

	env := []string{
		"NGINX_PORT=80",
	}

	ports := []string{
		"80/tcp",
	}

	image := "nginx:1.23.2-alpine"
	ctnName := "alpine_container"

	contSpec := hyp.NewContainerSpec(uuid.NewString(), image, ctnName, env, ports, volBindings, netSpecs)
	cobj, err := contSpec.CreateContainer(ctx, cli, ctnName)
	if err != nil {
		t.Fatal(err)
	}

	if cobj.UID == "" {
		t.Fatal(err)
	}

	cmdNetworkRemove(t, cobj.Networks[0].Name)
	cmdContainerRemove(t, ctnName)

}

func TestCreateContainerWithVolume(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	volName := "alpine_volume"

	volSpec := hyp.NewVolumeSpec(uuid.NewString(), volName)

	/*
		_, err = volSpec.CreateVolume(ctx, cli, volName)
		if err != nil {
			t.Fatal(err)
		}
	*/

	netSpecs := []hyp.NetworkSpec{}
	mountPt := "/usr/share/nginx/html"

	volBindings := map[string]hyp.VolumeSpec{
		mountPt: *volSpec,
	}

	env := []string{
		"NGINX_PORT=80",
	}

	ports := []string{
		"80/tcp",
	}

	image := "nginx:1.23.2-alpine"
	name := "alpine_container"

	contSpec := hyp.NewContainerSpec(uuid.NewString(), image, name, env, ports, volBindings, netSpecs)
	cobj, err := contSpec.CreateContainer(ctx, cli, name)
	if err != nil {
		t.Fatal(err)
	}

	if cobj.UID == "" {
		t.Fatal(err)
	}

	cmdContainerRemove(t, name)
	cmdVolumeRemove(t, cobj.Volumes[mountPt].Name)
}

func TestStartContainer(t *testing.T) {
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

	contSpec := hyp.NewContainerSpec(uuid.NewString(), image, name, env, ports, map[string]hyp.VolumeSpec{}, []hyp.NetworkSpec{})
	cobj, err := contSpec.CreateContainer(ctx, cli, name)
	if err != nil {
		t.Fatal(err)
	}

	if cobj.UID == "" {
		t.Fatal(err)
	}

	robj, err := cobj.StartContainer(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if robj.UID == "" {
		t.Fatal(err)
	}

	cmdContainerStop(t, name)
	cmdContainerRemove(t, name)

}
