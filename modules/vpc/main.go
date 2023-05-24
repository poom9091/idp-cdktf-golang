package vpc

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/vpc"
	"cdk.tf/go/stack/modules/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type output struct {
	VpcId         *string
	PublicSubnets *string
	Cidr          *string
}

func NewVPC(scope constructs.Construct, id string) *output {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("provider"), provider.Account)
	vpc := vpc.NewVpc(stack, jsii.String("AWS"), &vpc.VpcOptions{
		Name:               jsii.String("dev-vpc"),
		Cidr:               jsii.String("10.10.0.0/16"),
		Azs:                jsii.Strings("ap-southeast-1a", "ap-southeast-1b"),
		PrivateSubnets:     jsii.Strings("10.10.1.0/24", "10.10.2.0/24"),
		PublicSubnets:      jsii.Strings("10.10.10.0/24", "10.10.20.0/24"),
		EnableDnsHostnames: jsii.Bool(true),
		EnableNatGateway:   jsii.Bool(false),
		EnableVpnGateway:   jsii.Bool(true),
	})

	return &output{
		VpcId:         vpc.VpcIdOutput(),
		PublicSubnets: vpc.PublicSubnetsOutput(),
		Cidr:          vpc.VpcCidrBlockOutput(),
	}
}
