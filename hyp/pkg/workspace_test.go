package hyp_test

import (
	"context"
	"testing"

	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestDeployWorkspace(t *testing.T) {
	fakeCli := fake.NewSimpleClientset()

	expectedNS := "test_ns"
	workspaceName := "quarkus"
	workspaceDesc := "quarkus microprofile environment"
	workspaceImage := "nginx:1.12"
	workspacePorts := []int32{8443, 8080}

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: expectedNS}}

	cteatedNS, err := fakeCli.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("not expected error, got: %v", err)
	}

	wksp := hyp.NewWorkspace(workspaceName, workspaceDesc, workspaceImage, workspacePorts)

	dplObj, err := wksp.Deploy(context.Background(), fakeCli, expectedNS)
	if err != nil {
		t.Fatalf("error when deploying: %v", err)
	}

	if dplObj.Namespace != cteatedNS.Name {
		t.Fatalf("expected %s, got %s", dplObj.Namespace, cteatedNS.Name)
	}

}
