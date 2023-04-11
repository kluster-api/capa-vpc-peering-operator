/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package firewall

import (
	"context"
	upEC2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	kmc "kmodules.xyz/client-go/client"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
)

const (
	FirstPort      = "0"
	LastPort       = "65535"
	CidrAnnotation = "aws.upbound.io/peer-vpc-cidr"
)

type RouteInfo struct {
	RouteTable, Destination, Region, PeeringConnectionID string
}

func GetRoute(routeInfo RouteInfo, ownerRef []metav1.OwnerReference) *upEC2.Route {
	var route upEC2.Route
	route = upEC2.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name:            GetRouteName(routeInfo.RouteTable, routeInfo.Destination),
			OwnerReferences: ownerRef,
		},
		Spec: upEC2.RouteSpec{
			ForProvider: upEC2.RouteParameters_2{
				Region:                 &routeInfo.Region,
				RouteTableID:           &routeInfo.RouteTable,
				DestinationCidrBlock:   &routeInfo.Destination,
				VPCPeeringConnectionID: &routeInfo.PeeringConnectionID,
			},
		},
	}
	return &route
}

type RuleInfo struct {
	DestinationCidr, Region, SecurityGroup, ToPort, FromPort string
}

func GetRule(ruleInfo RuleInfo, ownerRef []metav1.OwnerReference) (*upEC2.SecurityGroupRule, error) {
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
			Name:            GetSGRuleName(ruleInfo.SecurityGroup, ruleInfo.DestinationCidr),
			OwnerReferences: ownerRef,
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

func CreateSecurityGroupRule(ctx context.Context, c client.Client, info RuleInfo, ownerRef []metav1.OwnerReference) error {
	sgRule, err := GetRule(info, ownerRef)
	if err != nil {
		return err
	}

	_, _, err = kmc.CreateOrPatch(ctx, c, sgRule, func(_ client.Object, _ bool) client.Object {
		return sgRule
	})
	if err != nil {
		return err
	}
	klog.Infof("rule created to %s for %s", info.SecurityGroup, info.DestinationCidr)
	return nil
}

func CreateRouteTableRoute(ctx context.Context, c client.Client, info RouteInfo, ownerRef []metav1.OwnerReference) error {
	route := GetRoute(info, ownerRef)
	_, _, err := kmc.CreateOrPatch(ctx, c, route, func(_ client.Object, _ bool) client.Object {
		return route
	})
	if err != nil {
		return err
	}

	klog.Infof("route created to table %s for %s", info.RouteTable, info.Destination)

	return nil
}

type VPCIdentifier struct {
	Name, Cidr string
}

func CheckCIDRConflict(ctx context.Context, c client.Client, ownVPC VPCIdentifier) error {
	vpcs := []VPCIdentifier{ownVPC}
	pcs := &upEC2.VPCPeeringConnectionList{}

	err := c.List(ctx, pcs)
	if err != nil {
		return err
	}

	for _, pc := range pcs.Items {
		if len(pc.Annotations[CidrAnnotation]) > 0 {
			vpcs = append(vpcs, VPCIdentifier{Name: *pc.Spec.ForProvider.PeerVPCID, Cidr: pc.Annotations[CidrAnnotation]})
		}
	}

	return nil
}
