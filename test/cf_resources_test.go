package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/sbalnojan/cf-test/modules/cloudformation"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-example using Terratest.
func TestCFResources(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

	CFOptions := &cloudformation.Options{
		CFFile:    "../examples/cloudformation-aws-example/cf_create_test.yml",
		StackName: expectedName,
		AWSRegion: "us-east-1",
	}
	defer cloudformation.DeleteStack(t, CFOptions)

	cloudformation.CreateStack(t, CFOptions)
	list := cloudformation.ListResources(t, CFOptions)
	filteredList := cloudformation.FilterResources(list, "DummyResource")
	assert.Contains(t, filteredList, "cloudformation-waitcondition-")

}
