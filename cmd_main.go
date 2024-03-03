//go:build cmd

package main

import (
	"dockman/backend/docker"
	"dockman/backend/multipass"
	"fmt"
)

func main() {
	dockerContainers := docker.GetDockerContainers()

	for _, container := range dockerContainers {
		fmt.Printf("Docker container: %s\n", container.GetID())
	}

	multipassVMs := multipass.GetMultipassVMs()

	for _, vm := range multipassVMs {
		fmt.Printf("Multipass VM: %s\n", vm.GetID())
	}
}
