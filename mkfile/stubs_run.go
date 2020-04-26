package mkfile

const runStub = `.PHONY: run
run: dockerize pull_dependencies ## Run a dockerized version of the app
	docker-compose up -d`
