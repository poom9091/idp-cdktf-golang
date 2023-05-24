package main

import (
	ecscluster "cdk.tf/go/stack/modules/ecs_cluster"
	ecsservice "cdk.tf/go/stack/modules/ecs_service"
	loadbalance "cdk.tf/go/stack/modules/load_balance"
	vpc "cdk.tf/go/stack/modules/vpc"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	vpcOutput := vpc.NewVPC(app, "vpc")

	ecsClusterOutput := ecscluster.NewCluster(app, "ecs_cluster", ecscluster.ConfigStack{
		ServiceName: "demo-idp",
	})

	ecsservice.NewService(app, "ecs_service", ecsservice.ConfigStack{
		ServiceName: "demo-idp",
		Cluster:     ecsClusterOutput.ClusterName,
	})

	loadbalance.NewALB(app, "alb", loadbalance.ConfigStack{
		ProjectName:   "demo-idp",
		VpcId:         vpcOutput.VpcId,
		PublicSubnets: vpcOutput.PublicSubnets,
		Cidr:          vpcOutput.Cidr,
	})

	app.Synth()
}
