package license

import (
	"testing"
	"text/template"
)

func TestLicenseTemplateCompiles(t *testing.T) {
	for name, stub := range stubs {
		_, err := template.New("").Parse(stub)
		if err != nil {
			t.Errorf("expected to parse: template %v", name)
		}
	}
}

func TestGetStub(t *testing.T) {
	stub, err := getStub("MIT")
	if err != nil {
		t.Errorf("expected MIT stub\ngot '%v'", stub)
	}

	stub, err = getStub("mit")
	if err != nil {
		t.Errorf("expected MIT stub\ngot '%v'", stub)
	}

	_, err = getStub("nonsense")
	if err == nil {
		t.Error("expected: error\ngot nil")
	}
}
