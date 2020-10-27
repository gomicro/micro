package mkfile

const dockerizeStub = `.PHONY: dockerize
dockerize: ecr_login  ## Create a docker image of the project
	@docker build \
	--build-arg BUILD_PATH=/go$${PWD/$$GOPATH} \
	-t {{ .Org }}/{{ .Name }} .`
