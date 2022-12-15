package hyp

/*
// add image of other to bundle b and concactenate names and description
func (b *Bundle) Add(other *Bundle) *Bundle {
	//change new added containers networks
	configs := changeNetwork(other.ContainerSpecs, b.Networks)

	b.ContainerSpecs = append(b.ContainerSpecs, configs...)
	b.Name = b.Name + "_" + other.Name
	b.Description = "\n\n" + other.Description

	return b
}
*/

/*
func changeNetwork(containerConfigs []ContainerSpec, networks []string) []ContainerSpec {

	configs := containerConfigs

	//change all container network for bundle ones
	for _, c := range containerConfigs {
		c.Networks = networks
		configs = append(configs, c)
	}

	return configs

}
*/

type BundleObject struct {
	ObjectMeta
	Infos []ContainerObject
}
