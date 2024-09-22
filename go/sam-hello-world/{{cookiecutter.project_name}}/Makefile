.PHONY: build
build:
	sam build

.PHONY: deploy
deploy: build
	@envs=$$(go run cmd/parameter_overrides.go .secrets.json);
	sam deploy --parameter-overrides "$$envs"

.PHONY: dev
dev: build
	@envs=$$(go run cmd/parameter_overrides.go .secrets.local.json);
	sam local start-api --port 3000 --parameter-overrides "$$envs"
