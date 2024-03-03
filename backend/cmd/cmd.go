package cmd

import "os/exec"

type Command struct {
	Cmd  string
	Args []string
}

func (c *Command) Exec(command Command) (string, error) {
	cmd := exec.Command(command.Cmd, command.Args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
