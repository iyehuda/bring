/*
Copyright © 2022 Yehuda Chikvashvili <yehudaac1@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bring",
	Short: "Fetch assets for offline usage",
	Long: `bring is a CLI that enables you to donwload assets fron online sources such as docker images, helm charts, etc.
Use this application to fetch every artifact necessary for your later offline amusement.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
