package integration

import (
	"bytes"
	"io"
	"testing"

	"github.com/iyehuda/bring/pkg/commands"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

const helpMessage = `Usage:
  bring `

type executionVars struct {
	app  *cobra.Command
	args []string
	out  io.Writer
	err  io.Writer
}

type executeOption func(vars *executionVars)

func withOutput(out io.Writer) executeOption {
	return func(vars *executionVars) {
		vars.out = out
	}
}

func withError(err io.Writer) executeOption {
	return func(vars *executionVars) {
		vars.err = err
	}
}

func execute(osArgs []string, opts ...executeOption) error {
	app := commands.NewApp()
	app.SetArgs(osArgs)

	execution := &executionVars{
		app:  app,
		args: osArgs,
		out:  nil,
		err:  nil,
	}

	for _, opt := range opts {
		opt(execution)
	}

	app.SetOut(execution.out)
	app.SetErr(execution.err)

	return app.Execute()
}

func assertErrorIf(t *testing.T, err error, condition bool) {
	if condition {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
}

func assertContainsIf(t *testing.T, buf *bytes.Buffer, substr string, condition bool) {
	content := buf.String()
	if condition {
		assert.Contains(t, content, substr)
	} else {
		assert.NotContains(t, content, substr)
	}
}

func assertUsagePrintedIf(t *testing.T, buf *bytes.Buffer, condition bool) {
	assertContainsIf(t, buf, helpMessage, condition)
}
