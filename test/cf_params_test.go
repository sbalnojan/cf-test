package test

import (
	"fmt"
	"testing"

	aws_cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/sbalnojan/cf-test/modules/cloudformation"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-example using Terratest.
func TestCFParams(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

	test := "TestParameter"
	TestParams := []*aws_cf.Parameter{&aws_cf.Parameter{
		ParameterKey:   &test,
		ParameterValue: &test,
	},
	}

	CFOptions := &cloudformation.Options{
		CFFile:     "../examples/cloudformation-aws-example/cf_create_params_test.yml",
		StackName:  expectedName,
		AWSRegion:  "us-west-1",
		Parameters: TestParams,
	}
	defer cloudformation.DeleteStack(t, CFOptions)

	cloudformation.CreateStack(t, CFOptions)
	list := cloudformation.ListResources(t, CFOptions)
	filteredList := cloudformation.FilterResources(list, "DummyResource")
	assert.Contains(t, filteredList, "cloudformation-waitcondition-")

}
