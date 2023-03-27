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

type VPCEndpointSubnetAssociationObservation struct {

	// The ID of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPCEndpointSubnetAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the subnet to be associated with the VPC endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The ID of the VPC endpoint with which the subnet will be associated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// Reference to a VPCEndpoint in ec2 to populate vpcEndpointId.
	// +kubebuilder:validation:Optional
	VPCEndpointIDRef *v1.Reference `json:"vpcEndpointIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpoint in ec2 to populate vpcEndpointId.
	// +kubebuilder:validation:Optional
	VPCEndpointIDSelector *v1.Selector `json:"vpcEndpointIdSelector,omitempty" tf:"-"`
}

// VPCEndpointSubnetAssociationSpec defines the desired state of VPCEndpointSubnetAssociation
type VPCEndpointSubnetAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointSubnetAssociationParameters `json:"forProvider"`
}

// VPCEndpointSubnetAssociationStatus defines the observed state of VPCEndpointSubnetAssociation.
type VPCEndpointSubnetAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointSubnetAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointSubnetAssociation is the Schema for the VPCEndpointSubnetAssociations API. Provides a resource to create an association between a VPC endpoint and a subnet.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointSubnetAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointSubnetAssociationSpec   `json:"spec"`
	Status            VPCEndpointSubnetAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointSubnetAssociationList contains a list of VPCEndpointSubnetAssociations
type VPCEndpointSubnetAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointSubnetAssociation `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointSubnetAssociation_Kind             = "VPCEndpointSubnetAssociation"
	VPCEndpointSubnetAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointSubnetAssociation_Kind}.String()
	VPCEndpointSubnetAssociation_KindAPIVersion   = VPCEndpointSubnetAssociation_Kind + "." + CRDGroupVersion.String()
	VPCEndpointSubnetAssociation_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointSubnetAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointSubnetAssociation{}, &VPCEndpointSubnetAssociationList{})
}
