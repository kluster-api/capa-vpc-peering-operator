//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccepterObservation) DeepCopyInto(out *AccepterObservation) {
	*out = *in
	if in.AllowRemoteVPCDNSResolution != nil {
		in, out := &in.AllowRemoteVPCDNSResolution, &out.AllowRemoteVPCDNSResolution
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccepterObservation.
func (in *AccepterObservation) DeepCopy() *AccepterObservation {
	if in == nil {
		return nil
	}
	out := new(AccepterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccepterParameters) DeepCopyInto(out *AccepterParameters) {
	*out = *in
	if in.AllowRemoteVPCDNSResolution != nil {
		in, out := &in.AllowRemoteVPCDNSResolution, &out.AllowRemoteVPCDNSResolution
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccepterParameters.
func (in *AccepterParameters) DeepCopy() *AccepterParameters {
	if in == nil {
		return nil
	}
	out := new(AccepterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequesterObservation) DeepCopyInto(out *RequesterObservation) {
	*out = *in
	if in.AllowRemoteVPCDNSResolution != nil {
		in, out := &in.AllowRemoteVPCDNSResolution, &out.AllowRemoteVPCDNSResolution
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequesterObservation.
func (in *RequesterObservation) DeepCopy() *RequesterObservation {
	if in == nil {
		return nil
	}
	out := new(RequesterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequesterParameters) DeepCopyInto(out *RequesterParameters) {
	*out = *in
	if in.AllowRemoteVPCDNSResolution != nil {
		in, out := &in.AllowRemoteVPCDNSResolution, &out.AllowRemoteVPCDNSResolution
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequesterParameters.
func (in *RequesterParameters) DeepCopy() *RequesterParameters {
	if in == nil {
		return nil
	}
	out := new(RequesterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Route) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteList) DeepCopyInto(out *RouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteList.
func (in *RouteList) DeepCopy() *RouteList {
	if in == nil {
		return nil
	}
	out := new(RouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteObservation) DeepCopyInto(out *RouteObservation) {
	*out = *in
	if in.CarrierGatewayID != nil {
		in, out := &in.CarrierGatewayID, &out.CarrierGatewayID
		*out = new(string)
		**out = **in
	}
	if in.CoreNetworkArn != nil {
		in, out := &in.CoreNetworkArn, &out.CoreNetworkArn
		*out = new(string)
		**out = **in
	}
	if in.DestinationCidrBlock != nil {
		in, out := &in.DestinationCidrBlock, &out.DestinationCidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationIPv6CidrBlock != nil {
		in, out := &in.DestinationIPv6CidrBlock, &out.DestinationIPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationPrefixListID != nil {
		in, out := &in.DestinationPrefixListID, &out.DestinationPrefixListID
		*out = new(string)
		**out = **in
	}
	if in.EgressOnlyGatewayID != nil {
		in, out := &in.EgressOnlyGatewayID, &out.EgressOnlyGatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceOwnerID != nil {
		in, out := &in.InstanceOwnerID, &out.InstanceOwnerID
		*out = new(string)
		**out = **in
	}
	if in.LocalGatewayID != nil {
		in, out := &in.LocalGatewayID, &out.LocalGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TransitGatewayID != nil {
		in, out := &in.TransitGatewayID, &out.TransitGatewayID
		*out = new(string)
		**out = **in
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCPeeringConnectionID != nil {
		in, out := &in.VPCPeeringConnectionID, &out.VPCPeeringConnectionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteObservation.
func (in *RouteObservation) DeepCopy() *RouteObservation {
	if in == nil {
		return nil
	}
	out := new(RouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteParameters) DeepCopyInto(out *RouteParameters) {
	*out = *in
	if in.CarrierGatewayID != nil {
		in, out := &in.CarrierGatewayID, &out.CarrierGatewayID
		*out = new(string)
		**out = **in
	}
	if in.CoreNetworkArn != nil {
		in, out := &in.CoreNetworkArn, &out.CoreNetworkArn
		*out = new(string)
		**out = **in
	}
	if in.DestinationCidrBlock != nil {
		in, out := &in.DestinationCidrBlock, &out.DestinationCidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationIPv6CidrBlock != nil {
		in, out := &in.DestinationIPv6CidrBlock, &out.DestinationIPv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DestinationPrefixListID != nil {
		in, out := &in.DestinationPrefixListID, &out.DestinationPrefixListID
		*out = new(string)
		**out = **in
	}
	if in.EgressOnlyGatewayID != nil {
		in, out := &in.EgressOnlyGatewayID, &out.EgressOnlyGatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.LocalGatewayID != nil {
		in, out := &in.LocalGatewayID, &out.LocalGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NATGatewayID != nil {
		in, out := &in.NATGatewayID, &out.NATGatewayID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RouteTableID != nil {
		in, out := &in.RouteTableID, &out.RouteTableID
		*out = new(string)
		**out = **in
	}
	if in.TransitGatewayID != nil {
		in, out := &in.TransitGatewayID, &out.TransitGatewayID
		*out = new(string)
		**out = **in
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCPeeringConnectionID != nil {
		in, out := &in.VPCPeeringConnectionID, &out.VPCPeeringConnectionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteParameters.
func (in *RouteParameters) DeepCopy() *RouteParameters {
	if in == nil {
		return nil
	}
	out := new(RouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRule) DeepCopyInto(out *SecurityGroupRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRule.
func (in *SecurityGroupRule) DeepCopy() *SecurityGroupRule {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroupRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRuleList) DeepCopyInto(out *SecurityGroupRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityGroupRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRuleList.
func (in *SecurityGroupRuleList) DeepCopy() *SecurityGroupRuleList {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroupRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRuleObservation) DeepCopyInto(out *SecurityGroupRuleObservation) {
	*out = *in
	if in.CidrBlocks != nil {
		in, out := &in.CidrBlocks, &out.CidrBlocks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FromPort != nil {
		in, out := &in.FromPort, &out.FromPort
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPv6CidrBlocks != nil {
		in, out := &in.IPv6CidrBlocks, &out.IPv6CidrBlocks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrefixListIds != nil {
		in, out := &in.PrefixListIds, &out.PrefixListIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupRuleID != nil {
		in, out := &in.SecurityGroupRuleID, &out.SecurityGroupRuleID
		*out = new(string)
		**out = **in
	}
	if in.Self != nil {
		in, out := &in.Self, &out.Self
		*out = new(bool)
		**out = **in
	}
	if in.SourceSecurityGroupID != nil {
		in, out := &in.SourceSecurityGroupID, &out.SourceSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ToPort != nil {
		in, out := &in.ToPort, &out.ToPort
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRuleObservation.
func (in *SecurityGroupRuleObservation) DeepCopy() *SecurityGroupRuleObservation {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRuleParameters) DeepCopyInto(out *SecurityGroupRuleParameters) {
	*out = *in
	if in.CidrBlocks != nil {
		in, out := &in.CidrBlocks, &out.CidrBlocks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FromPort != nil {
		in, out := &in.FromPort, &out.FromPort
		*out = new(float64)
		**out = **in
	}
	if in.IPv6CidrBlocks != nil {
		in, out := &in.IPv6CidrBlocks, &out.IPv6CidrBlocks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrefixListIds != nil {
		in, out := &in.PrefixListIds, &out.PrefixListIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.Self != nil {
		in, out := &in.Self, &out.Self
		*out = new(bool)
		**out = **in
	}
	if in.SourceSecurityGroupID != nil {
		in, out := &in.SourceSecurityGroupID, &out.SourceSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ToPort != nil {
		in, out := &in.ToPort, &out.ToPort
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRuleParameters.
func (in *SecurityGroupRuleParameters) DeepCopy() *SecurityGroupRuleParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRuleSpec) DeepCopyInto(out *SecurityGroupRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRuleSpec.
func (in *SecurityGroupRuleSpec) DeepCopy() *SecurityGroupRuleSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRuleStatus) DeepCopyInto(out *SecurityGroupRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRuleStatus.
func (in *SecurityGroupRuleStatus) DeepCopy() *SecurityGroupRuleStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCPeeringConnection) DeepCopyInto(out *VPCPeeringConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCPeeringConnection.
func (in *VPCPeeringConnection) DeepCopy() *VPCPeeringConnection {
	if in == nil {
		return nil
	}
	out := new(VPCPeeringConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCPeeringConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCPeeringConnectionList) DeepCopyInto(out *VPCPeeringConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCPeeringConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCPeeringConnectionList.
func (in *VPCPeeringConnectionList) DeepCopy() *VPCPeeringConnectionList {
	if in == nil {
		return nil
	}
	out := new(VPCPeeringConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCPeeringConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCPeeringConnectionObservation) DeepCopyInto(out *VPCPeeringConnectionObservation) {
	*out = *in
	if in.AcceptStatus != nil {
		in, out := &in.AcceptStatus, &out.AcceptStatus
		*out = new(string)
		**out = **in
	}
	if in.Accepter != nil {
		in, out := &in.Accepter, &out.Accepter
		*out = make([]AccepterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoAccept != nil {
		in, out := &in.AutoAccept, &out.AutoAccept
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PeerOwnerID != nil {
		in, out := &in.PeerOwnerID, &out.PeerOwnerID
		*out = new(string)
		**out = **in
	}
	if in.PeerRegion != nil {
		in, out := &in.PeerRegion, &out.PeerRegion
		*out = new(string)
		**out = **in
	}
	if in.PeerVPCID != nil {
		in, out := &in.PeerVPCID, &out.PeerVPCID
		*out = new(string)
		**out = **in
	}
	if in.Requester != nil {
		in, out := &in.Requester, &out.Requester
		*out = make([]RequesterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCPeeringConnectionObservation.
func (in *VPCPeeringConnectionObservation) DeepCopy() *VPCPeeringConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(VPCPeeringConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCPeeringConnectionParameters) DeepCopyInto(out *VPCPeeringConnectionParameters) {
	*out = *in
	if in.Accepter != nil {
		in, out := &in.Accepter, &out.Accepter
		*out = make([]AccepterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoAccept != nil {
		in, out := &in.AutoAccept, &out.AutoAccept
		*out = new(bool)
		**out = **in
	}
	if in.PeerOwnerID != nil {
		in, out := &in.PeerOwnerID, &out.PeerOwnerID
		*out = new(string)
		**out = **in
	}
	if in.PeerRegion != nil {
		in, out := &in.PeerRegion, &out.PeerRegion
		*out = new(string)
		**out = **in
	}
	if in.PeerVPCID != nil {
		in, out := &in.PeerVPCID, &out.PeerVPCID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Requester != nil {
		in, out := &in.Requester, &out.Requester
		*out = make([]RequesterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCPeeringConnectionParameters.
func (in *VPCPeeringConnectionParameters) DeepCopy() *VPCPeeringConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(VPCPeeringConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCPeeringConnectionSpec) DeepCopyInto(out *VPCPeeringConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCPeeringConnectionSpec.
func (in *VPCPeeringConnectionSpec) DeepCopy() *VPCPeeringConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(VPCPeeringConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCPeeringConnectionStatus) DeepCopyInto(out *VPCPeeringConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCPeeringConnectionStatus.
func (in *VPCPeeringConnectionStatus) DeepCopy() *VPCPeeringConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(VPCPeeringConnectionStatus)
	in.DeepCopyInto(out)
	return out
}
