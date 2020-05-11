package features

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const (
	outputFilename = "status.feature"
	errorKey       = "features"
)

type Features struct {
}

// New takes a name and returns a viable features struct
func New(name string) *Features {
	return &Features{}
}

// WriteFiles writes the current service struct to actualized feature files. It
// returns any errors it encounters with the template and writing the file to
// disk.
func (f *Features) WriteFiles() error {
	for outputFilename, stub := range stubs {
		dirName := filepath.Dir(outputFilename)

		_, err := os.Stat(dirName)
		if err != nil {
			err = os.MkdirAll(dirName, os.ModePerm)
			if err != nil {
				return fmt.Errorf("%v: cannot make directory for file: %v", errorKey, err.Error())
			}
		}

		tmpl, err := template.New("").Parse(stub)
		if err != nil {
			return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
		}

		fh, err := os.OpenFile(outputFilename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("%v: cannot open file: %v", errorKey, err.Error())
		}
		defer fh.Close()

		err = tmpl.Execute(fh, f)
		if err != nil {
			return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
		}
	}

	return nil
}
