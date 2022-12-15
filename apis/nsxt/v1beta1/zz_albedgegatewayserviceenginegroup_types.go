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

type ALBEdgeGatewayServiceEngineGroupObservation struct {

	// Number of deployed Virtual Services on this Service Engine Group.
	// Number of deployed virtual services for this Service Engine Group
	DeployedVirtualServices *float64 `json:"deployedVirtualServices,omitempty" tf:"deployed_virtual_services,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Service Engine Group Name which is attached to NSX-T Edge Gateway
	ServiceEngineGroupName *string `json:"serviceEngineGroupName,omitempty" tf:"service_engine_group_name,omitempty"`
}

type ALBEdgeGatewayServiceEngineGroupParameters struct {

	// An ID of NSX-T Edge Gateway. Can be looked up using
	// vcd_nsxt_edgegateway data source.
	// Edge Gateway ID in which ALB Service Engine Group should be located
	// +crossplane:generate:reference:type=EdgeGateway
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Reference to a EdgeGateway to populate edgeGatewayId.
	// +kubebuilder:validation:Optional
	EdgeGatewayIDRef *v1.Reference `json:"edgeGatewayIdRef,omitempty" tf:"-"`

	// Selector for a EdgeGateway to populate edgeGatewayId.
	// +kubebuilder:validation:Optional
	EdgeGatewayIDSelector *v1.Selector `json:"edgeGatewayIdSelector,omitempty" tf:"-"`

	// Maximum amount of Virtual Services to run on this Service Engine Group. Only for
	// Shared Service Engine Groups.
	// Maximum number of virtual services to be used in this Service Engine Group
	// +kubebuilder:validation:Optional
	MaxVirtualServices *float64 `json:"maxVirtualServices,omitempty" tf:"max_virtual_services,omitempty"`

	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Number of reserved Virtual Services for this Edge Gateway. Only for Shared
	// Service Engine Groups.
	// Number of reserved virtual services for this Service Engine Group
	// +kubebuilder:validation:Optional
	ReservedVirtualServices *string `json:"reservedVirtualServices,omitempty" tf:"reserved_virtual_services,omitempty"`

	// An ID of NSX-T Service Engine Group. Can be looked up using
	// vcd_nsxt_alb_service_engine_group data
	// source.
	// Service Engine Group ID to attach to this NSX-T Edge Gateway
	// +crossplane:generate:reference:type=ALBServiceEngineGroup
	// +kubebuilder:validation:Optional
	ServiceEngineGroupID *string `json:"serviceEngineGroupId,omitempty" tf:"service_engine_group_id,omitempty"`

	// Reference to a ALBServiceEngineGroup to populate serviceEngineGroupId.
	// +kubebuilder:validation:Optional
	ServiceEngineGroupIDRef *v1.Reference `json:"serviceEngineGroupIdRef,omitempty" tf:"-"`

	// Selector for a ALBServiceEngineGroup to populate serviceEngineGroupId.
	// +kubebuilder:validation:Optional
	ServiceEngineGroupIDSelector *v1.Selector `json:"serviceEngineGroupIdSelector,omitempty" tf:"-"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// ALBEdgeGatewayServiceEngineGroupSpec defines the desired state of ALBEdgeGatewayServiceEngineGroup
type ALBEdgeGatewayServiceEngineGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ALBEdgeGatewayServiceEngineGroupParameters `json:"forProvider"`
}

// ALBEdgeGatewayServiceEngineGroupStatus defines the observed state of ALBEdgeGatewayServiceEngineGroup.
type ALBEdgeGatewayServiceEngineGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ALBEdgeGatewayServiceEngineGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ALBEdgeGatewayServiceEngineGroup is the Schema for the ALBEdgeGatewayServiceEngineGroups API. Provides a resource to manage NSX-T ALB Service Engine Group assignment to Edge Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type ALBEdgeGatewayServiceEngineGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ALBEdgeGatewayServiceEngineGroupSpec   `json:"spec"`
	Status            ALBEdgeGatewayServiceEngineGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ALBEdgeGatewayServiceEngineGroupList contains a list of ALBEdgeGatewayServiceEngineGroups
type ALBEdgeGatewayServiceEngineGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ALBEdgeGatewayServiceEngineGroup `json:"items"`
}

// Repository type metadata.
var (
	ALBEdgeGatewayServiceEngineGroup_Kind             = "ALBEdgeGatewayServiceEngineGroup"
	ALBEdgeGatewayServiceEngineGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ALBEdgeGatewayServiceEngineGroup_Kind}.String()
	ALBEdgeGatewayServiceEngineGroup_KindAPIVersion   = ALBEdgeGatewayServiceEngineGroup_Kind + "." + CRDGroupVersion.String()
	ALBEdgeGatewayServiceEngineGroup_GroupVersionKind = CRDGroupVersion.WithKind(ALBEdgeGatewayServiceEngineGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ALBEdgeGatewayServiceEngineGroup{}, &ALBEdgeGatewayServiceEngineGroupList{})
}
