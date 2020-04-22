package mkfile

const testStub = `.PHONY: test
test: unit_test ## Run all available tests

.PHONY: unit_test
unit_test: ecr_login  ## Run unit tests
	go test ./...`
