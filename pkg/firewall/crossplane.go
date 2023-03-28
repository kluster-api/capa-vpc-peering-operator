package firewall

import (
	upEC2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
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

type RuleInfo struct {
	Cidr, Region, SecurityGroup, Port string
}

func GetRule(ruleInfo RuleInfo) (*upEC2.SecurityGroupRule, error) {
	var rule upEC2.SecurityGroupRule
	fPort, err := strconv.ParseFloat(ruleInfo.Port, 64)
	if err != nil {
		return nil, err
	}
	protocol := "tcp"
	typ := "ingress"

	rule = upEC2.SecurityGroupRule{
		ObjectMeta: metav1.ObjectMeta{
			Name: GetSGRuleName(ruleInfo.SecurityGroup, ruleInfo.Cidr, ruleInfo.Port),
		},
		Spec: upEC2.SecurityGroupRuleSpec{
			ForProvider: upEC2.SecurityGroupRuleParameters_2{
				Region:          &ruleInfo.Region,
				CidrBlocks:      []*string{&ruleInfo.Cidr},
				FromPort:        &fPort,
				ToPort:          &fPort,
				Protocol:        &protocol,
				SecurityGroupID: &ruleInfo.SecurityGroup,
				Type:            &typ,
			},
		},
	}
	return &rule, nil
}
