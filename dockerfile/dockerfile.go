package dockerfile

import (
	"fmt"
	"os"
	"text/template"
)

const (
	outputFilename = "Dockerfile"
	errorKey       = "dockerfile"
)

// Dockerfile is a struct representation of a dockerfile template
type Dockerfile struct {
}

// New takes a name and returns a viable dockerfile struct
func New() *Dockerfile {
	return &Dockerfile{}
}

// WriteFile writes the current dockerfile file struct to an actualized
// Dockerfile file. It returns any errors it encounters with the template and
// writing the file to disk.
func (d *Dockerfile) WriteFile() error {
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

	return nil
}
