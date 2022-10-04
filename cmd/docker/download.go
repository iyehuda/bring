/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package docker

import (
	"errors"

	"github.com/iyehuda/bring/pkg/docker"
	"github.com/spf13/cobra"
)

var downloadTarget string

// DownloadCmd represents the download command
var DownloadCmd = &cobra.Command{
	Use:     "download <image> [other images...]",
	Short:   "Download a set of docker images",
	Long:    `This command depends on having docker installed and logged in to your registry (if necessary).`,
	PreRunE: validateArgs,
	RunE:    dockerDownload,
	Args:    cobra.MinimumNArgs(1),
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if downloadTarget == "" {
		return errors.New("please specify a download target")
	}

	cmd.SilenceUsage = true

	return nil
}

func dockerDownload(cmd *cobra.Command, args []string) error {
	docker := docker.NewFetcher(args, downloadTarget)

	return docker.Fetch()
}

func init() {
	DownloadCmd.Flags().StringVar(&downloadTarget, "to", "", "output file path")
}
