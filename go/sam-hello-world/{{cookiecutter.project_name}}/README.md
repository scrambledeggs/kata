# {{ cookiecutter.project_name }}

This is a template for {{ cookiecutter.project_name }} - Below is a brief explanation of what we have generated for you:

```bash
.
├── db                            <-- Database folders
|   └── migrations                <-- Migration files by goose (or choose your poison)
├── events                        <-- Contains sample events for invoking the lambda function
├── handlers                      <-- Source code for lambda functions
│   └── ActionResourceV1          <-- Lambda function name
│      ├── main_test.go           <-- Lambda function unit test
|      └── main.go                <-- Lambda function code
├── scripts                       <-- Go code that executed separately
├── .secrets.local.json           <-- (gitignored) holds the ENV secrets json
├── go.mod                        <-- dependency manager
├── go.sum                        <-- modules
├── Makefile                      <-- Make to automate build
├── README.md                     <-- You are here
├── samconfig.toml                <-- Local Deployment Script to AWS
└── template.yaml                 <-- AWS SAM template file for building the infrastructure
```

## Requirements

* AWS CLI already configured with permissions
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process

1. Run the following commands
```bash
$ go mod init [project-name]
$ go mod tidy
```

### Local development

**Invoking function locally through local API Gateway**

1. make sure you have `.secrets.local.json` file
2. then run the following in your shell:
```bash
$ make dev
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/hello-world`

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

```yaml
...
Events:
    HelloWorldV1:
        Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
        Properties:
            Path: /hello-world
            Method: GET
```

## Deployment

To deploy your application for the first time,

1. make sure you have `.secrets.json` file
    - you can copy .secrets.local.json
2. then run the following in your shell:

```bash
$ make deploy
```

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
$ go test -v .
```

# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
$ brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
$ brew update
$ brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
$ choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
$ choco upgrade golang
```
