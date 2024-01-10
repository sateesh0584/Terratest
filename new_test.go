package test

import (
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformExample(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "/home/ubuntu", // Replace with the actual path
		Vars: map[string]interface{}{
			// You can add any variables that need to be passed to your Terraform configuration here.
		},
	}

	// Defer the destruction of the infrastructure until the test is complete
	defer terraform.Destroy(t, terraformOptions)

	// Init and Apply Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Sleep for some time to allow the instance to be provisioned
	time.Sleep(5 * time.Minute) // Adjust the duration based on your instance provisioning time

	// Run `terraform output` to get the value of an output variable
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Verify the instance ID is as expected
	assert.True(t, strings.HasPrefix(instanceID, "i-"))
}

