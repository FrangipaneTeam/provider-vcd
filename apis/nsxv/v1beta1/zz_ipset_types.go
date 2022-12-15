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

type IPSetObservation struct {

	// ID of IP set
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IPSetParameters struct {

	// An optional description for IP set.
	// IP set description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of IP addresses, CIDRs and ranges as strings.
	// A set of IP address, CIDR, IP range objects
	// +kubebuilder:validation:Required
	IPAddresses []*string `json:"ipAddresses" tf:"ip_addresses,omitempty"`

	// Toggle to enable inheritance to allow visibility at underlying scopes. Default true
	// Allows visibility in underlying scopes (Default is true)
	// +kubebuilder:validation:Optional
	IsInheritanceAllowed *bool `json:"isInheritanceAllowed,omitempty" tf:"is_inheritance_allowed,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// IPSetSpec defines the desired state of IPSet
type IPSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPSetParameters `json:"forProvider"`
}

// IPSetStatus defines the observed state of IPSet.
type IPSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPSet is the Schema for the IPSets API. Provides an IP set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPSetSpec   `json:"spec"`
	Status            IPSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPSetList contains a list of IPSets
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPSet `json:"items"`
}

// Repository type metadata.
var (
	IPSet_Kind             = "IPSet"
	IPSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPSet_Kind}.String()
	IPSet_KindAPIVersion   = IPSet_Kind + "." + CRDGroupVersion.String()
	IPSet_GroupVersionKind = CRDGroupVersion.WithKind(IPSet_Kind)
)

func init() {
	SchemeBuilder.Register(&IPSet{}, &IPSetList{})
}
