package firewall

import (
	upEC2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RouteInfo struct {
	RouteTable, Destination, Region, InternetGateway string
}

func GetRoute(routeInfo RouteInfo) *upEC2.Route {
	var route upEC2.Route
	route = upEC2.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name: GetRouteName(routeInfo.RouteTable, routeInfo.Destination),
		},
		Spec: upEC2.RouteSpec{
			ForProvider: upEC2.RouteParameters_2{
				Region:               &routeInfo.Region,
				RouteTableID:         &routeInfo.RouteTable,
				DestinationCidrBlock: &routeInfo.Destination,
				GatewayID:            &routeInfo.InternetGateway,
			},
		},
	}
	return &route
}
