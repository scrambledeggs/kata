SERVICE_NAME = {{ cookiecutter.project_name }}
AWS_REGION = ap-southeast-1

.PHONY: generate-secrets
generate-secrets:
	go run scripts/generate-secrets/main.go -env=$(ENV)

.PHONY: generate-endpoints
generate-endpoints: build
	go run scripts/generate-endpoints/main.go -path=.aws-sam/build/AllEndpoints/endpoints.yml

.PHONY: build
build:
	sam build

.PHONY: deploy
deploy: generate-secrets generate-endpoints
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.json); \
	sam deploy --parameter-overrides "$$envs" --region $(AWS_REGION) --stack-name $(SERVICE_NAME)-$(ENV) --no-confirm-changeset

.PHONY: dev
dev: generate-endpoints
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.local.json); \
	sam local start-api --port 3000 --parameter-overrides "$$envs"

.PHONY: dev-watch
dev-watch: generate-endpoints
	@envs=$$(go run scripts/generate-parameter-overrides/main.go -path=.secrets.local.json); \
	(air &) && sam local start-api --port 3000 --parameter-overrides "$$envs"

.PHONY: build-watch
build-watch:
	@endpoints=$$(cat .aws-sam/build/AllEndpoints/endpoints.yml); \
	sam build; \
	echo "$$endpoints" > .aws-sam/build/AllEndpoints/endpoints.yml
