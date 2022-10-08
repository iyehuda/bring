package commands

import (
	"github.com/iyehuda/bring/pkg/commands/docker"
	"github.com/spf13/cobra"
)

// NewApp creates new bring CLI cobra command
func NewApp() *cobra.Command {
	rootCmd := NewRootCommand()

	rootCmd.AddCommand(
		docker.NewCommand(),
	)

	return rootCmd
}
