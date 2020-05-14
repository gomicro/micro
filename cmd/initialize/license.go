package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/service/license"
)

func init() {
	InitializeCmd.AddCommand(LicenseCmd)
	funcs = append(funcs, licenseFunc)
}

// LicenseCmd represents the command that executes all of tasks for
// bootstrapping a License
var LicenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Initialize a License file",
	Run:   licenseFunc,
}

func licenseFunc(cmd *cobra.Command, args []string) {
	l := license.New(org)
	err := l.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
