package hyp

type ContainerObject struct {
	ObjectMeta
}

type ContainerTemplate struct {
	TemplateMeta
	Image        string
	Ports        []int32
	VolumeMounts []volumeMount
}

type volumeMount struct {
	Name      string
	MountPath string
	ReadOnly  bool
}

func NewContainerTemplate(name, namespace, image string, volumeMounts []volumeMount) ContainerTemplate {
	return ContainerTemplate{
		TemplateMeta: TemplateMeta{
			Name:      name,
			Namespace: namespace,
		},
		VolumeMounts: volumeMounts,
	}
}
