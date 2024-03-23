package wrapper

import (
	"encoding/json"
	"strings"

	"github.com/sharki13/dockman/backend/cmd"
)

func MakeDockerWrapper() Wrapper {
	return &DockerWrapper{}
}

type DockerWrapper struct {
}

type internalPsContainer struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
}

func (e *internalPsContainer) stateToStateType() StateType {
	if e.State == "running" {
		return StateRunning
	}
	return StateStopped
}

func (e *internalPsContainer) ToInstance() Instance {
	return Instance{
		ID:    e.ID,
		Name:  e.Names,
		State: e.stateToStateType(),
		Type:  DockerContainerInstance,
	}
}

var psCommand = cmd.Command{
	Cmd:  "docker",
	Args: []string{"ps", "-a", "--format", "json"},
}

func (w *DockerWrapper) GetType() WarpperType {
	return DockerWrapperType
}

func (w *DockerWrapper) GetInstacnes() ([]Instance, error) {
	ret := []Instance{}

	out, err := psCommand.Exec()
	if err != nil {
		if strings.Contains(out, "docker daemon is not running") {
			return ret, Err_DockerNotRunning
		}
		return ret, err
	}

	// iterate over each line in out and print
	for _, line := range strings.Split(out, "\n") {
		var container internalPsContainer
		if line == "" {
			continue
		}
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			return ret, err
		}
		ret = append(ret, container.ToInstance())
	}

	return ret, nil
}
