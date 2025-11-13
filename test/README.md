# Testing

This folder contains the testing setup using [Terratest](https://terratest.gruntwork.io/) for the AWS Lambda Terraform module.

## Prerequisites

- AWS credentials configured
- Datadog API and Application keys
- Go ~> 1.23.3
- Terraform ~> 1.9.0

## Test Structure

The test setup includes:

1. A main test file (`main_test.go`) that validates:
   - Infrastructure deployment
   - Output validation
2. A configuration file (`config.go`) that handles environment variables

## Running Tests

1. Make sure you have the required environment variables:
   - `TF_VAR_test_author` or `GITHUB_ACTOR`: Test author name
   - `TF_VAR_test_id` or `GITHUB_SHA`: Test ID
   - `TF_VAR_aws_region` or `AWS_REGION`: AWS region
   - `TF_VAR_dd_api_key` or `TERRAFORM_DD_API_KEY`: Datadog API key
   - `TF_VAR_dd_app_key` or `TERRAFORM_DD_APP_KEY`: Datadog Application key

2. Run the tests:

```bash
cd test && go test -v
```
