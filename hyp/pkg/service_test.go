package hyp_test

import (
	"context"
	"testing"

	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestCreateService(t *testing.T) {
	fakeCli := fake.NewSimpleClientset()

	expectedSvc := "test_svc"
	ns := "test_ns"
	ports := map[int32]int{
		80: 8080,
	}

	svcTpl := hyp.NewServiceTemplate(expectedSvc, ns, ports)

	svcObj, err := svcTpl.CreateService(context.Background(), fakeCli)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if svcObj.Name != expectedSvc {
		t.Fatalf("expected service name %s, got %v", expectedSvc, svcObj.Name)
	}

}
