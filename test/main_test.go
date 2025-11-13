package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

const terraformDir = "../testing"

func TestMain(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}

	terraformOptions := createTerraformOptions(config)
	defer terraform.Destroy(t, terraformOptions)

	// Run main test stages
	t.Run("infrastructure_deployment", func(t *testing.T) {
		testInfrastructureDeployment(t, terraformOptions)
	})
}

func createTerraformOptions(config *Config) *terraform.Options {
	return &terraform.Options{
		TerraformDir: terraformDir,
		MaxRetries:   0,
		Vars: map[string]interface{}{
			"test_author": fmt.Sprintf("ci-%s", config.Username),
			"test_id":     config.ShortSha,
			"aws_region":  config.AwsRegion,
			"aws_role_arn": config.AwsRoleArn,
			"dd_api_key":  config.DdApiKey,
			"dd_app_key":  config.DdAppKey,
		},
	}
}

func testInfrastructureDeployment(t *testing.T, terraformOptions *terraform.Options) {
	t.Run("initial_apply", func(t *testing.T) {
		terraform.InitAndApply(t, terraformOptions)
	})

	t.Run("idempotency", func(t *testing.T) {
		terraform.ApplyAndIdempotent(t, terraformOptions)
	})
}
