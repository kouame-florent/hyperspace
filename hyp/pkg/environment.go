package hyp

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

type Environment struct {
	Service   ServiceTemplate
	Container ContainerTemplate
	Volume    VolumeClaimTemplate
}

type EnvironmentObject struct {
	Service   ServiceObject
	Container ContainerObject
	Volume    VolumeClaimObject
}

func NewEnvironment(name, description, image string, containerPorts []int32) Environment {
	return Environment{}

}

func (e *Environment) Deploy(ctx context.Context, cls kubernetes.Interface, namespace string) (EnvironmentObject, error) {

	return EnvironmentObject{}, nil
}
