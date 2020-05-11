package gofiles

import (
	"fmt"
	"os"
	"text/template"
)

const (
	errorKey = "gofiles"
)

// Service is a struct representation of the service templates
type Service struct {
	Database bool
	Name     string
}

// New returns a viable service struct
func New(name string) *Service {
	return &Service{
		Name: name,
	}
}

// EnableDB turns on the inclusion of database usage in the output template
func (s *Service) EnableDB() {
	s.Database = true
}

// WriteFiles writes the current service struct to actualized go files. It
// returns any errors it encounters with the template and writing the file to
// disk.
func (s *Service) WriteFiles() error {
	for outputFilename, stub := range stubs {
		tmpl, err := template.New("").Parse(stub)
		if err != nil {
			return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
		}

		fh, err := os.OpenFile(outputFilename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("%v: cannot open file: %v", errorKey, err.Error())
		}
		defer fh.Close()

		err = tmpl.Execute(fh, s)
		if err != nil {
			return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
		}
	}

	return nil
}
