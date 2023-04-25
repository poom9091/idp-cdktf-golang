package base

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	aws "github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.AwsProvider
	// The code that defines your stack goes here
	return stack
}

func PrintProjectName(txt string) {
	fmt.Printf("%s \n", txt)
}
