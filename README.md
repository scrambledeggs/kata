# Kata - Repository of templates

~~This repository contains the application templates.~~

This repo currently only contains a template for [AWS SAM](https://aws.amazon.com/serverless/sam/) in [golang](https://go.dev/).

## Features

### HelloWorldV1 - the sample endpoint

Template has a basic handler HelloWorld under `/handlers/HelloWorldV1` accessible under `/v1/hello-world`

### API Gateway

Template has an API Gateway included that points to (test-/staging-)[project_name] upon deployment

### docs/api_contract

Template has a basic setup of OpenAPI documentation

### Makefile

Template has a basic Makefile setup that I've tried executing while making the template

Execute this makefile command for local development
```
$ make dev
```

Execute this makefile command to deploy project
```
$ make deploy ENV=test
```

### github actions

Template has a ready github actions that can deploy already, just adjust depending on your projects needs.

### air

Template implements [air](https://github.com/air-verse/air) for live-reloading, this works in the background

Execute this makefile command for local development with air live-reloading
```
$ make dev-watch
```

### AllEndpoints - the /endpoints endpoint

Template already has a ready endpoint for our Booky auth under AllEndpoints handler

### scripts/generate-endpoints

Creates an endpoint (BOOKY endpoint list) for /endpoints endpoint to read. Uses `docs/api_contract.yaml` to generate list of endpoints

### scripts/generate-secret

Reads from (env-)[project_name](by default) in [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/) and creates a file .secrets.json

### scripts/generate-parameter-overrides

Reads .secrets.json (by default) and creates a parameter-override string for `sam deploy` or sam local start-api`


## Single repo reason

Since SAM can't have multiple template yet in one repo as of this writing, I opted to put this under root so it can work this way for now https://github.com/aws/aws-sam-cli/issues/3555.

## Contributing

We welcome issue reports and pull requests to help improve these application templates.
