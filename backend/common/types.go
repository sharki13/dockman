package common

import "time"

type VMType string

const (
	VMTypeDocker    VMType = "docker"
	VMTypeKVM       VMType = "lima"
	VMTypeColima    VMType = "colima"
	VMTypeMultipass VMType = "multipass"
)

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

type VM interface {
	GetType() VMType
	GetID() string
	GetCreated() time.Time
	GetName() string
	// GetStatus() string
}
