package readme

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const (
	outputFilename = "README.md"
	errorKey       = "readme"
)

// Readme is a struct representation of a readme template
type Readme struct {
	Name        string
	Source      string
	Installable bool
}

// New takes a name and returns a viable readme struct
func New(name, source string) *Readme {
	return &Readme{
		Name:   name,
		Source: source,
	}
}

// EnableInstallable enables the installable details of the readme to be turned
// on.
func (r *Readme) EnableInstallable() {
	r.Installable = true
}

// ReleaseURL returns a string URL of the releases location if there is known to
// be one.
func (r *Readme) ReleaseURL() string {
	if strings.Contains(r.Source, "github.com") {
		return fmt.Sprintf("%v/releases", r.Source)
	}

	return ""
}

// WriteFile writes the current readme file struct to an actualized README file.
// It returns any errors it encounters with the template and writing the file to
// disk.
func (r *Readme) WriteFile() error {
	tmpl, err := template.New("").Parse(stub)
	if err != nil {
		return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
	}

	fh, err := os.OpenFile(outputFilename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%v: cannot open file: %v", errorKey, err.Error())
	}
	defer fh.Close()

	err = tmpl.Execute(fh, r)
	if err != nil {
		return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
	}

	return nil
}
