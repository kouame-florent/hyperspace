package hyp

type VolumeTemplate struct {
	TemplateMeta
	PersistenceVolumeClaim VolumeClaimTemplate
}
