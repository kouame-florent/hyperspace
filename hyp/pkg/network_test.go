package hyp_test

import (
	"context"
	"os/exec"
	"testing"

	"github.com/docker/docker/client"
	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

/*
func TestMain(m *testing.M) {



}
*/

func createClient(t *testing.T) (*client.Client, error) {

	client.NewClientWithOpts()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	t.Cleanup(func() {
		cli.Close()
	})

	return cli, nil
}

func createNetworkCleanUp(t *testing.T, name string) {
	err := runCmd("docker", []string{"network", "rm", name})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateNetwork(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	netName := "test_network"

	spec := hyp.NetworkSpec{
		Name: netName,
	}

	info, err := spec.CreateNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if info.ID == "" {
		t.Fatal(err)
	}

	createNetworkCleanUp(t, netName)
}

func TestRemoveNetwork(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	netName := "test_network"

	spec := hyp.NetworkSpec{
		Name: netName,
	}

	cinfo, err := spec.CreateNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if cinfo.ID == "" {
		t.Fatal(err)
	}

	err = cinfo.RemoveNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

}

func runCmd(command string, args []string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()

}
