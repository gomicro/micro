package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gomicro/service/gitignore"
	"github.com/gomicro/service/gofiles"
	"github.com/gomicro/service/license"
	"github.com/gomicro/service/readme"
)

var (
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

	RootCmd.Flags().StringVar(&org, "org", defaultOrg, "organization name")
	RootCmd.Flags().StringVar(&name, "name", defaultName, "service name")
	RootCmd.Flags().StringVar(&source, "source", defaultSource, "source location")
	RootCmd.Flags().BoolVar(&database, "db", false, "whether the service will have a database or not")
	RootCmd.Flags().BoolVar(&installable, "installable", false, "whether the service will be installable or not")
}

// RootCmd represents the command that executes all of tasks for bootstrapping a
// service
var RootCmd = &cobra.Command{
	Use:   "service",
	Short: "Generate a bootstrap of a new microservice",
	Run:   rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string) {
	l := license.New(org)
	err := l.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	r := readme.New(name, source)
	if installable {
		r.EnableInstallable()
	}

	err = r.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gi := gitignore.New(name)
	err = gi.WriteFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	gf := gofiles.New(name)
	if database {
		gf.EnableDB()
	}

	err = gf.WriteFiles()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Execute wraps the command executor to make it possible to trigger easily
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("Failed to execute: %v", err.Error())
		os.Exit(1)
	}
}
