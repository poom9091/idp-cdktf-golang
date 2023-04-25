package main

import (
	"fmt"

	base "cdk.tf/go/stack/base"

	"github.com/aws/constructs-go/constructs"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func Provider(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region:  jsii.String("ap-southeast-1"),
		Profile: jsii.String("myTerraform"),
	})
	// The code that defines your stack goes here
	return stack
}

func PrintProjectName(txt string) {
	fmt.Printf("%s \n", txt)
}
func main() {
	app := cdktf.NewApp(nil)

	base.PrintProjectName("Provide dev")
	aws.AwsProvider

	app.Synth()
}
