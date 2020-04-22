package mkfile

const cleanStub = `.PHONY: clean
clean: ## Cleans out all generated items
	-@rm -f output.txt
	-@rm -rf coverage
	-@rm -f coverage.txt
	-@docker-compose down`
