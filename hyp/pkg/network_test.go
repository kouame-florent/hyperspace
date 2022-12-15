package hyp_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

func TestCreateNetwork(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	expected := "test_network"

	spec := hyp.NewNetworkSpec(uuid.NewString(), expected)

	nobj, err := spec.CreateNetwork(ctx, cli, expected)
	if err != nil {
		t.Fatal(err)
	}

	if nobj.UID == "" {
		t.Fatal("Expected id, got empty string")
	}

	if nobj.Name != expected {
		t.Fatalf("Expected %s, got %s", expected, nobj.Name)
	}

	cmdNetworkRemove(t, expected)

}

func TestRemoveNetwork(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	netName := "test_network"

	spec := hyp.NewNetworkSpec(uuid.NewString(), netName)

	cres, err := spec.CreateNetwork(ctx, cli, netName)
	if err != nil {
		t.Fatal(err)
	}

	if cres.UID == "" {
		t.Fatal(err)
	}

	err = cres.RemoveNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

}
