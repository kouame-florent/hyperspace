package hyp

//type version = string

type bundle struct {
	version  string
	Networks []Tnetwork
	Volumes  []Tvolume
	Services []Tservice
}

type Tnetwork struct {
	name string
}

type Tvolume struct {
	name   string
	driver string
}

type Tservice struct {
	name        string
	image       string
	environment []string
	networks    []Tnetwork
	volumes     []Tvolume
	ports       []string
}
