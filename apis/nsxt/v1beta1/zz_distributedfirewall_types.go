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

type DistributedFirewallObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more blocks with Firewall Rule definitions. Order
	// defines firewall rule precedence
	// Ordered list of firewall rules
	// +kubebuilder:validation:Required
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type DistributedFirewallParameters struct {

	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// One or more blocks with Firewall Rule definitions. Order
	// defines firewall rule precedence
	// Ordered list of firewall rules
	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`

	// The ID of VDC Group to manage Distributed Firewall in. Can be looked
	// up using vcd_vdc_group resource or data source.
	// ID of VDC Group for Distributed Firewall
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-vcd/apis/org/v1beta1.Group
	// +kubebuilder:validation:Optional
	VdcGroupID *string `json:"vdcGroupId,omitempty" tf:"vdc_group_id,omitempty"`

	// Reference to a Group in org to populate vdcGroupId.
	// +kubebuilder:validation:Optional
	VdcGroupIDRef *v1.Reference `json:"vdcGroupIdRef,omitempty" tf:"-"`

	// Selector for a Group in org to populate vdcGroupId.
	// +kubebuilder:validation:Optional
	VdcGroupIDSelector *v1.Selector `json:"vdcGroupIdSelector,omitempty" tf:"-"`
}

type RuleObservation struct {

	// Firewall Rule ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleParameters struct {

	// Defines if it should ALLOW, DROP, REJECT traffic. REJECT is only
	// supported in VCD 10.2.2+
	// Defines if the rule should 'ALLOW', 'DROP', 'REJECT' matching traffic
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// An optional set of Application Port Profiles.
	// A set of Application Port Profile IDs. Leaving it empty means 'Any'
	// +crossplane:generate:reference:type=AppPortProfile
	// +kubebuilder:validation:Optional
	AppPortProfileIds []*string `json:"appPortProfileIds,omitempty" tf:"app_port_profile_ids,omitempty"`

	// References to AppPortProfile to populate appPortProfileIds.
	// +kubebuilder:validation:Optional
	AppPortProfileIdsRefs []v1.Reference `json:"appPortProfileIdsRefs,omitempty" tf:"-"`

	// Selector for a list of AppPortProfile to populate appPortProfileIds.
	// +kubebuilder:validation:Optional
	AppPortProfileIdsSelector *v1.Selector `json:"appPortProfileIdsSelector,omitempty" tf:"-"`

	// Comment field shown in UI
	// Comment that is shown next to rule in UI (VCD 10.3.2+)
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Description of firewall rule (not shown in UI)
	// Description is not shown in UI
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// - reverses value of destination_ids for
	// the rule to match everything except specified IDs.
	// Reverses firewall matching for to match all except Destinations Groups specified in 'destination_ids' (VCD 10.3.2+)
	// +kubebuilder:validation:Optional
	DestinationGroupsExcluded *bool `json:"destinationGroupsExcluded,omitempty" tf:"destination_groups_excluded,omitempty"`

	// A set of source object Firewall Groups (IP Sets or Security groups). Leaving it empty matches Any (all)
	// A set of Destination Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'
	// +kubebuilder:validation:Optional
	DestinationIds []*string `json:"destinationIds,omitempty" tf:"destination_ids,omitempty"`

	// One of IN, OUT, or IN_OUT. (default IN_OUT)
	// Direction on which Firewall Rule applies (One of 'IN', 'OUT', 'IN_OUT')
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Defines if the rule is enabled (default true)
	// Defined if Firewall Rule is active
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One of IPV4,  IPV6, or IPV4_IPV6 (default IPV4_IPV6)
	// Firewall Rule Protocol (One of 'IPV4', 'IPV6', 'IPV4_IPV6')
	// +kubebuilder:validation:Optional
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// Defines if logging for this rule is enabled (default false)
	// Defines if matching traffic should be logged
	// +kubebuilder:validation:Optional
	Logging *bool `json:"logging,omitempty" tf:"logging,omitempty"`

	// Explanatory name for firewall rule (uniqueness not enforced)
	// Firewall Rule name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// An optional set of Network Context Profiles. Can be
	// looked up using vcd_nsxt_network_context_profile data source.
	// A set of Network Context Profile IDs. Leaving it empty means 'Any'
	// +kubebuilder:validation:Optional
	NetworkContextProfileIds []*string `json:"networkContextProfileIds,omitempty" tf:"network_context_profile_ids,omitempty"`

	// - reverses value of source_ids for the rule to
	// match everything except specified IDs.
	// Reverses firewall matching for to match all except Source Groups specified in 'source_ids' (VCD 10.3.2+)
	// +kubebuilder:validation:Optional
	SourceGroupsExcluded *bool `json:"sourceGroupsExcluded,omitempty" tf:"source_groups_excluded,omitempty"`

	// A set of source object Firewall Groups (IP Sets or Security groups).
	// Leaving it empty matches Any (all)
	// A set of Source Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'
	// +crossplane:generate:reference:type=IPSet
	// +kubebuilder:validation:Optional
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// References to IPSet to populate sourceIds.
	// +kubebuilder:validation:Optional
	SourceIdsRefs []v1.Reference `json:"sourceIdsRefs,omitempty" tf:"-"`

	// Selector for a list of IPSet to populate sourceIds.
	// +kubebuilder:validation:Optional
	SourceIdsSelector *v1.Selector `json:"sourceIdsSelector,omitempty" tf:"-"`
}

// DistributedFirewallSpec defines the desired state of DistributedFirewall
type DistributedFirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DistributedFirewallParameters `json:"forProvider"`
}

// DistributedFirewallStatus defines the observed state of DistributedFirewall.
type DistributedFirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DistributedFirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DistributedFirewall is the Schema for the DistributedFirewalls API. The Distributed Firewall allows user to segment organization virtual data center entities, such as virtual machines, based on virtual machine names and attributes.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type DistributedFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DistributedFirewallSpec   `json:"spec"`
	Status            DistributedFirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributedFirewallList contains a list of DistributedFirewalls
type DistributedFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedFirewall `json:"items"`
}

// Repository type metadata.
var (
	DistributedFirewall_Kind             = "DistributedFirewall"
	DistributedFirewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DistributedFirewall_Kind}.String()
	DistributedFirewall_KindAPIVersion   = DistributedFirewall_Kind + "." + CRDGroupVersion.String()
	DistributedFirewall_GroupVersionKind = CRDGroupVersion.WithKind(DistributedFirewall_Kind)
)

func init() {
	SchemeBuilder.Register(&DistributedFirewall{}, &DistributedFirewallList{})
}
