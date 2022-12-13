package hyp_test

import (
	"context"
	"testing"

	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

func TestCreateNetwork(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	expected := "test_network"

	spec := hyp.NetworkSpec{
		Name: expected,
	}

	inf, err := spec.CreateNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if inf.ID == "" {
		t.Fatal("Expected id, got empty string")
	}

	if inf.Name != expected {
		t.Fatalf("Expected %s, got %s", expected, inf.Name)
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

	spec := hyp.NetworkSpec{
		Name: netName,
	}

	cres, err := spec.CreateNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if cres.ID == "" {
		t.Fatal(err)
	}

	err = cres.RemoveNetwork(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

}
