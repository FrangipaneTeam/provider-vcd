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

type InsertedMediaObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InsertedMediaParameters struct {

	// The name of the catalog where to find media file
	// catalog name where to find media file
	// +kubebuilder:validation:Required
	Catalog *string `json:"catalog" tf:"catalog,omitempty"`

	// Allows to pass answer to question in vCD
	// "The guest operating system has locked the CD-ROM door and is probably using the CD-ROM.
	// Disconnect anyway (and override the lock)?"
	// when ejecting from a VM which is powered on. True means "Yes" as answer to question. Default is true
	// When ejecting answers automatically to question yes
	// +kubebuilder:validation:Optional
	EjectForce *bool `json:"ejectForce,omitempty" tf:"eject_force,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// - The name of VM to be used to insert media file
	// VM in vApp in which media will be inserted or ejected
	// +kubebuilder:validation:Required
	VMName *string `json:"vmName" tf:"vm_name,omitempty"`

	// - The name of vApp to find
	// vApp to use
	// +kubebuilder:validation:Required
	VappName *string `json:"vappName" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// InsertedMediaSpec defines the desired state of InsertedMedia
type InsertedMediaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InsertedMediaParameters `json:"forProvider"`
}

// InsertedMediaStatus defines the observed state of InsertedMedia.
type InsertedMediaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InsertedMediaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InsertedMedia is the Schema for the InsertedMedias API. Provides a VMware Cloud Director resource for inserting or ejecting media (ISO) file for the VM. Create this resource for inserting the media, and destroy it for ejecting.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type InsertedMedia struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InsertedMediaSpec   `json:"spec"`
	Status            InsertedMediaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InsertedMediaList contains a list of InsertedMedias
type InsertedMediaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InsertedMedia `json:"items"`
}

// Repository type metadata.
var (
	InsertedMedia_Kind             = "InsertedMedia"
	InsertedMedia_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InsertedMedia_Kind}.String()
	InsertedMedia_KindAPIVersion   = InsertedMedia_Kind + "." + CRDGroupVersion.String()
	InsertedMedia_GroupVersionKind = CRDGroupVersion.WithKind(InsertedMedia_Kind)
)

func init() {
	SchemeBuilder.Register(&InsertedMedia{}, &InsertedMediaList{})
}
