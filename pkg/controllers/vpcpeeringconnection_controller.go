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

	"go.bytebuilders.dev/capa-vpc-peering-operator/pkg/firewall"

	"github.com/pkg/errors"
	ec2api "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	cpv1beta1 "sigs.k8s.io/cluster-api-provider-aws/v2/controlplane/eks/api/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// VPCPeeringConnectionReconciler reconciles a crossplane vpc peering connection object
type VPCPeeringConnectionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *VPCPeeringConnectionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	objKey := req.NamespacedName
	pc := &ec2api.VPCPeeringConnection{}

	if err := r.Get(ctx, objKey, pc); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if pc.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	if !firewall.IsConditionReady(pc.Status.Conditions) || len(pc.GetID()) == 0 {
		klog.Infof("%s condition false", pc.GetID())
		return ctrl.Result{}, nil
	}

	klog.Infof("====================== %s is ready ====================", pc.GetID())

	managedCP, err := firewall.GetManagedControlPlane(ctx, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}
	sgID, err := firewall.GetSecurityGroupID(managedCP)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Note: expected application VPC CIDR range in pc.spec.forProvider.tags[cidr] section, tags is a map[string]*string type
	cidr, found := pc.ObjectMeta.Annotations[firewall.CidrAnnotation]
	if !found {
		return ctrl.Result{}, errors.New("empty CIDR range in peering connection tags")
	}

	if err := firewall.CreateOrPatchRule(ctx, r.Client, firewall.RuleInfo{
		DestinationCidr: cidr,
		Region:          managedCP.Spec.Region,
		SecurityGroup:   sgID,
		FromPort:        firewall.FirstPort,
		ToPort:          firewall.LastPort,
	}, firewall.GetOwnerReference(pc)); err != nil {
		return ctrl.Result{}, err
	}
	klog.Infof("security group rule created...")

	var routeTableIDs []string
	for _, subnet := range managedCP.Spec.NetworkSpec.Subnets {
		if !subnet.IsPublic {
			routeTableIDs = append(routeTableIDs, *subnet.RouteTableID)
		}
	}
	var retErr error
	for _, tableID := range routeTableIDs {
		if err := firewall.CreateOrPatchRoute(ctx, r.Client, firewall.RouteInfo{
			RouteTable:          tableID,
			Destination:         cidr,
			Region:              managedCP.Spec.Region,
			PeeringConnectionID: pc.GetID(),
		}, firewall.GetOwnerReference(pc)); err != nil {
			retErr = errors.Wrap(retErr, err.Error())
		}
	}

	return ctrl.Result{}, retErr
}

// SetupWithManager sets up the controller with the Manager.
func (r *VPCPeeringConnectionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ec2api.VPCPeeringConnection{}).
		Watches(
			&source.Kind{Type: &cpv1beta1.AWSManagedControlPlane{}},
			handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
				reconcileReq := make([]reconcile.Request, 0)
				pcs := &ec2api.VPCPeeringConnectionList{}
				err := r.List(context.TODO(), pcs)
				if err != nil {
					return reconcileReq
				}

				for _, pc := range pcs.Items {
					reconcileReq = append(reconcileReq, reconcile.Request{NamespacedName: client.ObjectKey{Name: pc.Name, Namespace: pc.Namespace}})
				}

				return reconcileReq
			}),
		).
		Complete(r)
}
