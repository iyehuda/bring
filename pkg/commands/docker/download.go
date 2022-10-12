package docker

import (
	"errors"
	"fmt"

	"github.com/iyehuda/bring/pkg/docker"
	"github.com/iyehuda/bring/pkg/utils/commands"
	"github.com/spf13/cobra"
)

var errDownloadTargetNotSet = errors.New("please specify a download target")

type downloadOptions struct {
	downloadTarget string
}

// NewDownloadCommand creates new docker download command for downloading docker images.
func NewDownloadCommand() *cobra.Command {
	opts := &downloadOptions{}
	cmd := &cobra.Command{
		Use:     "download <image> [other images...]",
		Short:   "Download a set of docker images",
		Long:    `This command depends on having docker installed and logged in to your registry (if necessary).`,
		PreRunE: validateDownloadArgs(opts),
		RunE:    downloadImages(opts),
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.Flags().StringVar(&opts.downloadTarget, "to", "", "output file path")

	return cmd
}

func validateDownloadArgs(opts *downloadOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if opts.downloadTarget == "" {
			return errDownloadTargetNotSet
		}

		cmd.SilenceUsage = true

		return nil
	}
}

func downloadImages(opts *downloadOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		docker := docker.NewFetcher(args, opts.downloadTarget, &commands.LocalRunner{})

		if err := docker.Fetch(); err != nil {
			return fmt.Errorf("failed to download images: %w", err)
		}

		return nil
	}
}
