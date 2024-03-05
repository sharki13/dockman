package docker

import (
	"dockman/backend/cmd"
	"dockman/backend/common"
	"strings"
)

// GetDockerContainers retrieves a list of Docker containers.
// It executes the "docker ps" command and parses the output to obtain the container information.
// If the Docker daemon is not running, it returns an empty list of containers.
// It returns a slice of common.VM, which is an interface representing a virtual machine.
func GetDockerContainers() []common.VM {
	dockerCmd := cmd.Command{
		Cmd:  "docker",
		Args: []string{"ps"},
	}

	out, err := dockerCmd.Exec(dockerCmd)
	if err != nil {
		// if out contains "this error may indicate that the docker daemon is not running"
		// return an empty list of containers
		if strings.Contains(out, "this error may indicate that the docker daemon is not running") {
			return []common.VM{}
		}
		panic(err)
	}

	containers := parseDockerContainers(out)

	// Convert to VM interface
	ret := make([]common.VM, 0)
	for _, container := range containers {
		ret = append(ret, &container)
	}

	return ret
}

// parseDockerContainers parses the output of the "docker ps" command to obtain the container IDs.
// It then executes the "docker inspect" command for each container ID to retrieve detailed container information.
// It returns a slice of DockerContainer, which represents a Docker container.
func parseDockerContainers(out string) []DockerContainer {
	lines := strings.Split(out, "\n")
	ret := make([]DockerContainer, 0)
	containersIds := make([]string, 0)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 1 {
			continue
		}
		containersIds = append(containersIds, fields[0])
	}

	for _, id := range containersIds {
		dockerInspectCmd := cmd.Command{
			Cmd:  "docker",
			Args: []string{"inspect", id},
		}
		out, err := dockerInspectCmd.Exec(dockerInspectCmd)
		if err != nil {
			continue
		}

		ret = ParseDockerContainerFromJSON(out)
	}

	return ret
}
