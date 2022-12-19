package hyp

import (
	"context"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/pointer"
)

const (
	ContainerDefaultPortName = "http"
)

type Workspace struct {
	Name        string
	Description string
	Image       string
	Ports       []int32
	Volumes     []string
}

type DeploymentObject struct {
	ObjectMeta
	Namespace string
	CreatedAt time.Time
	DeletedAt time.Time
}

func NewWorkspace(name, description, image string, ports []int32) Workspace {
	return Workspace{
		Name:        name,
		Description: description,
		Image:       image,
		Ports:       ports,
	}

}

func (w *Workspace) Deploy(ctx context.Context, cls kubernetes.Interface, namespace string) (*DeploymentObject, error) {

	dpl := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      w.Name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": w.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": w.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  w.Name,
							Image: w.Image,
							Ports: containerPorts(w.Ports),
						},
					},
				},
			},
		},
	}

	ctnDepl, err := cls.AppsV1().Deployments(namespace).Create(ctx, dpl, metav1.CreateOptions{})
	if err != nil {
		return &DeploymentObject{}, err
	}
	return &DeploymentObject{
		ObjectMeta: ObjectMeta{
			UID:       string(ctnDepl.UID),
			Name:      ctnDepl.Name,
			CreatedAt: ctnDepl.CreationTimestamp.Time,
		},
		Namespace: ctnDepl.Namespace,
	}, nil
}

func containerPorts(ports []int32) []apiv1.ContainerPort {
	ctnPorts := []apiv1.ContainerPort{}

	for _, p := range ports {
		ctnPort := apiv1.ContainerPort{
			Name:          ContainerDefaultPortName,
			Protocol:      apiv1.ProtocolTCP,
			ContainerPort: p,
		}

		ctnPorts = append(ctnPorts, ctnPort)
	}

	return ctnPorts
}
