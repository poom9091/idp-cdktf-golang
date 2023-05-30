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
	// r := gin.Default()
	// r.GET("/vpc", func(c *gin.Context) {
	// 	out, err := exec.Command("cdktf", "apply", "--auto-approve", "vpc").Output()
	// 	if err != nil {
	// 		panic("Created vpc error: " + err.Error())
	// 	}
	// 	fmt.Println(out)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "vpc created",
	// 	})
	// })
	// r.Run()

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
