package test

import (
	"testing"

	"github.com/syltek/terratest-commons/config"
	"github.com/syltek/terratest-commons/tftest"
)

func TestMain(t *testing.T) {
	// Add config.WithAWS() and/or config.WithDatadog() to match the variables
	// your testing fixture declares (see ../testing/test.tf).
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}

	terraformOptions := tftest.Options(cfg, "../testing", nil)
	defer tftest.Destroy(t, terraformOptions)

	// Run main test stages
	t.Run("infrastructure_deployment", func(t *testing.T) {
		tftest.Deploy(t, terraformOptions)
	})

	// Add output-validation subtests here, before the deferred destroy.
}
