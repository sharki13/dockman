package wrapper

import "encoding/json"

// VMType represents the type of a virtual machine.
type WarpperType string

const (
	DockerWrapperType    WarpperType = "docker"
	MultipassWrapperType WarpperType = "multipass"
)

// Wrapper is an interface that represents a virtual machine.
type Wrapper interface {
	// returns the type of the virtual machine
	GetType() WarpperType
	GetInstances() ([]Instance, error)
}

type StateType int

const (
	StateRunning StateType = 1
	StateStopped StateType = 2
)

func (s StateType) String() string {
	switch s {
	case StateRunning:
		return "Running"
	case StateStopped:
		return "Stopped"
	default:
		return "Unknown"
	}
}

type InstanceType string

func (i InstanceType) String() string {
	return string(i)
}

const (
	DockerContainerInstance InstanceType = "docker-container"
	MultipassInstance       InstanceType = "multipass-vm"
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

func (i *Instance) JSON() ([]byte, error) {
	return json.Marshal(i)
}
