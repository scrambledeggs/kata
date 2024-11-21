## ----------------------------------------------------------------------
## {{ cookiecutter.project_name }} Makefile documentation
# This is a regular comment, that will not be displayed
## This is a help comment. Use double sharp for `make help` to output.
## Use spaces (not tabs) to align the commands and help comment properly.
## ----------------------------------------------------------------------

SERVICE_NAME = {{ cookiecutter.project_name }}
AWS_REGION = ap-southeast-1

help:                                          ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

.PHONY: generate-secrets
generate-secrets:                              ## Generate secrets file from AWS Secrets Manager based on ENV
	go run scripts/generate-secrets/main.go -env=$(ENV)

.PHONY: generate-endpoints
generate-endpoints: build                      ## Generate endpoints file based on docs/api_contract.yaml
	go run scripts/generate-endpoints/main.go -path=.aws-sam/build/AllEndpoints/endpoints.yml

.PHONY: build
build:                                         ## Build project in to .aws-sam
	sam build

.PHONY: deploy
deploy: generate-secrets generate-endpoints    ## Deploy project, will create parameter-overrides based on .secrets.json taken from AWS Secrets Manager
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.json); \
	sam deploy --parameter-overrides "$$envs" --region $(AWS_REGION) --stack-name $(SERVICE_NAME)-$(ENV) --no-confirm-changeset

.PHONY: dev
dev: generate-endpoints                        ## Build and run locally
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.local.json); \
	sam local start-api --port 3000 --parameter-overrides "$$envs"

.PHONY: dev-watch
dev-watch: generate-endpoints                  ## Build and run locally. Will maintain a watcher for changes and live reload the service
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.local.json); \
	(air &) && sam local start-api --port 3000 --parameter-overrides "$$envs"

.PHONY: dev-deploy
dev-deploy: generate-endpoints                 ## Deploy project, will create parameter-overrides based on .secrets.local.json
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.local.json); \
	sam deploy --parameter-overrides "$$envs" --region $(AWS_REGION) --stack-name $(SERVICE_NAME)-$(ENV) --no-confirm-changeset

.PHONY: build-watch
build-watch:                                   ## Build command executed by air.toml every live reload
	@endpoints=$$(cat .aws-sam/build/AllEndpoints/endpoints.yml); \
	sam build; \
	echo "$$endpoints" > .aws-sam/build/AllEndpoints/endpoints.yml
