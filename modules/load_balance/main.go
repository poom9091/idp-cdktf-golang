package loadbalance

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/aws/elb"
	"cdk.tf/go/stack/generated/aws/vpc"
	"cdk.tf/go/stack/modules/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigStack struct {
	ProjectName   string
	VpcId         *string
	PublicSubnets *string
	Cidr          *string
}

type output struct {
	AlbArn string
}

func NewALB(scope constructs.Construct, id string, config ConfigStack) (cdktf.TerraformStack, output) {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("provider"), provider.Account)

	portsList := cdktf.NewTerraformVariable(stack, jsii.String("ports"), &cdktf.TerraformVariableConfig{
		Type:    cdktf.VariableType_LIST_STRING(),
		Default: []string{"80", "443"},
	})

	portsListIter := cdktf.TerraformIterator_FromList(portsList.ListValue())

	sg := vpc.NewSecurityGroup(stack, jsii.String("security_group"), &vpc.SecurityGroupConfig{
		Name:  jsii.String("applicaton"),
		VpcId: jsii.String(*config.VpcId),
	})

	vpc.NewSecurityGroupRule(stack, jsii.String("security_group_rule_egress"), &vpc.SecurityGroupRuleConfig{
		Type:            jsii.String("egress"),
		FromPort:        cdktf.Token_AsNumber(0),
		ToPort:          cdktf.Token_AsNumber(0),
		CidrBlocks:      jsii.Strings(*config.Cidr),
		Protocol:        jsii.String("tcp"),
		SecurityGroupId: jsii.String(*sg.Id()),
	})

	vpc.NewSecurityGroupRule(stack, jsii.String("security_group_ingress"), &vpc.SecurityGroupRuleConfig{
		ForEach:         portsListIter,
		Type:            jsii.String("ingress"),
		FromPort:        cdktf.Token_AsNumber(portsListIter.Value()),
		ToPort:          cdktf.Token_AsNumber(portsListIter.Value()),
		CidrBlocks:      jsii.Strings(*config.Cidr),
		Protocol:        jsii.String("tcp"),
		SecurityGroupId: jsii.String(*sg.Id()),
	})

	alb := elb.NewAlb(stack, jsii.String("application_load_balancer"), &elb.AlbConfig{
		Name:                     jsii.String(config.ProjectName + "-alb"),
		Internal:                 jsii.Bool(false),
		Subnets:                  cdktf.Fn_Tolist(*config.PublicSubnets),
		LoadBalancerType:         jsii.String("application"),
		SecurityGroups:           jsii.Strings(*sg.Id()),
		EnableDeletionProtection: jsii.Bool(true),
	})

	return stack, output{
		AlbArn: *alb.Arn(),
	}
}
