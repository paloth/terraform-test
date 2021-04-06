package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-s3-example using Terratest.
func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	const (
		awsRegion          = "eu-west-1"
		expectedBucketName = "paloth-test-bucket."
		expectedStatus     = "Enabled"
	)

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../..",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketID := terraform.Output(t, terraformOptions, "bucket_name")

	assert.Equal(t, expectedBucketName, bucketID)

	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)

	assert.Equal(t, expectedStatus, actualStatus)
}
