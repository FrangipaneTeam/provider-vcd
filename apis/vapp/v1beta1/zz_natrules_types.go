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

type NATRulesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configures a NAT rule; see Rules below for details.
	// +kubebuilder:validation:Optional
	Rule []NATRulesRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type NATRulesParameters struct {

	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is false.
	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.
	// +kubebuilder:validation:Optional
	EnableIPMasquerade *bool `json:"enableIpMasquerade,omitempty" tf:"enable_ip_masquerade,omitempty"`

	// Enable or disable NAT. Default is true. To enable the NAT service, vcd_vapp_firewall_rules needs to be enabled as well.
	// Enable or disable NAT service. Default is `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// "One of: ipTranslation (use IP translation), portForwarding (use port forwarding). For ipTranslation fields vm_id, vm_nic_id, mapping_mode are required and external_ip is optional. For portForwarding fields vm_id, vm_nic_id, protocol, external_port and forward_to_port are required.
	// One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).
	// +kubebuilder:validation:Required
	NATType *string `json:"natType" tf:"nat_type,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Configures a NAT rule; see Rules below for details.
	// +kubebuilder:validation:Optional
	Rule []NATRulesRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The identifier of vApp.
	// vApp identifier
	// +crossplane:generate:reference:type=Vapp
	// +kubebuilder:validation:Optional
	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	// Reference to a Vapp to populate vappId.
	// +kubebuilder:validation:Optional
	VappIDRef *v1.Reference `json:"vappIdRef,omitempty" tf:"-"`

	// Selector for a Vapp to populate vappId.
	// +kubebuilder:validation:Optional
	VappIDSelector *v1.Selector `json:"vappIdSelector,omitempty" tf:"-"`

	// The name of VDC to use, optional if defined at provider level.
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NATRulesRuleObservation struct {

	// ID of the rule. Can be used to track syslog messages.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NATRulesRuleParameters struct {

	// External IP address to forward to or External IP address to map to VM.
	// External IP address to forward to or External IP address to map to VM
	// +kubebuilder:validation:Optional
	ExternalIP *string `json:"externalIp,omitempty" tf:"external_ip,omitempty"`

	// External port to forward. -1 value for any port.
	// External port to forward.
	// +kubebuilder:validation:Optional
	ExternalPort *float64 `json:"externalPort,omitempty" tf:"external_port,omitempty"`

	// Internal port to forward. -1 value for any port.
	// Internal port to forward.
	// +kubebuilder:validation:Optional
	ForwardToPort *float64 `json:"forwardToPort,omitempty" tf:"forward_to_port,omitempty"`

	// Mapping mode. One of: automatic, manual.
	// Mapping mode. One of: `automatic`, `manual`
	// +kubebuilder:validation:Optional
	MappingMode *string `json:"mappingMode,omitempty" tf:"mapping_mode,omitempty"`

	// Protocol to forward. One of: TCP (forward TCP packets), UDP (forward UDP packets), TCP_UDP (forward TCP and UDP packets).
	// Protocol to forward. One of: `TCP` (forward TCP packets), `UDP` (forward UDP packets), `TCP_UDP` (forward TCP and UDP packets).
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// VM to which this rule applies.
	// VM to which this rule applies.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-vcd/apis/vm/v1beta1.VM
	// +kubebuilder:validation:Optional
	VMID *string `json:"vmId,omitempty" tf:"vm_id,omitempty"`

	// Reference to a VM in vm to populate vmId.
	// +kubebuilder:validation:Optional
	VMIDRef *v1.Reference `json:"vmIdRef,omitempty" tf:"-"`

	// Selector for a VM in vm to populate vmId.
	// +kubebuilder:validation:Optional
	VMIDSelector *v1.Selector `json:"vmIdSelector,omitempty" tf:"-"`

	// VM NIC ID to which this rule applies.
	// VM NIC ID to which this rule applies.
	// +kubebuilder:validation:Required
	VMNicID *float64 `json:"vmNicId" tf:"vm_nic_id,omitempty"`
}

// NATRulesSpec defines the desired state of NATRules
type NATRulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATRulesParameters `json:"forProvider"`
}

// NATRulesStatus defines the observed state of NATRules.
type NATRulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATRulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATRules is the Schema for the NATRuless API. Provides a VMware Cloud Director vApp NAT resource. This can be used to create, modify, and delete NAT rules.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NATRules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATRulesSpec   `json:"spec"`
	Status            NATRulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATRulesList contains a list of NATRuless
type NATRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATRules `json:"items"`
}

// Repository type metadata.
var (
	NATRules_Kind             = "NATRules"
	NATRules_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATRules_Kind}.String()
	NATRules_KindAPIVersion   = NATRules_Kind + "." + CRDGroupVersion.String()
	NATRules_GroupVersionKind = CRDGroupVersion.WithKind(NATRules_Kind)
)

func init() {
	SchemeBuilder.Register(&NATRules{}, &NATRulesList{})
}
