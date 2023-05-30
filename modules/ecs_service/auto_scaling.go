package ecsservice

import (
	"fmt"

	"cdk.tf/go/stack/generated/aws/appautoscaling"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ConfigAutoScaling struct {
	ClusterArn  string
	ServiceName string
}

func NewAutoScaling(scope constructs.Construct, config ConfigAutoScaling) {
	scalingTarget := appautoscaling.NewAppautoscalingTarget(scope, jsii.String("scaling_target"), &appautoscaling.AppautoscalingTargetConfig{
		MaxCapacity:       jsii.Number(4),
		MinCapacity:       jsii.Number(1),
		ResourceId:        jsii.String(fmt.Sprintf("service/%v/%v", config.ClusterArn, config.ServiceName)),
		ScalableDimension: jsii.String("ecs:service:DesiredCount"),
		ServiceNamespace:  jsii.String("ecs"),
	})

	appautoscaling.NewAppautoscalingPolicy(scope, jsii.String("scaling_memory_policy"), &appautoscaling.AppautoscalingPolicyConfig{
		Name:              jsii.String("memory-autoscaling"),
		PolicyType:        jsii.String("TargetTrackingScaling"),
		ResourceId:        jsii.String(*scalingTarget.ResourceId()),
		ScalableDimension: jsii.String(*scalingTarget.ScalableDimension()),
		ServiceNamespace:  jsii.String(*scalingTarget.ServiceNamespace()),
		TargetTrackingScalingPolicyConfiguration: &appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration{
			PredefinedMetricSpecification: &appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification{
				PredefinedMetricType: jsii.String("ECSServiceAverageMemoryUtilization"),
			},
			TargetValue: jsii.Number(80),
		},
	})

	appautoscaling.NewAppautoscalingPolicy(scope, jsii.String("scaling_cpu_policy"), &appautoscaling.AppautoscalingPolicyConfig{
		Name:              jsii.String("cpu-autoscaling"),
		PolicyType:        jsii.String("TargetTrackingScaling"),
		ResourceId:        jsii.String(*scalingTarget.ResourceId()),
		ScalableDimension: jsii.String(*scalingTarget.ScalableDimension()),
		ServiceNamespace:  jsii.String(*scalingTarget.ServiceNamespace()),
		TargetTrackingScalingPolicyConfiguration: &appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration{
			PredefinedMetricSpecification: &appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification{
				PredefinedMetricType: jsii.String("ECSServiceAverageCPUUtilization"),
			},
			TargetValue: jsii.Number(60),
		},
	})
}
