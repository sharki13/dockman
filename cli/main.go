package main

import (
	"fmt"

	"github.com/sharki13/dockman/backend/cmd"
	"github.com/sharki13/dockman/backend/wrapper"
)

func main() {
	dockerWrapper := wrapper.MakeDockerWrapper()

	fmt.Printf("Wrapper type: %s\n", dockerWrapper.GetType())

	instances, err := dockerWrapper.GetInstacnes()
	if err != nil && err != cmd.Err_CommandNotFound && err != wrapper.Err_DockerNotRunning {
		fmt.Println(err)
		return
	}

	for _, instance := range instances {
		fmt.Println(instance.String())
	}
}
