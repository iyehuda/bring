/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package commands

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	outputWriter io.Writer = os.Stdout
)

// SetOutput overrides the default output stream
// It may be useful for testing or embedding in other packages
func SetOutput(out io.Writer) {
	outputWriter = out
}

// NewRootCommand creates a new root command for bring CLI
func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "bring",
		Short: "Fetch assets for offline usage",
		Long: `bring is a CLI that enables you to download assets from online sources such as docker images, helm charts, etc.
Use this application to fetch every artifact necessary for your later offline amusement.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetOut(outputWriter)

			return nil
		},
	}
}
