package loadbalance

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigStack struct {
	ServiceName string
	Cluster     string
}

type output struct {
}

func ALB(scope constructs.Construct, id string, config ConfigStack) (cdktf.TerraformStack, output) {
	stack := cdktf.NewTerraformStack(scope, &id)

	return stack
}
