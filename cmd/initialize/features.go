package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/micro/features"
)

func init() {
	InitializeCmd.AddCommand(FeaturesCmd)
	funcs = append(funcs, featuresFunc)
}

// FeaturesCmd represents the command that executes all of tasks for
// bootstrapping the functional tests
var FeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "Initialize functional tests",
	Run:   featuresFunc,
}

func featuresFunc(cmd *cobra.Command, args []string) {
	f := features.New(name)

	err := f.WriteFiles()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
