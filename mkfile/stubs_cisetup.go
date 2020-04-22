package mkfile

const cisetupStub = `.PHONY: ci_setup
ci_setup: ## Setup the ci environment
	@wget -O flow.tar.gz https://github.com/gomicro/flow/releases/latest/download/flow_linux_amd64.tar.gz
	@tar xvf flow.tar.gz -C /home/travis/.local/bin`
