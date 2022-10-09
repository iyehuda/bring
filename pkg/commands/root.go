package commands

import "github.com/spf13/cobra"

// NewRootCommand creates a new root command for bring CLI
func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "bring",
		Short: "Fetch assets for offline usage",
		Long: `bring is a CLI that enables you to download assets from online sources such as docker images, helm charts, etc.
Use this application to fetch every artifact necessary for your later offline amusement.`,
	}
}
