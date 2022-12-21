package hyp

type ContainerObject struct {
	ObjectMeta
}

type ContainerTemplate struct {
	TemplateMeta
	Image        string
	Ports        []int32
	VolumeMounts []VolumeMountTemplate
}

type VolumeMountTemplate struct {
	// mount name is the same as volume claim name
	Name      string
	MountPath string
	ReadOnly  bool
}

func NewVolumemountTemplate(name, mountPath string, readOnly bool) VolumeMountTemplate {
	return VolumeMountTemplate{
		Name:      name,
		MountPath: mountPath,
		ReadOnly:  readOnly,
	}
}

func NewContainerTemplate(name, namespace, image string, volumeMounts []VolumeMountTemplate) ContainerTemplate {
	return ContainerTemplate{
		TemplateMeta: TemplateMeta{
			Name:      name,
			Namespace: namespace,
		},
		VolumeMounts: volumeMounts,
	}
}
