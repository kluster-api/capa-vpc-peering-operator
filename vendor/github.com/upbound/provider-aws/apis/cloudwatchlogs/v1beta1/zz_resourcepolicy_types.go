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

type ResourcePolicyObservation struct {

	// The name of the CloudWatch log resource policy
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResourcePolicyParameters struct {

	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	// +kubebuilder:validation:Required
	PolicyDocument *string `json:"policyDocument" tf:"policy_document,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ResourcePolicySpec defines the desired state of ResourcePolicy
type ResourcePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyParameters `json:"forProvider"`
}

// ResourcePolicyStatus defines the observed state of ResourcePolicy.
type ResourcePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicy is the Schema for the ResourcePolicys API. Provides a resource to manage a CloudWatch log resource policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcePolicySpec   `json:"spec"`
	Status            ResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyList contains a list of ResourcePolicys
type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicy `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicy_Kind             = "ResourcePolicy"
	ResourcePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourcePolicy_Kind}.String()
	ResourcePolicy_KindAPIVersion   = ResourcePolicy_Kind + "." + CRDGroupVersion.String()
	ResourcePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ResourcePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicy{}, &ResourcePolicyList{})
}
