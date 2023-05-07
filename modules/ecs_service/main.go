package ecsservice

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/aws/ecr"
	"cdk.tf/go/stack/generated/aws/ecs"
	"cdk.tf/go/stack/generated/aws/iam"
	"cdk.tf/go/stack/modules/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigStack struct {
	ServiceName string
	Cluster     string
}

func ECS_service(scope constructs.Construct, id string, config ConfigStack) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("provider"), provider.Account)
	assumeRolePolicyJson := `
		{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Action": "sts:AssumeRole",
					"Principal": {
					  "Service": "ecs-tasks.amazonaws.com"
					},
					"Effect": "Allow",
					"Sid": ""
				}
			]
		}
	`
	iamRole := iam.NewIamRole(stack, jsii.String("iamRole"), &iam.IamRoleConfig{
		Name:             jsii.String(id + "iam_role"),
		AssumeRolePolicy: jsii.String(assumeRolePolicyJson),
	})

	ecr.NewEcrRepository(stack, jsii.String("ecr"), &ecr.EcrRepositoryConfig{
		Name: jsii.String(id),
	})

	containerDefinitions := `
	{
		name      = "${var.environment}-${var.project_name}-${var.service}-container"
		image     = var.container_tag != null ? "${var.ecr_url}:${var.container_tag}" : "${var.ecr_url}:latest"
		essential = true
		logConfiguration = {
		  logDriver     = "awslogs"
		  secretOptions = null
		  options = {
			awslogs-group         = "/ecs/${var.environment}-${var.project_name}-${var.service}"
			awslogs-region        = var.region
			awslogs-stream-prefix = "ecs"
		  }
		}

		dependsOn = var.dependsOn_sidercar_container == null ? [] : var.dependsOn_sidercar_container

		secrets     = var.env_secret
		environment = var.env
		portMappings = [
		  {
			containerPort = var.container_port
		  }
		]
	  },
	`

	ecs.NewEcsTaskDefinition(stack, jsii.String("ecsTaskDefinition"), &ecs.EcsTaskDefinitionConfig{
		NetworkMode:             jsii.String("awsvpc"),
		Family:                  jsii.String(id + "_task_definition"),
		TaskRoleArn:             jsii.String(*iamRole.Arn()),
		RequiresCompatibilities: jsii.Strings("FARGATE"),
		Cpu:                     jsii.String("256"),
		Memory:                  jsii.String("512"),
		ExecutionRoleArn:        jsii.String(*iamRole.Arn()),
		ContainerDefinitions:    jsii.String(containerDefinitions),
	})

	ecs.NewEcsService(stack, jsii.String("ecsService"), &ecs.EcsServiceConfig{
		Name:    jsii.String(id + "cluster"),
		Cluster: jsii.String(config.Cluster),
	})

	return stack
}
