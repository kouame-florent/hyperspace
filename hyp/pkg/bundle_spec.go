package hyp

// aggregate container specs in same package
type BundleSpec struct {
	SpecMeta
	Name string

	//networks used by all containers in the bundle
	Networks    []string
	Description string

	ContainerSpecs []ContainerSpec
}

func NewBundleSpec(name string, descripton string, networks []string, containerSpecs []ContainerSpec) *BundleSpec {

	//change all container network for bundle ones
	//configs := changeNetwork(containerSpecs, networks)

	return &BundleSpec{

		Name:           name,
		Networks:       networks,
		Description:    descripton,
		ContainerSpecs: containerSpecs,
		//ContainerSpecs: configs,
	}
}

/*
// create and start each container from the bundle and give them the same network id.
// return info from running container
func (b *BundleSpec) Deploy(ctx context.Context, cli *client.Client) ([]ContainerObject, error) {
	ctnInfos := []ContainerObject{}
	for _, o := range b.ContainerSpecs {

		//create container
		cinf, err := o.CreateContainer(ctx, cli)
		if err != nil {
			return []ContainerObject{}, err
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
}
*/
