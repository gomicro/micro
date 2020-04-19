package gitignore

import (
	"fmt"
	"os"
	"text/template"
)

const (
	outputFilename = ".gitignore"
	errorKey       = "gitignore"
	stub           = "{{ range .Ignores }}{{ . }}\n{{ end }}"
)

// Gitignore is a struct representation of a gitignore template
type Gitignore struct {
	Ignores []string
}

// New takes a name and returns a viable license struct
func New(ignores ...string) *Gitignore {
	return &Gitignore{
		Ignores: ignores,
	}
}

// AddIgnore takes an addtional file to include in the ignore stuct, and adds it
// to the list of files that will go into the generated template.
func (g *Gitignore) AddIgnore(ignore string) {
	g.Ignores = append(g.Ignores, ignore)
}

// WriteFile writes the current gitignore file struct to an actualized
// .gitignore file. It returns any errors it encounters with the template and
// writing the file to disk.
func (g *Gitignore) WriteFile() error {
	tmpl, err := template.New("").Parse(stub)
	if err != nil {
		return fmt.Errorf("%v: cannot parse template: %v", errorKey, err.Error())
	}

	fh, err := os.OpenFile(outputFilename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%v: cannot open file: %v", errorKey, err.Error())
	}
	defer fh.Close()

	err = tmpl.Execute(fh, g)
	if err != nil {
		return fmt.Errorf("%v: template failed execution: %v", errorKey, err.Error())
	}

	return nil
}
