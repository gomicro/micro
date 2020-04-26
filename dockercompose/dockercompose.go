package dockercompose

import (
	"fmt"
	"os"
	"text/template"
)

const (
	outputFilename = "docker-compose.yml"
	errorKey       = "dockercomposefile"
)

// DockerCompose  is a struct representation of a docker compose template
type DockerCompose struct {
	Org      string
	Name     string
	Database bool
}

// New takes a name and returns a viable dockercompose struct
func New(org, name string) *DockerCompose {
	return &DockerCompose{
		Org:  org,
		Name: name,
	}
}

// EnableDB turns on the inclusion of database usage in the output template
func (d *DockerCompose) EnableDB() {
	d.Database = true
}

// WriteFile writes the current dockercompose file struct to an actualized
// docker-compose file. It returns any errors it encounters with the template
// and writing the file to disk.
func (d *DockerCompose) WriteFile() error {
	tmpl, err := template.New("").Parse(stub)
	if err != nil {
		return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
	}

	fh, err := os.OpenFile(outputFilename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%v: cannot open file: %v", errorKey, err.Error())
	}
	defer fh.Close()

	err = tmpl.Execute(fh, d)
	if err != nil {
		return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
	}

	envFh, err := os.Create(".env.test")
	if err != nil {
		return fmt.Errorf("%v: cannot create file: %v", errorKey, err.Error())
	}
	defer envFh.Close()

	return nil
}
