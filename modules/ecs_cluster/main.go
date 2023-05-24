package ecscluster

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/aws/ecs"
	"cdk.tf/go/stack/modules/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigStack struct {
	ServiceName string
}

type output struct {
	ClusterName string
}

func NewCluster(scope constructs.Construct, id string, config ConfigStack) output {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("provider"), provider.Account)
	cluster := ecs.NewEcsCluster(stack, jsii.String("ecs_cluster"), &ecs.EcsClusterConfig{
		Name: jsii.String(id),
	})

	// cdktf.NewTerraformOutput(stack, jsii.String("cluster_name"), &cdktf.TerraformOutputConfig{
	// 	Value: cluster.Name(),
	// })

	return output{*cluster.Name()}
}
