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

type SharedWithObservation struct {

	// The name of the subject (group or user) which we are sharing with.
	// Name of the subject (group or user) with which we are sharing
	SubjectName *string `json:"subjectName,omitempty" tf:"subject_name,omitempty"`
}

type SharedWithParameters struct {

	// The access level for the user or group to which we are sharing. (Only ReadOnly is available)
	// The access level for the user or group to which we are sharing. (Only ReadOnly is available)
	// +kubebuilder:validation:Required
	AccessLevel *string `json:"accessLevel" tf:"access_level,omitempty"`

	// The ID of a group which we are sharing with. Required if user_id is not set.
	// ID of the group to which we are sharing. Required if user_id is not set
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The ID of a user which we are sharing with. Required if group_id is not set.
	// ID of the user to which we are sharing. Required if group_id is not set
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type VdcAccessControlObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// one or more blocks defining a subject to which we are sharing.
	// See shared_with below for detail. It cannot be used if shared_with_everyone is set.
	// +kubebuilder:validation:Optional
	SharedWith []SharedWithObservation `json:"sharedWith,omitempty" tf:"shared_with,omitempty"`
}

type VdcAccessControlParameters struct {

	// Access level when the VDC is shared with everyone (only ReadOnly is available). Required when shared_with_everyone is set.
	// Access level when the VDC is shared with everyone (only ReadOnly is available). Required when shared_with_everyone is set
	// +kubebuilder:validation:Optional
	EveryoneAccessLevel *string `json:"everyoneAccessLevel,omitempty" tf:"everyone_access_level,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations.
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// one or more blocks defining a subject to which we are sharing.
	// See shared_with below for detail. It cannot be used if shared_with_everyone is set.
	// +kubebuilder:validation:Optional
	SharedWith []SharedWithParameters `json:"sharedWith,omitempty" tf:"shared_with,omitempty"`

	// Whether the VDC is shared with everyone.
	// Whether the VDC is shared with everyone
	// +kubebuilder:validation:Required
	SharedWithEveryone *bool `json:"sharedWithEveryone" tf:"shared_with_everyone,omitempty"`
}

// VdcAccessControlSpec defines the desired state of VdcAccessControl
type VdcAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VdcAccessControlParameters `json:"forProvider"`
}

// VdcAccessControlStatus defines the observed state of VdcAccessControl.
type VdcAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VdcAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VdcAccessControl is the Schema for the VdcAccessControls API. Provides a VMware Cloud Director Org VDC access control resource. This can be used to share VDC across users or groups.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type VdcAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VdcAccessControlSpec   `json:"spec"`
	Status            VdcAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VdcAccessControlList contains a list of VdcAccessControls
type VdcAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VdcAccessControl `json:"items"`
}

// Repository type metadata.
var (
	VdcAccessControl_Kind             = "VdcAccessControl"
	VdcAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VdcAccessControl_Kind}.String()
	VdcAccessControl_KindAPIVersion   = VdcAccessControl_Kind + "." + CRDGroupVersion.String()
	VdcAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(VdcAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&VdcAccessControl{}, &VdcAccessControlList{})
}
