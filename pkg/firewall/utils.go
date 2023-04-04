package firewall

import (
	"context"
	"errors"
	"fmt"
	infrav2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
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
