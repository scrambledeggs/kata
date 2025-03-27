## ----------------------------------------------------------------------
## {{ cookiecutter.project_name }} Makefile documentation
# This is a regular comment, that will not be displayed
## This is a help comment. Use double sharp for `make help` to output.
## Use spaces (not tabs) to align the commands and help comment properly.
## ----------------------------------------------------------------------

PROJECT_NAME = {{ cookiecutter.project_name }}
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
	sam deploy --parameter-overrides "$$envs" --region $(AWS_REGION) --stack-name $(PROJECT_NAME)-$(ENV) --no-confirm-changeset

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
	sam deploy --parameter-overrides "$$envs" --region $(AWS_REGION) --stack-name $(PROJECT_NAME)-$(ENV) --no-confirm-changeset

.PHONY: dev-invoke
dev-invoke:                                    ## Invoke one-time a lambda NAME function with data from the EVENT (optional)
	@if [ -z "$(NAME)" ]; then echo "Error: Please provide NAME of lambda to invoke."; exit 1; fi
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.local.json); \
	sam local invoke $(NAME) --parameter-overrides "$$envs"  --event $(EVENT)

.PHONY: build-watch
build-watch:                                   ## Build command executed by air.toml every live reload
	@endpoints=$$(cat .aws-sam/build/AllEndpoints/endpoints.yml); \
	sam build; \
	echo "$$endpoints" > .aws-sam/build/AllEndpoints/endpoints.yml

{% if cookiecutter.with_own_db == "y" -%}
.PHONY: db-add
db-add:                                        ## Add a new migration file opt. NAME=new_migration_file
	@if [ -z "$(NAME)" ]; then echo "Error: Please provide NAME for the migration file."; exit 1; fi
{%- if cookiecutter.db_migration == "goose" %}
	goose -dir db/migrations create $(NAME) sql
{%- elif cookiecutter.db_migration == "dbmate" %}
	dbmate --migrations-dir db/migrations new $(NAME)
{%- endif %}

.PHONY: db-up
db-up:                                         ## Migrate the DB to the most recent version available
{%- if cookiecutter.db_migration == "goose" %}
	GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=$(shell cat .secrets.local.json | sed -n 's/.*"DatabaseURL": *"\([^"]*\)".*/\1/p') \
	goose -dir db/migrations up
{%- elif cookiecutter.db_migration == "dbmate" %}
	DATABASE_URL=$(shell cat .secrets.local.json | sed -n 's/.*"DatabaseURL": *"\([^"]*\)".*/\1/p') \
	dbmate --migrations-dir db/migrations --no-dump-schema migrate
{%- endif %}

.PHONY: db-down
db-down:                                       ## Roll back the migration version by 1
{%- if cookiecutter.db_migration == "goose" %}
	GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=$(shell cat .secrets.local.json | sed -n 's/.*"DatabaseURL": *"\([^"]*\)".*/\1/p') \
	goose -dir db/migrations down
{%- elif cookiecutter.db_migration == "dbmate" %}
	DATABASE_URL=$(shell cat .secrets.local.json | sed -n 's/.*"DatabaseURL": *"\([^"]*\)".*/\1/p') \
	dbmate --migrations-dir db/migrations --no-dump-schema down
{%- endif %}
{%- endif %}

{% if cookiecutter.use_sqlc -%}
.PHONY: sqlc-gen
sqlc-gen:                                      ## Generate sqlc code in sqlc.yaml
	sqlc generate
{%- endif %}
