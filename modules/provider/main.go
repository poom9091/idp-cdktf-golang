package provider

import (
	"cdk.tf/go/stack/generated/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

var Account = &aws.AwsProviderConfig{
	Region:  jsii.String("ap-southeast-1"),
	Profile: jsii.String("myTerraform"),
}

func SetProvider(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("provider"), Account)
	return stack
}
