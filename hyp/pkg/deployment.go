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
	Pod PodTemplate
}

func NewDeploymentTemplate(name, namespace string, pod PodTemplate) DeploymentTemplate {
	return DeploymentTemplate{
		TemplateMeta: TemplateMeta{
			Name:      name,
			Namespace: namespace,
		},
		Pod: pod,
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
					Containers: containers(d.Pod.Containers),
					Volumes:    volumes(d.Pod.Volumes),
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

	for _, ctp := range tpls {

		ct := apiv1.Container{
			Name:         ctp.Name,
			Image:        ctp.Image,
			Ports:        containerPorts(ctp.Ports),
			VolumeMounts: volumeMount(ctp.VolumeMounts),
		}

		ctns = append(ctns, ct)
	}
	return ctns
}

func volumeMount(vmtpls []VolumeMountTemplate) []apiv1.VolumeMount {
	volMounts := []apiv1.VolumeMount{}
	for _, vtpl := range vmtpls {
		volMount := apiv1.VolumeMount{
			Name:      vtpl.Name,
			MountPath: vtpl.MountPath,
			ReadOnly:  vtpl.ReadOnly,
		}
		volMounts = append(volMounts, volMount)
	}

	return volMounts
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

func volumes(vtpls []VolumeTemplate) []apiv1.Volume {
	volumes := []apiv1.Volume{}

	for _, vtpl := range vtpls {
		vol := apiv1.Volume{
			Name: vtpl.Name,
			VolumeSource: apiv1.VolumeSource{
				PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
					ClaimName: vtpl.Name,
				},
			},
		}

		volumes = append(volumes, vol)
	}

	return volumes
}
