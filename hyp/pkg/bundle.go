package hyp

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/client"
	"github.com/google/uuid"
)

// aggregate container configs in same package
type Bundle struct {
	ID   string
	Name string

	//networks used by all containers
	Networks    []string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	ContainerSpecs []ContainerSpec
}

func NewBundle(name string, descripton string, networks []string, containerSpecs []ContainerSpec) *Bundle {

	//change all container network for bundle ones
	configs := changeNetwork(containerSpecs, networks)

	return &Bundle{
		ID:             uuid.New().String(),
		Name:           name,
		Networks:       networks,
		Description:    descripton,
		ContainerSpecs: configs,
	}
}

// create and start each container from the bundle and give them the same network id.
// return info from running container
func (b *Bundle) Deploy(ctx context.Context, cli *client.Client) ([]ContainerInfo, error) {
	ctnInfos := []ContainerInfo{}
	for _, o := range b.ContainerSpecs {

		//create container
		id, err := o.CreateContainer(ctx, cli)
		if err != nil {
			return []ContainerInfo{}, err
		}

		//start container
		inf, err := o.StartContainer(ctx, cli, id)
		if err != nil {
			return []ContainerInfo{}, err
		}

		ctnInfos = append(ctnInfos, inf)
	}

	fmt.Printf("---> INFO LENGHT: %d", len(ctnInfos))
	return ctnInfos, nil
}

// add image of other to bundle b and concactenate names and description
func (b *Bundle) Add(other *Bundle) *Bundle {
	//change new added containers networks
	configs := changeNetwork(other.ContainerSpecs, b.Networks)

	b.ContainerSpecs = append(b.ContainerSpecs, configs...)
	b.Name = b.Name + "_" + other.Name
	b.Description = "\n\n" + other.Description

	return b
}

func changeNetwork(containerConfigs []ContainerSpec, networks []string) []ContainerSpec {

	configs := containerConfigs

	//change all container network for bundle ones
	for _, c := range containerConfigs {
		c.Networks = networks
		configs = append(configs, c)
	}

	return configs

}
