package main

import (
	"fmt"
	"os"

	"github.com/iyehuda/bring/pkg/commands"
)

func main() {
	app := commands.NewApp()
	if err := app.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
