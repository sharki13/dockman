package backend

import "time"

type VMType string

const (
	VMTypeDocker    VMType = "docker"
	VMTypeKVM       VMType = "lima"
	VMTypeColima    VMType = "colima"
	VMTypeMultipass VMType = "multipass"
)

type VM interface {
	GetType() VMType
	GetID() string
	GetCreated() time.Time
}
