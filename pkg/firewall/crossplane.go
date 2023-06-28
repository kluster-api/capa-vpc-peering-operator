/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package firewall

import (
	"context"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	kmc "kmodules.xyz/client-go/client"
	ecapi "kubeform.dev/provider-aws/apis/ec2/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	FirstPort      = "0"
	LastPort       = "65535"
	CidrAnnotation = "aws.upbound.io/peer-vpc-cidr"
)

type RouteInfo struct {
	RouteTable, Destination, Region, PeeringConnectionID string
}

func getRoute(routeInfo RouteInfo, ownerRef []metav1.OwnerReference) *ecapi.Route {
	route := ecapi.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name:            GetRouteName(routeInfo.RouteTable, routeInfo.Destination),
			OwnerReferences: ownerRef,
		},
		Spec: ecapi.RouteSpec{
			ForProvider: ecapi.RouteParameters{
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

func getRule(ruleInfo RuleInfo, ownerRef []metav1.OwnerReference) (*ecapi.SecurityGroupRule, error) {
	var rule ecapi.SecurityGroupRule
	toPort, err := strconv.ParseFloat(ruleInfo.ToPort, 64)
	if err != nil {
		return nil, err
	}

	fromPort, err := strconv.ParseFloat(ruleInfo.FromPort, 64)
	if err != nil {
		return nil, err
	}

	protocol := "tcp"
	typ := "ingress"

	rule = ecapi.SecurityGroupRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:            GetSGRuleName(ruleInfo.SecurityGroup, ruleInfo.DestinationCidr),
			OwnerReferences: ownerRef,
		},
		Spec: ecapi.SecurityGroupRuleSpec{
			ForProvider: ecapi.SecurityGroupRuleParameters{
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
	sgRule, err := getRule(info, ownerRef)
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
	route := getRoute(info, ownerRef)
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
