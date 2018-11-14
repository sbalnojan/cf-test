package cloudformation

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

//belongs into utils...

func FilterResources(list []*cloudformation.StackResource, CmpId string) string {
	res := ""
	for _, s := range list {
		if *s.LogicalResourceId == CmpId {
			res = *s.PhysicalResourceId
		}
	}
	return res
}

func ListResources(t *testing.T, CFOptions *Options) []*cloudformation.StackResource {
	resOutput, err := ListResourcesE(t, CFOptions)
	if err != nil {
		fmt.Println("Got error getting resources:")
		fmt.Println(err.Error())
		t.Fatal(err)
	}
	return resOutput
}
func ListResourcesE(t *testing.T, CFOptions *Options) ([]*cloudformation.StackResource, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(CFOptions.AWSRegion),
	}))
	svc := cloudformation.New(sess)

	resInput := &cloudformation.DescribeStackResourcesInput{
		StackName: aws.String(CFOptions.StackName)}
	resOutput, err := svc.DescribeStackResources(resInput)
	if err != nil {
		return nil, err
	}

	return resOutput.StackResources, nil
}
