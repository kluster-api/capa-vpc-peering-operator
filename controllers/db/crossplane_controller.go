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
	"k8s.io/apimachinery/pkg/util/managedfields"
	"k8s.io/klog/v2"
	"kubedb/aws-peering-connection-operator/pkg/firewall"

	"k8s.io/apimachinery/pkg/runtime"
	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	expv1beta1 "sigs.k8s.io/cluster-api-provider-aws/exp/api/v1beta1"
)

// Reconciler reconciles a Crossplane object
type Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type SecurityGroups

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

	securityGroupID := managedCP.Status.Network.SecurityGroups[infrav1.SecurityGroupEKSNodeAdditional].ID
	if securityGroupID == "" {
		return ctrl.Result{},errors.New("no security group id found")
	}

	var routeTables []string
	for _,subnet := range managedCP.Spec.NetworkSpec.Subnets {
		if subnet.IsPublic == true {
			routeTables = append(routeTables, subnet.ID)
		}
	}

	destinationCIDR := "0.0.0.0/0"

	for _,tableID := range routeTables {
		var route upEC2.Route
		if err := r.Get(ctx,client.ObjectKey{Name: firewall.GetRouteName(tableID,destinationCIDR)},&route); err != nil && client.IgnoreNotFound(err) == nil{
			route = *firewall.GetRoute(firewall.RouteInfo{
				RouteTable:      tableID,
				Destination:     destinationCIDR,
				Region:          managedCP.Spec.Region,
				InternetGateway: *managedCP.Spec.NetworkSpec.VPC.InternetGatewayID,
			})
			if er := r.Client.Create(ctx,&route); er != nil {
				return ctrl.Result{},er
			}
			klog.Infof("route: %s created",firewall.GetRouteName(tableID,destinationCIDR))
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ekscontrolplanev1.AWSManagedControlPlane{}).
		Complete(r)
}
