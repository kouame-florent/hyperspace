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
	// UID from container engine if applicable
	UID string
	// Object name in container engine
	Name      string
	CreatedAt time.Time
	DeletedAt time.Time
	Labels    map[string]string
	Status    status
}

type SpecMeta struct {
	// ID for persistence
	ID string
	// User defined name
	Tag       string
	CreatedAt time.Time
	DeletedAt time.Time
	Labels    map[string]string
}
