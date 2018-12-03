package cloudformation

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// CreateStack creates the supplied stack
// throws error
func CreateStack(t *testing.T, CFOptions *Options) {
	err := CreateStackE(t, CFOptions)
	if err != nil {
		fmt.Println("== ", t.Name(), " == Got error waiting for stack to be created")
		fmt.Println(err)
		t.Fatal(err)
	}

	fmt.Println("== ", t.Name(), " == Created stack "+CFOptions.StackName)

}

// CreateStackE creates the supplied stack
// returns possible error
func CreateStackE(t *testing.T, CFOptions *Options) error {

	cwd, _ := os.Getwd()
	templateBodyBin, err := ioutil.ReadFile(path.Join(cwd, CFOptions.CFFile))
	if err != nil {
		return err
	}
	templateBody := string(templateBodyBin)

	if err := CreateStackStrE(t, CFOptions, templateBody); err != nil {
		return err
	}
	return nil
}

// CreateStackStr creates stack from supplied templateBody
// throws error
func CreateStackStr(t *testing.T, CFOptions *Options, templateBody string) {
	err := CreateStackStrE(t, CFOptions, templateBody)
	if err != nil {
		t.Fatal(err)
	}
}

// CreateStackStrE creates stack from supplied templateBody
// returns error
func CreateStackStrE(t *testing.T, CFOptions *Options, templateBody string) error {

	svc := NewCFClient(t, CFOptions.AWSRegion)

	input := &cloudformation.CreateStackInput{
		TemplateBody: aws.String(templateBody),
		StackName:    aws.String(CFOptions.StackName),
		Parameters:   CFOptions.Parameters,
		Capabilities: []*string{aws.String("CAPABILITY_IAM")}}

	fmt.Println("== ", t.Name(), " == Creating stack: ", CFOptions.StackName)

	_, err := svc.CreateStack(input)
	if err != nil {
		fmt.Println("== ", t.Name(), " == Got error creating stack:")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("== ", t.Name(), " == Waiting for stack to be created")

	// Wait until stack is created
	desInput := &cloudformation.DescribeStacksInput{StackName: aws.String(CFOptions.StackName)}
	err = svc.WaitUntilStackCreateComplete(desInput)
	if err != nil {
		return err
	}
	return nil
}
