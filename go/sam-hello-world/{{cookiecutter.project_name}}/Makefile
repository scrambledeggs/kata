SERVICE = {{ cookiecutter.project_name }}
AWS_REGION = ap-southeast-1

.PHONY: generate-endpoints
generate-endpoints: build
	go run scripts/generate-endpoints/main.go ./.aws-sam/build/AllEndpoints/endpoints.yml

.PHONY: build
build:
	sam build

.PHONY: deploy
deploy: generate-endpoints
	@envs=$$(go run scripts/generate-parameter-overrides/main.go .secrets.json); \
	sam deploy --parameter-overrides "$$envs" --region $(AWS_REGION) --stack-name $(SERVICE)-$(ENV) --no-confirm-changeset

.PHONY: dev
dev: generate-endpoints
	@envs=$$(go run scripts/generate-parameter-overrides/main.go .secrets.local.json); \
	sam local start-api --port 3000 --parameter-overrides "$$envs"
