# cf-test

cf-test is a "module" for the Go library terratest (https://github.com/gruntwork-io/terratest). It enables
you to use the variety of helper functions and tools that terratest offers, but enables you to provide
your infrastructure using Cloudformation (not just Terraform).

## Why base this on terratest?

The reason this "module" is not a standalone library but integrates into terratest is that terratest:

- has an amazing variety of features
- is mostly focused on AWS (although that is changing)
- is actively developed
- is used in practice to maintain a huge repository of IaC
- also enables testing of Packer and Docker, which can integrate with Cloudformation

## Don't confuse provider with test module

There are two folders called "cloudformation", one for the provider written here,
and one for the AWS service which you can test, written at Terratest.

```
--- modules
----- cloudformation <- new! don't confuse with modules/aws/cloudformation in Terratest.
```

# Introduction to the module

Read the README.md of terratest, especially take note that the goal of terratest is
to deploy real infrastructure, with real costs.

The basic pattern for writing an automated test with cf-test is:

1.  Use cf-test to create & destroy your stack with a randomized stack name.
2.  Use cf-test helper functions to retrieve physical ids, outputs, exports etc. to test against.
3.  Use terratest AWS modules and helper functions to validate the behavior of your infrastructure.

## Example list

There are already some examples contained in this module which are also run as
part of testing the module itself, such as:

1.  [Helper Functions](/test/cf_resources_test.go): A usage example of retrieving physical ids (something that is
    not possible with Terraform)
2.  [Tagging an EC2 Instance](/test/cf_aws_example_test.go): A simple tag an instance with CF then retrieve
    and assert it's tag test.

## Usage Example 1

- Use the new module "cloudformation" to create & defer destroy your stack;
- Possibly retrieve a physical id from the stack
- Use core terratest functions to retrieve public IPs etc.
- Use core terratest functionality to test whatever you want to test.

Example adapted from: https://github.com/gruntwork-io/terratest/blob/master/test/terraform_aws_example_test.go

is at test > cr_create_test.go:

- cf_ec2_instances.yml is a CF containing a EC2 t2.micro instance with one tag.
  Let's check whether the tagging actually works.

- module "cloudformation"

```go
expectedName := fmt.Sprintf("terratest-cf-example-%s", random.UniqueId())

CFOptions := &cloudformation.Options{
  CFFile:    "cf_ec2_instance.yml",
  StackName: expectedName,
  AWSRegion: "us-east-1",
}
defer cloudformation.DeleteStack(t, CFOptions)

// create the ec2 instance
cloudformation.CreateStack(t, CFOptions)
```

- retrieve the physical ID, use it to check for tags with core terratest:

```go
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
```
