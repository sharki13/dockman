package main

import (
	"fmt"

	"github.com/sharki13/dockman/backend/cmd"
)

func main() {
	lsCommand := cmd.Command{
		Cmd:  "docker",
		Args: []string{"ps"},
	}

	out, err := lsCommand.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)
}
