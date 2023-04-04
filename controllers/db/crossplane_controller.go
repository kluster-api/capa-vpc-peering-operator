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

package db

import (
	"context"
	"errors"
	upEC2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	kmc "kmodules.xyz/client-go/client"
	"kubedb/aws-peering-connection-operator/pkg/firewall"
	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	firstPort = "0"
	lastPort  = "65535"
)

// Reconciler reconciles a Crossplane object
type Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// PCReconciler reconciles a crossplane vpc peering connection object
type PCReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r PCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	objKey := req.NamespacedName
	pc := &upEC2.VPCPeeringConnection{}

	if err := r.Get(ctx, objKey, pc); err != nil {
		return ctrl.Result{}, err
	}

	if pc.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	managedCP, err := firewall.GetManagedControlPlane(ctx, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}
	sgID, err := firewall.GetSecurityGroupID(managedCP)
	if err != nil {
		return ctrl.Result{}, err
	}

	cidr, found := pc.Spec.ForProvider.Tags["CIDR"]
	if !found {
		return ctrl.Result{}, errors.New("empty CIDR range in peering connection tags")
	}

	//TODO: we need the CIDR range of application VPC
	//frontend will apply a peering connection resource from that we can get
	//the VPC ID of the application cluster
	if err := firewall.CreateSecurityGroupRule(ctx, r.Client, firewall.RuleInfo{
		DestinationCidr: *cidr,
		Region:          managedCP.Spec.Region,
		SecurityGroup:   sgID,
		ToPort:          firstPort,
		FromPort:        lastPort,
	}); err != nil {
		return ctrl.Result{}, err
	}

	klog.Infof("security group rule created...")

	return ctrl.Result{}, nil
}

//+kubebuilder:rbac:groups=db.appscode.com,resources=crossplanes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=db.appscode.com,resources=crossplanes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=db.appscode.com,resources=crossplanes/finalizers,verbs=update

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	objKey := req.NamespacedName
	managedCP := &ekscontrolplanev1.AWSManagedControlPlane{}

	if err := r.Get(ctx, objKey, managedCP); err != nil {
		return ctrl.Result{}, err
	}

	if managedCP.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	if managedCP.Status.Ready != true {
		return ctrl.Result{}, nil
	}

	securityGroupID, err := firewall.GetSecurityGroupID(managedCP)
	if err != nil {
		return ctrl.Result{}, err
	}
	sgRule, err := firewall.GetRule(firewall.RuleInfo{
		DestinationCidr: "0.0.0.0/0",
		Region:          managedCP.Spec.Region,
		SecurityGroup:   securityGroupID,
		ToPort:          firstPort,
		FromPort:        lastPort,
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	_, _, err = kmc.CreateOrPatch(ctx, r.Client, sgRule, func(_ client.Object, _ bool) client.Object {
		return sgRule
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	klog.Info("Successfully updated sg rules")

	firewall.GetVpcIDs(ctx, r.Client)

	/*

		var routeTables []string
		for _, subnet := range managedCP.Spec.NetworkSpec.Subnets {
			if subnet.IsPublic == true {
				routeTables = append(routeTables, subnet.ID)
			}
		}

		destinationCIDR := "0.0.0.0/0"

		for _, tableID := range routeTables {
			var route upEC2.Route
			if err := r.Get(ctx, client.ObjectKey{Name: firewall.GetRouteName(tableID, destinationCIDR)}, &route); err != nil && client.IgnoreNotFound(err) == nil {
				route = *firewall.GetRoute(firewall.RouteInfo{
					RouteTable:      tableID,
					Destination:     destinationCIDR,
					Region:          managedCP.Spec.Region,
					InternetGateway: *managedCP.Spec.NetworkSpec.VPC.InternetGatewayID,
				})
				if er := r.Client.Create(ctx, &route); er != nil {
					return ctrl.Result{}, er
				}
				klog.Infof("route: %s created", firewall.GetRouteName(tableID, destinationCIDR))
			}
		}

	*/

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ekscontrolplanev1.AWSManagedControlPlane{}).
		Complete(r)
}

func (r *PCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&upEC2.VPCPeeringConnection{}).
		Complete(r)
}
