/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package cmd

import (
	"errors"

	"github.com/iyehuda/bring/pkg/docker"
	"github.com/spf13/cobra"
)

var downloadTarget string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:     "download <image> [other images...]",
	Short:   "Download a set of docker images",
	Long:    `This command depends on having docker installed and logged in to your registry (if necessary).`,
	PreRunE: validateArgs,
	RunE:    dockerDownload,
	Args:    cobra.MinimumNArgs(1),
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if downloadTarget == "" {
		return errors.New("Please specify a download target")
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	return nil
}

func dockerDownload(cmd *cobra.Command, args []string) error {
	docker := docker.NewFetcher(args, downloadTarget)

	return docker.Fetch()
}

func init() {
	dockerCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVar(&downloadTarget, "to", "", "output file path")
}
