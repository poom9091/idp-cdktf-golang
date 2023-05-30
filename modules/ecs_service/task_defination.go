package ecsservice

import (
	"fmt"

	"cdk.tf/go/stack/generated/aws/ecs"
	"cdk.tf/go/stack/generated/aws/iam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ConfigTaskDefinatoin struct {
	EcsServiceName  string
	EcrImage        string
	Cpu             string
	Memory          string
	LogGroup        string
	Region          string
	Environment     string
	ApplicationPort float64
	Secrets         []string
}

func NewTaskDefinatoin(scope constructs.Construct, config ConfigTaskDefinatoin) {
	iamRole := iam.NewIamRole(scope, jsii.String("iamRole"), &iam.IamRoleConfig{
		Name: jsii.String(fmt.Sprintf("%s-task-role", config.EcsServiceName)),
		AssumeRolePolicy: jsii.String(`
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
		`),
	})

	applicatoinPolicy := iam.NewIamPolicy(scope, jsii.String("iamPolicy"), &iam.IamPolicyConfig{
		Name: jsii.String(fmt.Sprintf("%s-task-policy", config.EcsServiceName)),
		Policy: jsii.String(`
		{
			"Version": "2012-10-17",
			"Statement": [
			  {
				"Action": [
				  "ssm:*"
				],
				"Effect": "Allow",
				"Resource": "*"
			  }
			]
		 }`),
	})

	iam.NewIamRolePolicyAttachment(scope, jsii.String("iamRoleEcsPolicyAttachment"), &iam.IamRolePolicyAttachmentConfig{
		Role:      iamRole.Name(),
		PolicyArn: jsii.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
	})

	iam.NewIamRolePolicyAttachment(scope, jsii.String("iamRoleEcsAdditionalPolicyAttachment"), &iam.IamRolePolicyAttachmentConfig{
		Role:      iamRole.Name(),
		PolicyArn: applicatoinPolicy.Arn(),
	})

	ecs.NewEcsTaskDefinition(scope, jsii.String("task-definition"), &ecs.EcsTaskDefinitionConfig{
		Family:           jsii.String(fmt.Sprintf("%s-task-defination", config.EcsServiceName)),
		NetworkMode:      jsii.String("awsvpc"),
		TaskRoleArn:      iamRole.Arn(),
		Cpu:              jsii.String(config.Cpu),
		Memory:           jsii.String(config.Memory),
		ExecutionRoleArn: iamRole.Arn(),
		ContainerDefinitions: jsii.String(fmt.Sprintf(`
		[
			{
			  name      = %v 
			  image     = %v 
			  essential = true
			  logConfiguration = {
				logDriver     = "awslogs"
				secretOptions = null
				options = {
				  awslogs-group         = %v
				  awslogs-region        = %v 
				  awslogs-stream-prefix = "ecs"
				}
			  }
	
			  secrets     = %v 
			  environment = %v 
			  portMappings = [
				{
				  containerPort = %v
				}
			  ]
			},
	
		],`,
			config.EcsServiceName,
			config.EcrImage,
			config.LogGroup,
			config.Region,
			config.Secrets,
			config.Environment,
			config.ApplicationPort,
		)),
	})
}
