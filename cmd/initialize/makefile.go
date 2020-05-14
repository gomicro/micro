package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/service/mkfile"
)

func init() {
	InitializeCmd.AddCommand(MakefileCmd)
	funcs = append(funcs, makefileFunc)
}

// MakefileCmd represents the command that executes all of tasks for
// bootstrapping a Makefile
var MakefileCmd = &cobra.Command{
	Use:   "makefile",
	Short: "Generate a bootstrap of a Makefile",
	Run:   makefileFunc,
}

func makefileFunc(cmd *cobra.Command, args []string) {
	m := mkfile.New(org, name)
	err := m.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
