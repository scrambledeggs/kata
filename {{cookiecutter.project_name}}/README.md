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

1. This is provided with .secrets.local.json and has empty ENV variables

### Environment

1. Default environment handler is in [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/). `generate-secrets` downloads the environment and generates .secrets.json
2. Format is in JSON format, keys should be in CamelCase (sample data)
```json
{
  "AlternateDomainName": "test-{{ cookiecutter.project_name }}.booky.ph",
  "AppEnv": "test",
  "CertificateArn": "arn:aws:acm:us-east-1:123456789012:certificate/5da5njye-this-test-yeah-1eb057af5006",
  "DistributionId": "E3TESTX633ONLY",
  "SecurityGroups": "sg-test32bae2b93test,sg-test32bae2b93just",
  "Subnets": "subnet-test32bae2b93test"
}
```
3. This is then read by `Makefile` using `generate-parameter-overrides`
4. Development `make dev` uses `.secrets.local.json`
    4.1. Optionally use `make dev-watch` to use live-reloading while developing
5. Deployment (`make deploy`) uses (and builds) using `.secrets.json`

### Local development

**Invoking function locally through local API Gateway**

1. Make sure you have `.secrets.local.json` file
2. Run the following in your shell:
```bash
$ make dev
```
  2.1. Optionally use `make dev-watch` for live-reloading while developing

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/v1/hello-world`

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

```yaml
...
Events:
    HelloWorldV1:
        Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
        Properties:
            Path: /v1/hello-world
            Method: GET
```

## Deployment

### CLI
To deploy your application for the first time,

1. Make sure you have secrets manager named under (env-){{ cookiecutter.project_name }}
    - This would be used when generate-secrets executes
2. Run the following in your shell:
```bash
$ make deploy ENV=test
```

### Github actions
1. Actions -> Start Deployment
2. Run workflow
- Choose branch [main]
- Deploy to TEST Env
- Run workflow

### Gacha to make CloudFront working for now
1. After deploying, copy Output->CloudFrontDistributionId to secrets->DistributionId, redeploy
2. Create Route53 -> Hosted Zone -> Record
- Choose Route53
- Hosted Zones
- booky.ph
- Create Record
    - Record name: [env]-[{{ cookiecutter.project_name }}]
    - Record type: CNAME
    - Value: Output->CloudFrontDistributionDomainName
    - TTL: 300
    - Routing Policy: Simple

## Testing

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
