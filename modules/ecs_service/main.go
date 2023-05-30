package ecsservice

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/aws/ecr"
	"cdk.tf/go/stack/modules/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigStack struct {
	ServiceName     string
	ClusterArn      string
	Cpu             string
	Memory          string
	LogRetention    float64
	ApplicationPort float64
}

func NewService(scope constructs.Construct, id string, config ConfigStack) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("provider"), provider.Account)

	ecr := ecr.NewEcrRepository(stack, jsii.String("ecr"), &ecr.EcrRepositoryConfig{
		Name: jsii.String(id),
	})

	NewAutoScaling(stack, ConfigAutoScaling{
		ClusterArn:  config.ClusterArn,
		ServiceName: config.ServiceName,
	})

	NewCloudWatch(stack, ConfigCloudWatch{
		EcsServiceName: config.ServiceName,
		LogRetention:   config.LogRetention,
	})

	NewTaskDefinatoin(stack, ConfigTaskDefinatoin{
		EcsServiceName:  config.ServiceName,
		EcrImage:        *ecr.RepositoryUrl(),
		Cpu:             config.Cpu,
		Memory:          config.Memory,
		LogGroup:        "",
		Region:          "",
		Environment:     "",
		ApplicationPort: config.ApplicationPort,
		Secrets:         nil,
	})

	// NewEcsService(stack, ConfigEcsService{
	// 	EcsServiceName:
	// 	EcsClusterId:
	// 	TaskDefinitionArn     string
	// 	DesiredCount          float64
	// 	MinimumHealthyPercent float64
	// 	MaximumPercent        float64
	// 	Subnets               *string
	// 	SecurityGroupId       *string
	// 	AlbTargetArn          *string
	// 	ContainerName         *string
	// 	ContainerPort         float64
	// })

	return stack
}
