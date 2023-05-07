package main

import (
	ecscluster "cdk.tf/go/stack/modules/ecs_cluster"
	ecsservice "cdk.tf/go/stack/modules/ecs_service"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	_, ecsClusterOutput := ecscluster.ECS_Cluster(app, "ecs_cluster", ecscluster.ConfigStack{
		ServiceName: "demo-idp",
	})

	ecsservice.ECS_service(app, "ecs_service", ecsservice.ConfigStack{
		ServiceName: "demo-idp",
		Cluster:     ecsClusterOutput.ClusterName,
	})

	app.Synth()
}
