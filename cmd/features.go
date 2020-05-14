package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/service/features"
)

func init() {
	RootCmd.AddCommand(FeaturesCmd)
	funcs = append(funcs, featuresFunc)
}

// FeaturesCmd represents the command that executes all of tasks for
// bootstrapping the functional tests
var FeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "Generate a bootstrap of functional tests",
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
