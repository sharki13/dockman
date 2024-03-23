package wrapper

// VMType represents the type of a virtual machine.
type WarpperType string

const (
	DockerWrapperType WarpperType = "docker"
)

// Wrapper is an interface that represents a virtual machine.
type Wrapper interface {
	// returns the type of the virtual machine
	GetType() WarpperType
	GetInstacnes() ([]Instance, error)
}

type StateType string

const (
	StateRunning StateType = "running"
	StateStopped StateType = "stopped"
)

func (s StateType) String() string {
	return string(s)
}

type InstanceType string

func (i InstanceType) String() string {
	return string(i)
}

const (
	DockerContainerInstance InstanceType = "docker-container"
)

type Instance struct {
	ID    string       `json:"ID"`
	Name  string       `json:"Name"`
	State StateType    `json:"State"`
	Type  InstanceType `json:"Type"`
}

func (i *Instance) String() string {
	return "[" + i.Type.String() + "] " + i.Name + " " + i.State.String()
}
