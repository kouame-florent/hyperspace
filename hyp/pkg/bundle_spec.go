package hyp

import (
	"context"

	"github.com/docker/docker/client"
)

// aggregate container specs in same package
type BundleSpec struct {
	SpecMeta
	//Name string

	//networks used by all containers in the bundle
	Networks    []NetworkSpec
	Description string

	ContainerSpecs []*ContainerSpec
}

func NewBundleSpec(name string, descripton string, networks []NetworkSpec, containerSpecs []*ContainerSpec) *BundleSpec {

	//set all container network for bundle ones
	/*
		for _, cs := range containerSpecs {
			cs.Networks = networks

		}
	*/
	return &BundleSpec{

		Networks:       networks,
		Description:    descripton,
		ContainerSpecs: containerSpecs,
	}
}

// create and start each container from the bundle
// return object from running container
func (b *BundleSpec) Deploy(ctx context.Context, cli *client.Client) ([]*ContainerObject, error) {
	/*
		objs := []ContainerObject{}
		for _, s := range b.ContainerSpecs {

			//create container
			obj, err := s.CreateContainer(ctx, cli)
			if err != nil {
				return []*ContainerObject{}, err
			}

			//start container
			inf, err := o.StartContainer(ctx, cli, cinf.ID)
			if err != nil {
				return []ContainerObject{}, err
			}

			ctnInfos = append(ctnInfos, inf)
		}

		fmt.Printf("---> INFO LENGHT: %d", len(ctnInfos))
		return ctnInfos, nil
	*/
	return []*ContainerObject{}, nil
}
