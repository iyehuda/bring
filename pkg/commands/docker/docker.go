package docker

import "github.com/spf13/cobra"

// NewCommand creates new docker command for managing docker images.
func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "docker",
		Short: "Manage docker images",
		Long:  `Use this command to download, retag and publish docker images`,
	}

	command.AddCommand(
		NewDownloadCommand(),
		NewUploadCommand(),
	)

	return command
}
