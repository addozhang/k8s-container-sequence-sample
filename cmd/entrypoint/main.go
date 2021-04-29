package main

import (
	"github.com/spf13/cobra"
	"os"
)

var(
	rootCommand = &cobra.Command{
		Use:   "entrypoint",
	}
)

func init() {
	rootCommand.AddCommand(waitCommand)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}