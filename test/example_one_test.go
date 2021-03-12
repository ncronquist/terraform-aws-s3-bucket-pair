package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestS3BucketPair(t *testing.T) {
	workingDir := "../examples/example-one"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where the Terraform code is located
		TerraformDir: workingDir,
	})

	// Save the terraformOptions so other test stages can use them
	test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

	// At the end of the test, clean up all the resources we created
	defer test_structure.RunTestStage(t, "tearodown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.Destroy(t, terraformOptions)
	})

	// Deploy the infrastructure
	test_structure.RunTestStage(t, "deploy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "test", func() {
		testSomething(t, workingDir)
	})
}

func testSomething(t *testing.T, workingDir string) {
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	t.Logf("Working Directory %s", terraformOptions.TerraformDir)
	t.Log("Write a test here")
}
