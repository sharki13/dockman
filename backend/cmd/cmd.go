package cmd

import (
	"errors"
	"os/exec"
)

var Err_CommandNotFound = errors.New("command not found")

// Command represents a command to be executed.
type Command struct {
	Cmd  string   // The command to be executed.
	Args []string // The arguments for the command.
}

func (c *Command) checkIfCommandExist() bool {
	_, err := exec.LookPath(c.Cmd)
	return err == nil
}

// Exec executes the specified command and returns the combined output as a string.
// If an error occurs during execution, it is also returned.
func (c *Command) Exec() (string, error) {
	if !c.checkIfCommandExist() {
		return "", Err_CommandNotFound
	}

	cmd := exec.Command(c.Cmd, c.Args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}

	return string(out), nil
}
