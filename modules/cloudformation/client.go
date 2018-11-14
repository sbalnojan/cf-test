package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

func NewCFClient(t *testing.T, Region string) *aws_cf.CloudFormation {
	client, err := NewCFClientE(Region)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func NewCFClientE(Region string) (*aws_cf.CloudFormation, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
	})
	if err != nil {
		return nil, err
	}

	return aws_cf.New(sess), nil
}
