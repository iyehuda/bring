/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package utils

import "os/exec"

// CommandRunner is able to run a command in some environment
type CommandRunner interface {
	Run(cmd *exec.Cmd) error
}

// LocalCommandRunner runs commands locally
type LocalCommandRunner struct{}

// Run runs a command locally
func (r *LocalCommandRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
