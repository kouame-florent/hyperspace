package hyp

import "time"

type ContainerInfo struct {
	ID        string
	Name      string
	Network   NetworkInfo
	Volume    VolumeInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BundleInfo struct {
	ID    string
	Infos []ContainerInfo
}

func NewBundleInfo() *BundleInfo {
	return &BundleInfo{}
}

func (b *BundleInfo) StopBundle() {

}

func (b *BundleInfo) RestartBundle() {

}

func (b *BundleInfo) DeleteBundle() {

}
