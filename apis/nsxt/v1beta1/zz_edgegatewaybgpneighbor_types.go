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

type EdgeGatewayBGPNeighborObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EdgeGatewayBGPNeighborParameters struct {

	// BGP Allow-as-in feature is used to allow the BGP speaker to accept the BGP updates even if its own BGP AS number is in the AS-Path attribute.
	// A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false')
	// +kubebuilder:validation:Optional
	AllowAsIn *bool `json:"allowAsIn,omitempty" tf:"allow_as_in,omitempty"`

	// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
	// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
	// +kubebuilder:validation:Optional
	BfdDeadMultiple *float64 `json:"bfdDeadMultiple,omitempty" tf:"bfd_dead_multiple,omitempty"`

	// Should Bidirectional Forwarding Detection (BFD) be enabled
	// BFD configuration for failure detection
	// +kubebuilder:validation:Optional
	BfdEnabled *bool `json:"bfdEnabled,omitempty" tf:"bfd_enabled,omitempty"`

	// Time interval (in milliseconds) between heartbeat packets
	// Time interval (in milliseconds) between heartbeat packets
	// +kubebuilder:validation:Optional
	BfdInterval *float64 `json:"bfdInterval,omitempty" tf:"bfd_interval,omitempty"`

	// The ID of the edge gateway (NSX-T only). Can be looked up using
	// vcd_nsxt_edgegateway datasource
	// Edge gateway ID for BGP Neighbor Configuration
	// +crossplane:generate:reference:type=EdgeGateway
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Reference to a EdgeGateway to populate edgeGatewayId.
	// +kubebuilder:validation:Optional
	EdgeGatewayIDRef *v1.Reference `json:"edgeGatewayIdRef,omitempty" tf:"-"`

	// Selector for a EdgeGateway to populate edgeGatewayId.
	// +kubebuilder:validation:Optional
	EdgeGatewayIDSelector *v1.Selector `json:"edgeGatewayIdSelector,omitempty" tf:"-"`

	// BGP Neighbor Graceful Restart Mode. One of:
	// One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
	// +kubebuilder:validation:Optional
	GracefulRestartMode *string `json:"gracefulRestartMode,omitempty" tf:"graceful_restart_mode,omitempty"`

	// Time interval (in seconds) before declaring a BGP peer dead
	// Time interval (in seconds) before declaring a peer dead
	// +kubebuilder:validation:Optional
	HoldDownTimer *float64 `json:"holdDownTimer,omitempty" tf:"hold_down_timer,omitempty"`

	// BGP Neighbor IP Address (IPv4 or IPv6)
	// BGP Neighbor IP address (IPv4 or IPv6)
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The ID of the IP Prefix List to be used for filtering incoming BGP routes
	// An optional IP Prefix List ID for filtering 'IN' direction.
	// +crossplane:generate:reference:type=EdgeGatewayBGPIPPrefixList
	// +kubebuilder:validation:Optional
	InFilterIPPrefixListID *string `json:"inFilterIpPrefixListId,omitempty" tf:"in_filter_ip_prefix_list_id,omitempty"`

	// Reference to a EdgeGatewayBGPIPPrefixList to populate inFilterIpPrefixListId.
	// +kubebuilder:validation:Optional
	InFilterIPPrefixListIDRef *v1.Reference `json:"inFilterIpPrefixListIdRef,omitempty" tf:"-"`

	// Selector for a EdgeGatewayBGPIPPrefixList to populate inFilterIpPrefixListId.
	// +kubebuilder:validation:Optional
	InFilterIPPrefixListIDSelector *v1.Selector `json:"inFilterIpPrefixListIdSelector,omitempty" tf:"-"`

	// Time interval (in seconds) between sending keep-alive messages to a BGP peer
	// Time interval (in seconds) between sending keep alive messages to a peer
	// +kubebuilder:validation:Optional
	KeepAliveTimer *float64 `json:"keepAliveTimer,omitempty" tf:"keep_alive_timer,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The ID of the IP Prefix List to be used for filtering outgoing BGP routes
	// An optional IP Prefix List ID for filtering 'OUT' direction.
	// +crossplane:generate:reference:type=EdgeGatewayBGPIPPrefixList
	// +kubebuilder:validation:Optional
	OutFilterIPPrefixListID *string `json:"outFilterIpPrefixListId,omitempty" tf:"out_filter_ip_prefix_list_id,omitempty"`

	// Reference to a EdgeGatewayBGPIPPrefixList to populate outFilterIpPrefixListId.
	// +kubebuilder:validation:Optional
	OutFilterIPPrefixListIDRef *v1.Reference `json:"outFilterIpPrefixListIdRef,omitempty" tf:"-"`

	// Selector for a EdgeGatewayBGPIPPrefixList to populate outFilterIpPrefixListId.
	// +kubebuilder:validation:Optional
	OutFilterIPPrefixListIDSelector *v1.Selector `json:"outFilterIpPrefixListIdSelector,omitempty" tf:"-"`

	// BGP Neighbor Password
	// Neighbor password
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// BGP Neighbor Remote Autonomous System (AS) Number
	// Remote Autonomous System (AS) number
	// +kubebuilder:validation:Required
	RemoteAsNumber *string `json:"remoteAsNumber" tf:"remote_as_number,omitempty"`

	// Route filtering by IP Address family. One of DISABLED, IPV4, IPV6
	// One of 'DISABLED', 'IPV4', 'IPV6'
	// +kubebuilder:validation:Optional
	RouteFiltering *string `json:"routeFiltering,omitempty" tf:"route_filtering,omitempty"`
}

// EdgeGatewayBGPNeighborSpec defines the desired state of EdgeGatewayBGPNeighbor
type EdgeGatewayBGPNeighborSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EdgeGatewayBGPNeighborParameters `json:"forProvider"`
}

// EdgeGatewayBGPNeighborStatus defines the observed state of EdgeGatewayBGPNeighbor.
type EdgeGatewayBGPNeighborStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EdgeGatewayBGPNeighborObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeGatewayBGPNeighbor is the Schema for the EdgeGatewayBGPNeighbors API. Provides a resource to manage NSX-T Edge Gateway BGP Neighbors and their configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type EdgeGatewayBGPNeighbor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeGatewayBGPNeighborSpec   `json:"spec"`
	Status            EdgeGatewayBGPNeighborStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeGatewayBGPNeighborList contains a list of EdgeGatewayBGPNeighbors
type EdgeGatewayBGPNeighborList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeGatewayBGPNeighbor `json:"items"`
}

// Repository type metadata.
var (
	EdgeGatewayBGPNeighbor_Kind             = "EdgeGatewayBGPNeighbor"
	EdgeGatewayBGPNeighbor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EdgeGatewayBGPNeighbor_Kind}.String()
	EdgeGatewayBGPNeighbor_KindAPIVersion   = EdgeGatewayBGPNeighbor_Kind + "." + CRDGroupVersion.String()
	EdgeGatewayBGPNeighbor_GroupVersionKind = CRDGroupVersion.WithKind(EdgeGatewayBGPNeighbor_Kind)
)

func init() {
	SchemeBuilder.Register(&EdgeGatewayBGPNeighbor{}, &EdgeGatewayBGPNeighborList{})
}
