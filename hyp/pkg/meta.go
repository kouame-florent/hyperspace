package hyp

import "time"

type status = string

const (
	//created, restarting, running, removing, paused, exited, or dead

	Running  status = "running"
	Created  status = "created"
	Stopped  status = "stopped"
	Removing status = "removing"
)

type ObjectMeta struct {
	UID       string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateMeta struct {
	Name      string
	Namespace string
	CreatedAt time.Time
	UpdatedAt time.Time
}
