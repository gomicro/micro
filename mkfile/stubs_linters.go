package mkfile

const lintersStub = `.PHONY: linters
linters: ## Run all the linters
	docker run -v $$PWD:/go$${PWD/$$GOPATH} --workdir /go$${PWD/$$GOPATH} gomicro/golinters`
