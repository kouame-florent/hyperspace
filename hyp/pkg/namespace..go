package hyp

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceObject struct {
	ObjectMeta
}

func CreateNamespace(ctx context.Context, cls kubernetes.Interface, name string) (*NamespaceObject, error) {
	nsName := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: name}}
	ns, err := cls.CoreV1().Namespaces().Create(ctx, nsName, metav1.CreateOptions{})
	if err != nil {
		return &NamespaceObject{}, err
	}

	return &NamespaceObject{
		ObjectMeta: ObjectMeta{
			UID:  string(ns.UID),
			Name: ns.Name,
		},
	}, nil
}

func Namespace(ctx context.Context, cls kubernetes.Interface, name string) (NamespaceObject, error) {
	ns, err := cls.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return NamespaceObject{}, err
	}
	return NamespaceObject{
		ObjectMeta: ObjectMeta{
			UID:       string(ns.UID),
			Name:      ns.Name,
			CreatedAt: ns.CreationTimestamp.Time,
		},
	}, nil
}
