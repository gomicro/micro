package mkfile

const dockerizeStub = `.PHONY: dockerize
dockerize: ecr_login  ## Create a docker image of the project
	docker build -t {{ .Org }}/{{ .Name }} .`
