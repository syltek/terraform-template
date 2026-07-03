# Testing

This folder contains the testing setup using [Terratest](https://terratest.gruntwork.io/) for this Terraform module.

## Prerequisites

- Go ~> 1.25.3
- Terraform ~> 1.9.0
- Access to the private [`syltek/terratest-commons`](https://github.com/syltek/terratest-commons) Go module:

```bash
go env -w GOPRIVATE=github.com/syltek/*
gh auth setup-git   # or: git config --global url."git@github.com:syltek/".insteadOf "https://github.com/syltek/"
```

## Test Structure

The test setup includes:

1. A main test file (`main_test.go`) that deploys `../testing` and checks the apply is idempotent
2. Configuration and deployment scaffolding from the shared [`syltek/terratest-commons`](https://github.com/syltek/terratest-commons) library

Adapt `main_test.go` to your module: pass `config.WithAWS()` and/or `config.WithDatadog()` to `config.Load` to match the variables `../testing/test.tf` declares, add module-specific variables via the `extraVars` argument of `tftest.Options`, and add output-validation subtests before the deferred destroy.

## Running Tests

1. Make sure you have the required environment variables:
   - `TF_VAR_test_author` or `GITHUB_ACTOR`: Test author name
   - `TF_VAR_test_id` or `GITHUB_SHA`: Test ID
   - With `config.WithAWS()`: `TF_VAR_aws_region`/`AWS_REGION` and `TF_VAR_aws_role_arn`/`AWS_ROLE_ARN`
   - With `config.WithDatadog()`: `TF_VAR_dd_api_key`/`TERRAFORM_DD_API_KEY` and `TF_VAR_dd_app_key`/`TERRAFORM_DD_APP_KEY`

2. Run the tests:

```bash
cd test && go test -v
```
