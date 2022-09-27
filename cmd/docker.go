/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package cmd

import (
	"github.com/iyehuda/bring/cmd/docker"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Manage docker images",
	Long:  `Use this command to download, retag and publish docker images`,
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.AddCommand(docker.DownloadCmd)
}
