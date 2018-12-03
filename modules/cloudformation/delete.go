package cloudformation

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// DeleteStack delets the supplied stack
func DeleteStack(t *testing.T, CFOptions *Options) {
	err := DeleteStackE(t, CFOptions)
	if err != nil {
		fmt.Println("Got error waiting for stack to be deleted")
		fmt.Println(err)
		t.Fatal(err)
	}

	fmt.Println("Deleted stack:", CFOptions.StackName)
}

// DeleteStackE delets the supplied stack
// returns possible error
func DeleteStackE(t *testing.T, CFOptions *Options) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(CFOptions.AWSRegion),
	}))

	svc := cloudformation.New(sess)

	delInput := &cloudformation.DeleteStackInput{
		StackName: aws.String(CFOptions.StackName)}
	_, err := svc.DeleteStack(delInput)
	if err != nil {
		fmt.Println("Got error deleting stack:", CFOptions.StackName)
		fmt.Println(err.Error())
		return err
	}

	// Wait until stack is deleted
	desInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(CFOptions.StackName)}
	err = svc.WaitUntilStackDeleteComplete(desInput)

	if err != nil {
		return err
	}
	return nil
}
