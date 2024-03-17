package common

import (
	"errors"
	"time"
)

// VMType represents the type of a virtual machine.
type VMType string

const (
	VMTypeDocker    VMType = "docker"
	VMTypeKVM       VMType = "lima"
	VMTypeColima    VMType = "colima"
	VMTypeMultipass VMType = "multipass"
)

// Status represents the status of a virtual machine.
type Status string

const (
	StatusRunning  Status = "running"
	StatusStopped  Status = "stopped"
	StatusPaused   Status = "paused"
	StatusUnknown  Status = "unknown"
	StatusStarting Status = "starting"
	StatusStopping Status = "stopping"
	StatusBroken   Status = "broken"
)

// VM is an interface that represents a virtual machine.
type VM interface {
	// returns the type of the virtual machine
	GetType() VMType

	// returns the ID of the virtual machine
	GetID() string

	// returns the creation time of the virtual machine
	GetCreated() time.Time

	// returns the name of the virtual machine
	GetName() string

	// GetStatus() string
}

var Err_DockerNotRunning = errors.New("docker daemon not running")
var Err_CommandNotFound = errors.New("command not found")
