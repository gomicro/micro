package mkfile

import (
	"fmt"
	"os"
	"text/template"
)

const (
	outputFilename = "Makefile"
	errorKey       = "mkfile"
)

// Makefile is a struct representation of a Makefile template
type Makefile struct {
	Org  string
	Name string
}

// New takes an org and name of a project to properly construct a new Makefile
// template.
func New(org, name string) *Makefile {
	return &Makefile{
		Org:  org,
		Name: name,
	}
}

// WriteFile writes the current makefile file struct to an actualized Makefile
// file.  It returns any errors it encounters with the template and writing the
// file to disk.
func (m *Makefile) WriteFile() error {
	tmpl, err := template.New("makefile").Parse(stub)
	if err != nil {
		return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
	}

	for name, s := range stubs {
		tmpl.New(name).Parse(s)
		if err != nil {
			return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
		}
	}

	fh, err := os.OpenFile(outputFilename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%v: cannot open file: %v", errorKey, err.Error())
	}
	defer fh.Close()

	err = tmpl.ExecuteTemplate(fh, "makefile", m)
	if err != nil {
		return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
	}

	return nil
}
