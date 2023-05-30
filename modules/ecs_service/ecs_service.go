package ecsservice

import (
	"fmt"

	"cdk.tf/go/stack/generated/aws/ecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigEcsService struct {
	EcsServiceName        string
	EcsClusterId          string
	TaskDefinitionArn     string
	DesiredCount          float64
	MinimumHealthyPercent float64
	MaximumPercent        float64
	Subnets               *string
	SecurityGroupId       *string
	AlbTargetArn          *string
	ContainerName         *string
	ContainerPort         float64
}

func NewEcsService(scope constructs.Construct, config ConfigEcsService) {
	ecs.NewEcsService(scope, jsii.String("main"), &ecs.EcsServiceConfig{
		Name:                            jsii.String(fmt.Sprintf("%v-log-stream", config.EcsServiceName)),
		Cluster:                         jsii.String(config.EcsClusterId),
		TaskDefinition:                  jsii.String(config.TaskDefinitionArn),
		DesiredCount:                    jsii.Number(config.DesiredCount),
		DeploymentMinimumHealthyPercent: jsii.Number(config.MinimumHealthyPercent),
		DeploymentMaximumPercent:        jsii.Number(config.MaximumPercent),
		LaunchType:                      jsii.String("FARGATE"),
		NetworkConfiguration: &ecs.EcsServiceNetworkConfiguration{
			Subnets:        cdktf.Fn_Tolist(*config.Subnets),
			SecurityGroups: jsii.Strings(*config.SecurityGroupId),
			AssignPublicIp: jsii.Bool(true),
		},
		DeploymentController: &ecs.EcsServiceDeploymentController{
			Type: jsii.String("ECS"),
		},
		LoadBalancer: &ecs.EcsServiceLoadBalancer{
			TargetGroupArn: jsii.String(*config.AlbTargetArn),
			ContainerName:  jsii.String(*config.ContainerName),
			ContainerPort:  jsii.Number(config.ContainerPort),
		},
	})
}
