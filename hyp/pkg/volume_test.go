package hyp_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

func TestCreateVolume(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	expected := "test_volume"

	spec := hyp.NewVolumeSpec(uuid.NewString(), expected)
	rsrc, err := spec.CreateVolume(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	if rsrc.Name != expected {
		t.Fatalf("Expected %s, got %s", expected, rsrc.Name)
	}

	cmdVolumeRemove(t, expected)

}

func TestRemoveVolume(t *testing.T) {
	ctx := context.Background()
	cli, err := createClient(t)
	if err != nil {
		t.Fatal(err)
	}

	expected := "test_volume"

	spec := hyp.NewVolumeSpec(uuid.NewString(), expected)
	cinf, err := spec.CreateVolume(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

	err = cinf.RemoveVolume(ctx, cli)
	if err != nil {
		t.Fatal(err)
	}

}
