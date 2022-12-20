package hyp

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

const (
	ServiceDefaultPortName = "http"
)

type ServiceObject struct {
	ObjectMeta
}

type ServiceTemplate struct {
	TemplateMeta
	Ports map[int32]int
}

func NewServiceTemplate(name, namespace string, servicePorts map[int32]int) ServiceTemplate {
	return ServiceTemplate{
		TemplateMeta: TemplateMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (s *ServiceTemplate) CreateService(ctx context.Context, cls kubernetes.Interface) (ServiceObject, error) {

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: s.Name,
		},
		Spec: corev1.ServiceSpec{
			Ports: servicePort(s.Ports),
			Type:  corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": s.Name,
			},
		},
	}

	crtSvc, err := cls.CoreV1().Services(s.Namespace).Create(ctx, svc, metav1.CreateOptions{})
	if err != nil {
		return ServiceObject{}, nil
	}

	return ServiceObject{
		ObjectMeta: ObjectMeta{
			UID:  string(crtSvc.UID),
			Name: crtSvc.Name,
		},
	}, nil
}

// map between port and target port
func servicePort(ports map[int32]int) []corev1.ServicePort {

	svcPorts := []corev1.ServicePort{}

	for k, v := range ports {
		port := corev1.ServicePort{
			Name:       ServiceDefaultPortName,
			Protocol:   corev1.ProtocolTCP,
			Port:       k,
			TargetPort: intstr.FromInt(v),
		}
		svcPorts = append(svcPorts, port)
	}

	return svcPorts
}
