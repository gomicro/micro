package ackrc

import (
	"fmt"
	"os"
	"text/template"
)

const (
	outputFilename = ".ackrc"
	errorKey       = "ackrc"
)

// Ackrc is a struct representation of a ackrc template
type Ackrc struct {
}

// New takes a name and returns a viable ackrc struct
func New(name, source string) *Ackrc {
	return &Ackrc{}
}

// WriteFile writes the current ackrc file struct to an actualized .ackrc file.
// It returns any errors it encounters with the template and writing the file to
// disk.
func (r *Ackrc) WriteFile() error {
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
