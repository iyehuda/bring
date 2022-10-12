package tests

import (
	"os/exec"
	"strings"

	"github.com/stretchr/testify/mock"
)

type MockCommandRunner struct {
	mock.Mock
}

func (r *MockCommandRunner) Run(cmd *exec.Cmd) error {
	args := r.Called(cmd)

	return args.Error(0)
}

type CommandPredicate func(cmd *exec.Cmd) bool

func CommandIncludes(str string) CommandPredicate {
	return func(cmd *exec.Cmd) bool {
		return strings.Contains(strings.Join(cmd.Args, " "), str)
	}
}

// FakeCommandRunner enables unit testing modules that run commands with
// a predictable error.
type FakeCommandRunner struct {
	Err error
}

// Run runs a command with a predefined error.
func (r *FakeCommandRunner) Run(cmd *exec.Cmd) error {
	return r.Err
}
