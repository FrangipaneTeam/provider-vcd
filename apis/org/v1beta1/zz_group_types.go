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

type GroupObservation struct {

	// The ID of the Organization group
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Read only) The set of user names that belong to this group. It's only populated if the users
	// are created after the group (with depends_on the given group).
	// Read only. Set of user names that belong to the group
	UserNames []*string `json:"userNames,omitempty" tf:"user_names,omitempty"`
}

type GroupParameters struct {

	// The description of Organization group
	// Description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Identity provider type for this this group. One of SAML or
	// INTEGRATED. Note LDAP must be configured to create INTEGRATED groups and names must
	// match LDAP group names. If LDAP is not configured - it will return 403 errors.
	// Identity provider type - 'SAML' or 'INTEGRATED' for LDAP
	// +kubebuilder:validation:Required
	ProviderType *string `json:"providerType" tf:"provider_type,omitempty"`

	// The role of the group. Role names can be retrieved from the organization. Both built-in roles and
	// custom built can be used. The roles normally available are:
	// Existing role name to assign
	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. Provides a VMware Cloud Director Organization group. This can be used to create, update, and delete organization groups defined in SAML or LDAP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
