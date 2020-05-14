package mkfile

const testStub = `.PHONY: test
test: unit_test functional_test ## Run all available tests

.PHONY: functional_test
functional_test: ## Run the functional tests against the running service
	docker run -i $(TTY) -v $$PWD/features:/usr/app/features --network=$(APP)_services gomicro/cucumber cucumber $$CUKE_TAGS

.PHONY: unit_test
unit_test: ## Run unit tests
	go test ./...`
