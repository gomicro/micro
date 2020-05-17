package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/micro/cmd/initialize"
)

func init() {
	RootCmd.AddCommand(initialize.InitializeCmd)
}

// RootCmd represents the root command for the micro cli
var RootCmd = &cobra.Command{
	Use:   "micro",
	Short: "A cli tool for bootstrapping a new microservice",
}

// Execute wraps the command executor to make it possible to trigger easily
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("Failed to execute: %v", err.Error())
		os.Exit(1)
	}
}
