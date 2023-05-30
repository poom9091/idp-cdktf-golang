package ecsservice

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/cloudwatch"
)

type ConfigCloudWatch struct {
	EcsServiceName string
	LogRetention   float64
}

func NewCloudWatch(scope constructs.Construct, config ConfigCloudWatch) {
	logGroup := cloudwatch.NewCloudwatchLogGroup(scope, jsii.String("main"), &cloudwatch.CloudwatchLogGroupConfig{
		Name:            jsii.String(fmt.Sprintf("ecs/%v", config.EcsServiceName)),
		RetentionInDays: jsii.Number(config.LogRetention),
	})

	cloudwatch.NewCloudwatchLogStream(scope, jsii.String("main"), &cloudwatch.CloudwatchLogStreamConfig{
		Name:         jsii.String(fmt.Sprintf("%v-log-stream", config.EcsServiceName)),
		LogGroupName: logGroup.Name(),
	})
}
