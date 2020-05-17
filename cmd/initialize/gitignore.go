package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/micro/gitignore"
)

func init() {
	InitializeCmd.AddCommand(GitignoreCmd)
	funcs = append(funcs, gitignoreFunc)
}

// GitignoreCmd represents the command that executes all of tasks for
// bootstrapping a Gitignore
var GitignoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Initialize a .gitignore file",
	Run:   gitignoreFunc,
}

func gitignoreFunc(cmd *cobra.Command, args []string) {
	gi := gitignore.New(name)
	err := gi.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
