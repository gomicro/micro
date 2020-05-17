package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/micro/readme"
)

func init() {
	InitializeCmd.AddCommand(ReadmeCmd)
	funcs = append(funcs, readmeFunc)
}

// ReadmeCmd represents the command that executes all of tasks for bootstrapping
// a Readme
var ReadmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Initialize a Readme file",
	Run:   readmeFunc,
}

func readmeFunc(cmd *cobra.Command, args []string) {
	r := readme.New(name, source)
	if installable {
		r.EnableInstallable()
	}

	err := r.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
