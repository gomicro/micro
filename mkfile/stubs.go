package mkfile

var stubs = map[string]string{
	"cisetup":   cisetupStub,
	"clean":     cleanStub,
	"coverage":  coverageStub,
	"dockerize": dockerizeStub,
	"ecrlogin":  ecrloginStub,
	"help":      helpStub,
	"linters":   lintersStub,
	"pulldeps":  pulldepsStub,
	"run":       runStub,
	"test":      testStub,
	"vars":      varsStub,
}

var stub = `{{ template "vars" . }}

.PHONY: all
all: test

{{ template "cisetup" . }}

{{ template "clean" . }}

{{ template "coverage" . }}

{{ template "dockerize" . }}

{{ template "ecrlogin" . }}

{{ template "help" . }}

{{ template "linters" . }}

{{ template "pulldeps" . }}

{{ template "run" . }}

{{ template "test" . }}
`
