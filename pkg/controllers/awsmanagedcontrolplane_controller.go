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

package controllers

import (
	"context"
	"errors"

	"go.bytebuilders.dev/capa-vpc-peering-operator/pkg/firewall"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ec2api "kubeform.dev/provider-aws/apis/ec2/v1alpha1"
	eksapi "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta1"
	expapi "sigs.k8s.io/cluster-api-provider-aws/v2/exp/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// AWSManagedControlPlaneReconciler reconciles a Crossplane object
type AWSManagedControlPlaneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *AWSManagedControlPlaneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	objKey := req.NamespacedName
	managedCP := &eksapi.AWSManagedControlPlane{}

	if err := r.Get(ctx, objKey, managedCP); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if managedCP.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	if !managedCP.Status.Ready {
		return ctrl.Result{}, nil
	}

	pcs := &ec2api.VPCPeeringConnectionList{}
	if err := r.Client.List(ctx, pcs); err != nil {
		return ctrl.Result{}, err
	}

	if len(pcs.Items) == 0 {
		return ctrl.Result{}, nil
	}

	securityGroupID, err := firewall.GetSecurityGroupID(managedCP)
	if err != nil {
		return ctrl.Result{}, err
	}

	var routeTableIDs []string
	for _, subnet := range managedCP.Spec.NetworkSpec.Subnets {
		if !subnet.IsPublic {
			routeTableIDs = append(routeTableIDs, *subnet.RouteTableID)
		}
	}

	var retErr error

	for _, pc := range pcs.Items {
		if !firewall.CheckCrossplaneCondition(pc.Status.Conditions) {
			klog.Infof("%s condition false", pc.GetID())
			continue
		}
		klog.Infof("====================== %s is ready ====================", pc.GetID())

		cidr := pc.ObjectMeta.Annotations[firewall.CidrAnnotation]

		if err = firewall.CreateSecurityGroupRule(ctx, r.Client, firewall.RuleInfo{
			DestinationCidr: cidr,
			Region:          managedCP.Spec.Region,
			SecurityGroup:   securityGroupID,
			ToPort:          firewall.LastPort,
			FromPort:        firewall.FirstPort,
		}, firewall.GetOwnerReference(&pc)); err != nil {
			klog.Errorf("failed to add rule in %s for %s", securityGroupID, pc.GetID())
			retErr = errors.Join(retErr, err)
		}

		for _, tableID := range routeTableIDs {
			if err = firewall.CreateRouteTableRoute(ctx, r.Client, firewall.RouteInfo{
				RouteTable:          tableID,
				Destination:         cidr,
				Region:              managedCP.Spec.Region,
				PeeringConnectionID: pc.GetID(),
			}, firewall.GetOwnerReference(&pc)); err != nil {
				klog.Errorf("failed to add route in %s for %s", tableID, pc.GetID())
				retErr = errors.Join(retErr, err)
			}
		}
	}

	return ctrl.Result{}, retErr
}

// SetupWithManager sets up the controller with the Manager.
func (r *AWSManagedControlPlaneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&eksapi.AWSManagedControlPlane{}).
		Watches(
			&source.Kind{Type: &expapi.AWSManagedMachinePool{}},
			handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
				reconcileReq := make([]reconcile.Request, 0)
				managedCP, err := firewall.GetManagedControlPlane(context.TODO(), r.Client)
				if err != nil {
					return reconcileReq
				}
				reconcileReq = append(reconcileReq, reconcile.Request{NamespacedName: client.ObjectKey{Name: managedCP.Name, Namespace: managedCP.Namespace}})
				return reconcileReq
			}),
		).
		Complete(r)
}
