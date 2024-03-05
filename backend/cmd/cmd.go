package cmd

import "os/exec"

// Command represents a command to be executed.
type Command struct {
	Cmd  string   // The command to be executed.
	Args []string // The arguments for the command.
}

// Exec executes the specified command and returns the combined output as a string.
// If an error occurs during execution, it is also returned.
func (c *Command) Exec(command Command) (string, error) {
	cmd := exec.Command(command.Cmd, command.Args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
