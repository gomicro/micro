package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultShell = "zsh"
)

var (
	shell string
)

func init() {
	RootCmd.AddCommand(CompletionCmd)

	CompletionCmd.Flags().StringVar(&shell, "shell", defaultShell, "desired shell to generate completions for")
}

// CompletionCmd represents the command for generating completion files for the
// micro cli.
var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion files for the micro cli",
	Run:   completionFunc,
}

func completionFunc(cmd *cobra.Command, args []string) {
	var err error
	switch strings.ToLower(shell) {
	case "bash":
		err = RootCmd.GenBashCompletion(os.Stdout)
	case "fish":
		err = RootCmd.GenFishCompletion(os.Stdout, false)
	case "ps", "powershell", "power_shell":
		err = RootCmd.GenPowerShellCompletion(os.Stdout)
	case "zsh":
		err = RootCmd.GenZshCompletion(os.Stdout)
	default:
	}

	if err != nil {
		fmt.Printf("Failed to execute: %v", err.Error())
		os.Exit(1)
	}
}
