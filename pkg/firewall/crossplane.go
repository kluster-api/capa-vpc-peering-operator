package firewall

import (
	"context"
	upEC2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmc "kmodules.xyz/client-go/client"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	DestinationCidr, Region, SecurityGroup, ToPort, FromPort string
}

func GetRule(ruleInfo RuleInfo) (*upEC2.SecurityGroupRule, error) {
	var rule upEC2.SecurityGroupRule
	toPort, err := strconv.ParseFloat(ruleInfo.ToPort, 64)
	if err != nil {
		return nil, err
	}

	fromPort, err := strconv.ParseFloat(ruleInfo.FromPort, 64)

	protocol := "tcp"
	typ := "ingress"

	rule = upEC2.SecurityGroupRule{
		ObjectMeta: metav1.ObjectMeta{
			Name: GetSGRuleName(ruleInfo.SecurityGroup, ruleInfo.DestinationCidr),
		},
		Spec: upEC2.SecurityGroupRuleSpec{
			ForProvider: upEC2.SecurityGroupRuleParameters_2{
				Region:          &ruleInfo.Region,
				CidrBlocks:      []*string{&ruleInfo.DestinationCidr},
				ToPort:          &toPort,
				FromPort:        &fromPort,
				Protocol:        &protocol,
				SecurityGroupID: &ruleInfo.SecurityGroup,
				Type:            &typ,
			},
		},
	}
	return &rule, nil
}

func GetVpcIDs(ctx context.Context, c client.Client) []string {
	pcLIst := &upEC2.VPCPeeringConnectionList{}
	c.List(ctx, pcLIst)

	var idList []string

	for _, pc := range pcLIst.Items {
		idList = append(idList, pc.GetID())
	}
	return idList
}

func CreateSecurityGroupRule(ctx context.Context, c client.Client, info RuleInfo) error {
	sgRule, err := GetRule(info)
	if err != nil {
		return err
	}

	_, _, err = kmc.CreateOrPatch(ctx, c, sgRule, func(_ client.Object, _ bool) client.Object {
		return sgRule
	})
	if err != nil {
		return err
	}
	return nil
}
