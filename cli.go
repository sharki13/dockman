//go:build cli

package main

import (
	"dockman/backend/common"
	"dockman/backend/docker"
	"dockman/backend/multipass"
	"fmt"
)

func main() {
	globalContainersList := []common.VM{}

	dockerContainers, err := docker.GetDockerContainers()
	if err != nil {
		if err == common.Err_DockerNotRunning {
			fmt.Printf("Docker daemon not running\n")
		} else if err == common.Err_CommandNotFound {
			fmt.Printf("Docker command not found\n")
		} else {
			fmt.Printf("Error getting docker containers: %s\n", err)
			panic(err)
		}
	}

	globalContainersList = append(globalContainersList, dockerContainers...)

	multipassVMs, err := multipass.GetMultipassVMs()
	if err != nil {
		fmt.Printf("Error getting multipass VMs: %s\n", err)
	}

	globalContainersList = append(globalContainersList, multipassVMs...)

	for _, container := range globalContainersList {
		fmt.Printf("----------------\n")
		fmt.Printf("ID: %s\n", container.GetID())
		fmt.Printf("Type: %s\n", container.GetType())
		fmt.Printf("Created: %s\n", container.GetCreated())
		fmt.Printf("Name: %s\n", container.GetName())
	}
}
