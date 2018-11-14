package cloudformation

import (
	aws_cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

type Options struct {
	CFFile     string              //Path to the Cloudformation template, yes singular currently!
	StackName  string              //The name of the stack to be.
	AWSRegion  string              //AWS Region
	Parameters []*aws_cf.Parameter //Template Parameters as aws cloudformation sdk params.
}