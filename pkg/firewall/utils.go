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
	"errors"
	"fmt"
	"strings"

	crossplanev1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	upEC2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	infrav2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetRouteName(routeTable, destination string) string {
	st := fmt.Sprintf("%s-%s", routeTable, destination)
	st = strings.ReplaceAll(st, ".", "-")
	st = strings.ReplaceAll(st, "/", "-")
	return st
}

func GetSGRuleName(securityGroup, cidr string) string {
	st := fmt.Sprintf("%s-%s", securityGroup, cidr)
	st = strings.ReplaceAll(st, ".", "-")
	st = strings.ReplaceAll(st, "/", "-")
	return st
}

func GetManagedControlPlane(ctx context.Context, c client.Client) (*ekscontrolplanev1.AWSManagedControlPlane, error) {
	managedCPList := &ekscontrolplanev1.AWSManagedControlPlaneList{}
	err := c.List(ctx, managedCPList)
	if err != nil {
		return nil, err
	}

	for _, mCP := range managedCPList.Items {
		return &mCP, nil
	}
	return nil, errors.New("failed to get any managed controlplane resource")
}

func GetSecurityGroupID(managedCP *ekscontrolplanev1.AWSManagedControlPlane) (string, error) {
	sg, found := managedCP.Status.Network.SecurityGroups[infrav2.SecurityGroupNode]
	if !found {
		return "", errors.New("no security group id found")
	}

	return sg.ID, nil
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

func CheckCrossplaneCondition(conditions []crossplanev1.Condition) bool {
	for i := range conditions {
		klog.Infof("condition status: %s", conditions[i].Status)
		if conditions[i].Status != "True" {
			return false
		}
	}
	return true
}

func GetOwnerReference(pc *upEC2.VPCPeeringConnection) []metav1.OwnerReference {
	return []metav1.OwnerReference{
		{
			APIVersion: pc.APIVersion,
			Kind:       pc.Kind,
			Name:       pc.Name,
			UID:        pc.UID,
		},
	}
}
