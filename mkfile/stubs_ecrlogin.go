package mkfile

const ecrloginStub = `.PHONY: ecr_login
ecr_login:  ## Login to the ECR using local credentials
	@eval $$(flow aws ecr auth)`
