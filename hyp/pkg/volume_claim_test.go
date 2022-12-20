package hyp_test

import (
	"context"
	"testing"

	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestCreateVolumeClaim(t *testing.T) {
	fakeCli := fake.NewSimpleClientset()

	expectedName := "test_vol"
	ns := "test_ns"
	diksSize := "4Gi"

	volTpl := hyp.NewVolumeClaimTemplate(expectedName, ns, diksSize)

	vcl, err := volTpl.CreateVolumeClaim(context.Background(), fakeCli)
	if err != nil {
		t.Fatalf("not expected err, got: %v", vcl)
	}

	if vcl.Name != expectedName {
		t.Fatalf("expected: %s, got: %s", expectedName, vcl.Name)
	}

}
