package test

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	Username  string // The GitHub user name
	ShortSha  string // The short Git commit hash
	AwsRegion string // The AWS region
	AwsRoleArn string // The AWS role ARN
	DdApiKey    string // The Datadog API key
	DdAppKey    string // The Datadog Application key
}

// GetConfig returns a Config object with the values from the environment variables
// It first checks local test environment variables (TF_VAR_test_author, TF_VAR_test_id, TF_VAR_aws_region),
// then falls back to global environment variables (GITHUB_ACTOR, GITHUB_SHA, AWS_REGION).
// If any of the required values are missing, it returns an error.
func GetConfig() (*Config, error) {
	var errstrings []string
	var config Config
	var found bool

	// First, check local test variables
	// Local test variables that might be set (TF_VAR_test_author, TF_VAR_test_id, TF_VAR_aws_region)
	config.Username, found = os.LookupEnv("TF_VAR_test_author")
	if !found || config.Username == "" {
		config.Username, found = os.LookupEnv("GITHUB_ACTOR")
		if !found || config.Username == "" {
			errstrings = append(errstrings, "Username (either TF_VAR_test_author or GITHUB_ACTOR) must be set")
		}
	}

	config.ShortSha, found = os.LookupEnv("TF_VAR_test_id")
	if !found || config.ShortSha == "" {
		config.ShortSha, found = os.LookupEnv("GITHUB_SHA")
		if !found || config.ShortSha == "" {
			errstrings = append(errstrings, "ShortSha (either TF_VAR_test_id or GITHUB_SHA) must be set")
		}
	}

	// Only trim if the short SHA is longer than 8 characters
	if len(config.ShortSha) > 8 {
		config.ShortSha = config.ShortSha[:8]
	}

	config.AwsRegion, found = os.LookupEnv("TF_VAR_aws_region")
	if !found || config.AwsRegion == "" {
		config.AwsRegion, found = os.LookupEnv("AWS_REGION")
		if !found || config.AwsRegion == "" {
			errstrings = append(errstrings, "AwsRegion (either TF_VAR_aws_region or AWS_REGION) must be set")
		}
	}

	config.DdApiKey, found = os.LookupEnv("TF_VAR_dd_api_key")
	if !found || config.DdApiKey == "" {
		config.DdApiKey, found = os.LookupEnv("TERRAFORM_DD_API_KEY")
		if !found || config.DdApiKey == "" {
			errstrings = append(errstrings, "DdApiKey (either TF_VAR_dd_api_key or TERRAFORM_DD_API_KEY) must be set")
		}
	}

	config.DdAppKey, found = os.LookupEnv("TF_VAR_dd_app_key")
	if !found || config.DdAppKey == "" {
		config.DdAppKey, found = os.LookupEnv("TERRAFORM_DD_APP_KEY")
		if !found || config.DdAppKey == "" {
			errstrings = append(errstrings, "DdAppKey (either TF_VAR_dd_app_key or TERRAFORM_DD_APP_KEY) must be set")
		}
	}

	config.AwsRoleArn, found = os.LookupEnv("TF_VAR_aws_role_arn")
	if !found || config.AwsRoleArn == "" {
		config.AwsRoleArn, found = os.LookupEnv("AWS_ROLE_ARN")
		if !found || config.AwsRoleArn == "" {
			errstrings = append(errstrings, "AwsRoleArn (either TF_VAR_aws_role_arn or AWS_ROLE_ARN) must be set")
		}
	}

	// If we have any missing variables, return an error
	if len(errstrings) > 0 {
		return nil, errors.New(strings.Join(errstrings, ", "))
	}

	return &config, nil
}
