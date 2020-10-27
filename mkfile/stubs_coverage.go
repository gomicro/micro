package mkfile

const coverageStub = `.PHONY: coverage
coverage: ## Generates the code coverage from all the tests
	@docker run -v $$PWD:/go$${PWD/$$GOPATH} --workdir /go$${PWD/$$GOPATH} gomicro/gocover`
