//go:build cmd

package main

import (
	backend "dockmaster/backend"
	"fmt"
)

func main() {
	dockerContainers := backend.GetDockerContainers()

	for _, container := range dockerContainers {
		fmt.Printf("Docker container: %s\n", container.GetID())
	}

	multipassVMs := backend.GetMultipassVMs()

	for _, vm := range multipassVMs {
		fmt.Printf("Multipass VM: %s\n", vm.GetID())
	}
}
