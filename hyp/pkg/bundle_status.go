package hyp

type BundleStatus struct {
	ID    string
	Infos []ContainerStatus
}

func NewBundleInfo() *BundleStatus {
	return &BundleStatus{}
}
