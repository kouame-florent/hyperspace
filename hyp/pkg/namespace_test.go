package hyp_test

import (
	"context"
	"testing"

	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestCreateNamespace(t *testing.T) {
	fakeCli := fake.NewSimpleClientset()

	expected := "test_namesapce"

	nsObj, err := hyp.CreateNamespace(context.Background(), fakeCli, expected)
	if err != nil {
		t.Fatalf("error when creating namespace: %v", err)
	}

	gotNs, err := fakeCli.CoreV1().Namespaces().Get(context.Background(), expected, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("error when getting namespace: %v", err)
	}

	if nsObj.Name != gotNs.Name {
		t.Fatalf("expected name %s, but got %s ", expected, gotNs.Name)
	}

}

func TestNamespace(t *testing.T) {
	expected := "test_namesapce"

	fakeCli := fake.NewSimpleClientset(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: expected,
		},
	})

	nsObj, err := hyp.Namespace(context.Background(), fakeCli, expected)
	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if nsObj.Name != expected {
		t.Fatalf("expected %s, got %s", expected, nsObj.Name)
	}
}
