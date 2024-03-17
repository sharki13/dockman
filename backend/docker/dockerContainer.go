package docker

import (
	"dockman/backend/common"
	"encoding/json"
	"time"
)

func (d *DockerContainer) GetType() common.VMType {
	return common.VMTypeDocker
}

func (d *DockerContainer) GetID() string {
	return d.ID
}

func (d *DockerContainer) GetCreated() time.Time {
	return d.Created
}

func (d *DockerContainer) GetName() string {
	return d.Name
}

func ParseDockerContainerFromJSON(in string) []DockerContainer {
	out := []DockerContainer{}

	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		panic(err)
	}

	return out
}
