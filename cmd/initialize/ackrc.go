package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/micro/ackrc"
)

func init() {
	InitializeCmd.AddCommand(AckrcCmd)
	funcs = append(funcs, ackrcFunc)
}

// AckrcCmd represents the command that executes all of tasks for bootstrapping
// a .ackrc
var AckrcCmd = &cobra.Command{
	Use:   "ackrc",
	Short: "Initialize a .ackrc file",
	Run:   ackrcFunc,
}

func ackrcFunc(cmd *cobra.Command, args []string) {
	a := ackrc.New(name, source)

	err := a.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
