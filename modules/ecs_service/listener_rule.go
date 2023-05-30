package ecsservice

import (
	"cdk.tf/go/stack/generated/aws/elb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ConfigTargetGroup struct {
	Listener_arn        string
	Service_domain_name string
	Rule_Priority       float64
	Target_Group        string
	Port                float64
	VpcId               string
}

type OutputTargetGroup struct {
	targetGroup string
}

func NewTargetGroup(scope constructs.Construct, config ConfigTargetGroup) OutputTargetGroup {
	targetGroup := elb.NewAlbTargetGroup(scope, jsii.String("target_group"), &elb.AlbTargetGroupConfig{
		Name:       jsii.String(""),
		Port:       jsii.Number(config.Port),
		Protocol:   jsii.String("HTTP"),
		VpcId:      jsii.String(config.VpcId),
		TargetType: jsii.String("ip"),
	})

	elb.NewAlbListenerRule(scope, jsii.String("listener_rule"), &elb.AlbListenerRuleConfig{
		ListenerArn: &config.Listener_arn,
		Priority:    jsii.Number(config.Rule_Priority),
		Action: &elb.AlbListenerRuleActionForward{
			TargetGroup: &elb.AlbListenerRuleActionForwardTargetGroup{
				Arn: jsii.String(*targetGroup.Arn()),
			},
		},
		Condition: &elb.AlbListenerRuleConditionHostHeader{
			Values: jsii.Strings(config.Service_domain_name),
		},
	})

	return OutputTargetGroup{
		targetGroup: *targetGroup.Arn(),
	}
}
