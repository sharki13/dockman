package docker

import (
	"dockman/backend/cmd"
	"dockman/backend/common"
	"strings"
)

func GetDockerContainers() []common.VM {
	dockerCmd := cmd.Command{
		Cmd:  "docker",
		Args: []string{"ps"},
	}

	out, err := dockerCmd.Exec(dockerCmd)
	if err != nil {
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
