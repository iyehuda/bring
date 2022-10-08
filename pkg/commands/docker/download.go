/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package docker

import (
	"errors"

	"github.com/iyehuda/bring/pkg/docker"
	"github.com/iyehuda/bring/pkg/utils/commands"
	"github.com/spf13/cobra"
)

var downloadTarget string

// NewDownloadCommand creates new docker download command for downloading docker images
func NewDownloadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download <image> [other images...]",
		Short:   "Download a set of docker images",
		Long:    `This command depends on having docker installed and logged in to your registry (if necessary).`,
		PreRunE: validateDownloadArgs,
		RunE:    downloadImages,
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.Flags().StringVar(&downloadTarget, "to", "", "output file path")

	return cmd
}

func validateDownloadArgs(cmd *cobra.Command, args []string) error {
	if downloadTarget == "" {
		return errors.New("please specify a download target")
	}

	cmd.SilenceUsage = true

	return nil
}

func downloadImages(cmd *cobra.Command, args []string) error {
	docker := docker.NewFetcher(args, downloadTarget, &commands.LocalRunner{})

	return docker.Fetch()
}
