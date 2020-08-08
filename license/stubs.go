package license

import (
	"fmt"
	"strings"
)

func getStub(name string) (string, error) {
	name = strings.ToUpper(name)
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ReplaceAll(name, ".", "")
	name = strings.ReplaceAll(name, " ", "")

	stub, found := stubs[name]
	if !found {
		return "", fmt.Errorf("license: license template not found")
	}

	return stub, nil
}

var (
	stubs = map[string]string{
		"APACHE20":        apache20,
		"APACHELICENSE20": apache20,
		"ISC":             isc,
		"MIT":             mit,
	}
)
