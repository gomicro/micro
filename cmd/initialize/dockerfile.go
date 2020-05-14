package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/service/dockerfile"
)

func init() {
	InitializeCmd.AddCommand(DockerfileCmd)
	funcs = append(funcs, dockerfileFunc)
}

// DockerfileCmd represents the command that executes all of tasks for
// bootstrapping a Dockerfile
var DockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Generate a bootstrap of a Dockerfile",
	Run:   dockerfileFunc,
}

func dockerfileFunc(cmd *cobra.Command, args []string) {
	d := dockerfile.New()
	err := d.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
