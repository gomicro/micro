package license

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
	"time"
)

const (
	outputFilename = "LICENSE.md"
	errorKey       = "license"
)

// License is a struct representation of a license template
type License struct {
	Year string
	Name string
}

// New takes a name and returns a viable license struct
func New(name string) *License {
	return &License{
		Year: strconv.FormatInt(int64(time.Now().Year()), 10),
		Name: name,
	}
}

// WriteFile writes the current license file struct to an actualized LICENSE
// file. It returns any errors it encounters with the template and writing the
// file to disk.
func (l *License) WriteFile() error {
	stub, err := getStub("MIT")
	if err != nil {
		return err
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

	err = tmpl.Execute(fh, l)
	if err != nil {
		return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
	}

	return nil
}
