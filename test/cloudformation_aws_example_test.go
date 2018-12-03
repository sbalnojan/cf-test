package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/sbalnojan/cf-test/modules/cloudformation"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-example using Terratest.
func TestCFCreateStack(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

	CFOptions := &cloudformation.Options{
		CFFile:    "../examples/cloudformation-aws-example/cf_create_test.yml",
		StackName: expectedName,
		AWSRegion: "us-west-2",
	}
	defer cloudformation.DeleteStack(t, CFOptions)

	cloudformation.CreateStack(t, CFOptions)
	list := cloudformation.ListResources(t, CFOptions)
	filteredList := cloudformation.FilterResources(list, "DummyResource")
	assert.Contains(t, filteredList, "cloudformation-waitcondition-")
}

func TestCFTagInstance(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

	CFOptions := &cloudformation.Options{
		CFFile:    "../examples/cloudformation-aws-example/cf_ec2_instance.yml",
		StackName: expectedName,
		AWSRegion: "us-east-1",
	}
	defer cloudformation.DeleteStack(t, CFOptions)

	// create the ec2 instance
	cloudformation.CreateStack(t, CFOptions)

	// retrieve it's physical id using it's logical id
	list := cloudformation.ListResources(t, CFOptions)
	pID := cloudformation.FilterResources(list, "EC2Instance")
	fmt.Println("physical id: ", pID)

	// get IP using core terratest
	IP := aws.GetPublicIpOfEc2Instance(t, pID, CFOptions.AWSRegion)

	// check that it's there...
	fmt.Println("public Ip: ", IP)

	// tag the resource
	instanceTags := aws.GetTagsForEc2Instance(t, CFOptions.AWSRegion, pID)

	testingTag, containsTestingTag := instanceTags["Name"]
	assert.True(t, containsTestingTag)
	assert.Equal(t, "test-instance", testingTag)
}

func TestCFOutputs(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

	CFOptions := &cloudformation.Options{
		CFFile:    "../examples/cloudformation-aws-example/cf_outputs.yml",
		StackName: expectedName,
		AWSRegion: "us-east-1",
	}
	defer cloudformation.DeleteStack(t, CFOptions)

	// create the stack
	cloudformation.CreateStack(t, CFOptions)

	// outputs := cloudformation.ListExports(t, CFOptions)

	// setting our expected Output

	// assert.Equal(t, nil, outputs)

}
