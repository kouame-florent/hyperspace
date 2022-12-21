package hyp

type PodTemplate struct {
	TemplateMeta
	Containers []ContainerTemplate
	Volumes    []VolumeTemplate
}

type PodObject struct {
}
