package hyp

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

type Environment struct {
	Service    ServiceTemplate
	Deployment DeploymentTemplate
}

type EnvironmentObject struct {
	Service   ServiceObject
	Container ContainerObject
}

type CreateEnvironmentParams struct {
	Name           string
	Description    string
	Namespace      string
	Image          string
	ClaimSize      string
	ServicePorts   map[int32]int
	ContainerPorts []int32
	VolumeMounts   []VolumeMountTemplate
}

func NewEnvironment(p CreateEnvironmentParams) Environment {

	//svcTpl := NewServiceTemplate(p.Name, p.Namespace, p.ServicePorts)
	//depTpl := NewContainerTemplate(p.Name, p.Namespace, p.Image, p.VolumeMounts)
	//volTpl := NewVolumeClaimTemplate(p.Name, p.Namespace, p.ClaimSize)

	return Environment{
		/*
			Service:    svcTpl,
			Deployment: depTpl,
			Volume:     volTpl,
		*/
	}

}

func (e *Environment) Deploy(ctx context.Context, cls kubernetes.Interface) (EnvironmentObject, error) {

	/*
		svcObj, err := e.Service.CreateService(ctx, cls)
		if err != nil {
			return EnvironmentObject{}, err
		}
	*/
	//ctnObj, err := e.Conta

	return EnvironmentObject{}, nil
}
