package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/service/gofiles"
)

func init() {
	InitializeCmd.AddCommand(GofilesCmd)
	funcs = append(funcs, gofilesFunc)
}

// GofilesCmd represents the command that executes all of tasks for
// bootstrapping the  Go files
var GofilesCmd = &cobra.Command{
	Use:   "gofiles",
	Short: "Generate a bootstrap of Go files",
	Run:   gofilesFunc,
}

func gofilesFunc(cmd *cobra.Command, args []string) {
	gf := gofiles.New(name)
	if database {
		gf.EnableDB()
	}

	err := gf.WriteFiles()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
