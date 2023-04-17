/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SnapshotCreateVolumePermissionObservation struct {

	// A combination of "snapshot_id-account_id".
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SnapshotCreateVolumePermissionParameters struct {

	// An AWS Account ID to add create volume permissions. The AWS Account cannot be the snapshot's owner
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A snapshot ID
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EBSSnapshot
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Reference to a EBSSnapshot in ec2 to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDRef *v1.Reference `json:"snapshotIdRef,omitempty" tf:"-"`

	// Selector for a EBSSnapshot in ec2 to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDSelector *v1.Selector `json:"snapshotIdSelector,omitempty" tf:"-"`
}

// SnapshotCreateVolumePermissionSpec defines the desired state of SnapshotCreateVolumePermission
type SnapshotCreateVolumePermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotCreateVolumePermissionParameters `json:"forProvider"`
}

// SnapshotCreateVolumePermissionStatus defines the observed state of SnapshotCreateVolumePermission.
type SnapshotCreateVolumePermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotCreateVolumePermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCreateVolumePermission is the Schema for the SnapshotCreateVolumePermissions API. Adds create volume permission to an EBS Snapshot
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnapshotCreateVolumePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotCreateVolumePermissionSpec   `json:"spec"`
	Status            SnapshotCreateVolumePermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCreateVolumePermissionList contains a list of SnapshotCreateVolumePermissions
type SnapshotCreateVolumePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotCreateVolumePermission `json:"items"`
}

// Repository type metadata.
var (
	SnapshotCreateVolumePermission_Kind             = "SnapshotCreateVolumePermission"
	SnapshotCreateVolumePermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotCreateVolumePermission_Kind}.String()
	SnapshotCreateVolumePermission_KindAPIVersion   = SnapshotCreateVolumePermission_Kind + "." + CRDGroupVersion.String()
	SnapshotCreateVolumePermission_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotCreateVolumePermission_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotCreateVolumePermission{}, &SnapshotCreateVolumePermissionList{})
}
