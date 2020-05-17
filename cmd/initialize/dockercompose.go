package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/gomicro/micro/dockercompose"
)

func init() {
	InitializeCmd.AddCommand(DockercomposeCmd)
	funcs = append(funcs, dockercomposeFunc)
}

// DockercomposeCmd represents the command that executes all of tasks for
// bootstrapping a Dockercompose
var DockercomposeCmd = &cobra.Command{
	Use:   "dockercompose",
	Short: "Initialize a docker-compose file",
	Run:   dockercomposeFunc,
}

func dockercomposeFunc(cmd *cobra.Command, args []string) {
	d := dockercompose.New(org, name)
	if database {
		d.EnableDB()
	}

	err := d.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
