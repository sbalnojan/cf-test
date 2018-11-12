# CF module for terratest

This can only be used in conjunction with the amazing
(terraform/packer/...) integration testing library terratest (https://github.com/gruntwork-io/terratest).

## Why not write your own lib?

Because I find terratest to be the best AWS infrastructure testing lib out there.

Besides, the code base here is super small compared to terratest.

## Diff

```
--- modules
----- cloudformation <- new! don't confuse with modules/aws/cloudformation in terratest.
```

## Usage Example

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

## To Do

- write an outputs module
- write an ImportVal etc. overwrite helper function.
- write an ECS Cluster test. (=> first write a terratest AWS ECS module?)
- Set Parameters in https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/#CreateStackInput
  to use testing parameters (names, sizes etc.)
