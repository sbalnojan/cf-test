package test

import (
	"fmt"
	"testing"

	"github.com/sbalnojan/cf-test/modules/cloudformation"
	"github.com/sbalnojan/terratest/modules/random"
)

func TestOverwrite(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())
	expectedName2 := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

	CFOptions := &cloudformation.Options{
		CFFile:    "../examples/cloudformation-aws-example/cf_outputs.yml",
		StackName: expectedName,
		AWSRegion: "us-west-1",
	}
	defer cloudformation.DeleteStack(t, CFOptions)

	cloudformation.CreateStack(t, CFOptions)

	CFOptions2 := &cloudformation.Options{
		CFFile:    "../examples/cloudformation-aws-example/cf_dependencies.yml",
		StackName: expectedName2,
		AWSRegion: "us-west-1",
	}
	defer cloudformation.DeleteStack(t, CFOptions2)

	cloudformation.CreateStack(t, CFOptions2)

}
