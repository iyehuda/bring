/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package commands

import "os/exec"

// Runner is able to run a command in some environment
type Runner interface {
	Run(cmd *exec.Cmd) error
}

// LocalRunner runs commands locally
type LocalRunner struct{}

// Run runs a command locally
func (r *LocalRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
