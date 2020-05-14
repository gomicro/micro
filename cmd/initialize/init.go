package initialize

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	funcs []func(cmd *cobra.Command, args []string)

	// Flag vars
	org         string
	name        string
	source      string
	database    bool
	installable bool
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defaultName := ""
	defaultOrg := ""
	defaultSource := ""

	pathParts := strings.Split(wd, "/github.com/")
	if len(pathParts) > 1 {
		repoParts := strings.Split(pathParts[len(pathParts)-1], "/")
		if len(repoParts) > 1 {
			defaultOrg = repoParts[0]
			defaultName = repoParts[1]
			defaultSource = fmt.Sprintf("https://github.com/%v/%v", defaultOrg, defaultName)
		}

	}

	InitializeCmd.Flags().StringVar(&org, "org", defaultOrg, "organization name")
	InitializeCmd.Flags().StringVar(&name, "name", defaultName, "service name")
	InitializeCmd.Flags().StringVar(&source, "source", defaultSource, "source location")
	InitializeCmd.Flags().BoolVar(&database, "db", false, "whether the service will have a database or not")
	InitializeCmd.Flags().BoolVar(&installable, "installable", false, "whether the service will be installable or not")
}

// InitializeCmd executes all of the initialize commands together rather than as
// individual commands.
var InitializeCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new microservice",
	Run:   initializeFunc,
}

func initializeFunc(cmd *cobra.Command, args []string) {
	for i := range funcs {
		funcs[i](cmd, args)
	}
}
