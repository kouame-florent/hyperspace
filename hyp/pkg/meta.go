package hyp

import "time"

type state = string

const (
	//created, restarting, running, removing, paused, exited, or dead

	running  state = "running"
	created  state = "created"
	stopped  state = "stopped"
	removing state = "removing"
)

type StatusMeta struct {
	ID        string
	CreatedAt time.Time
	DeletedAt time.Time
	Labels    map[string]string
	Stat      state
}

type SpecMeta struct {
	ID        string
	CreatedAt time.Time
	DeletedAt time.Time
	Labels    map[string]string
}
