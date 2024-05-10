# Go Certificate

![Layered Architecture](https://github.com/malaquiasdev/ekoa-certificate-generator/blob/main/doc/dynamodb/architecture/ekoa-certificate-generator.drawio.png?raw=true)

## Setup

### Prerequisites 📝

Before you begin, ensure you have met the following requirements:

- You must have an AWS Credentials
- Configure the [AWS CLI](https://aws.amazon.com/pt/cli/)
- You have installed the [Golang](https://go.dev/)
- You have installed the [Terraform](https://www.terraform.io)

### Create the infra on AWS 🏗️

[Go to Terraform doc](terraform/readme.md)

## Running local

### Create configs

Create a `.env` file at the root of the project. Make sure you follow the [`.env.example`](.env.example) file as a guide.

### Create the infra on AWS 🏗️

[Go to Terraform doc](terraform/readme.md)

### Run project

```sh
$ go run cmd/main.go
```

## DynamoDB Schema

[Go to DynamoDB doc](doc/dynamodb/README.md)
