package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	aws_cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func ListOutputs(t *testing.T, CFOptions *Options) error {
	return nil
}

func ListExports(t *testing.T, CFOptions *Options) []*aws_cf.Export {
	svc := NewCFClient(t, CFOptions.AWSRegion)

	resInput := &aws_cf.ListExportsInput{
		NextToken: aws.String("")}
	resOutput, err := svc.ListExports(resInput)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return resOutput.Exports
}
