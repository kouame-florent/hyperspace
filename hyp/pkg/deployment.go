package hyp

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/utils/pointer"
)

const (
	ContainerDefaultPortName = "http"
)

type DeploymentObject struct {
	ObjectMeta
}

type DeploymentTemplate struct {
	TemplateMeta
	Containers []ContainerTemplate
}

func NewDeploymentTemplate(name, namespace string, containers []ContainerTemplate) DeploymentTemplate {
	return DeploymentTemplate{
		TemplateMeta: TemplateMeta{
			Name:      name,
			Namespace: namespace,
		},
		Containers: containers,
	}
}

func (d *DeploymentTemplate) CreateDeployment(ctx context.Context, cls kubernetes.Interface) (DeploymentObject, error) {

	dpl := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.Name,
			Namespace: d.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": d.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": d.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: containers(d.Containers),
				},
			},
		},
	}

	k8sDepl, err := cls.AppsV1().Deployments(d.Namespace).Create(ctx, dpl, metav1.CreateOptions{})
	if err != nil {
		return DeploymentObject{}, err
	}

	return DeploymentObject{
		ObjectMeta: ObjectMeta{
			UID:  string(k8sDepl.UID),
			Name: k8sDepl.Name,
		},
	}, nil
}

func containers(tpls []ContainerTemplate) []apiv1.Container {

	ctns := []apiv1.Container{}

	for _, tp := range tpls {
		ct := apiv1.Container{
			Name:  tp.Name,
			Image: tp.Image,
			Ports: containerPorts(tp.Ports),
		}

		ctns = append(ctns, ct)
	}

	return ctns
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
