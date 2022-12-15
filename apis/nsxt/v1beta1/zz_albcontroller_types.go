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

type ALBControllerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ALB Controller version (e.g. 20.1.3)
	// NSX-T ALB Controller version
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ALBControllerParameters struct {

	// An optional description NSX-T ALB Controller
	// NSX-T ALB Controller description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// License type of ALB Controller (ENTERPRISE or BASIC)
	// NSX-T ALB License type. One of 'BASIC', 'ENTERPRISE'. Must not be used from VCD 10.4.0 onwards
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// The password for ALB Controller. Password will not be refreshed.
	// NSX-T ALB Controller Password
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The URL of ALB Controller
	// NSX-T ALB Controller URL
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// The username for ALB Controller
	// NSX-T ALB Controller Username
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// ALBControllerSpec defines the desired state of ALBController
type ALBControllerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ALBControllerParameters `json:"forProvider"`
}

// ALBControllerStatus defines the observed state of ALBController.
type ALBControllerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ALBControllerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ALBController is the Schema for the ALBControllers API. Provides a resource to manage NSX-T ALB Controller for Providers. It helps to integrate VMware Cloud Director with NSX-T Advanced Load Balancer deployment. Controller instances are registered with VMware Cloud Director instance. Controller instances serve as a central control plane for the load-balancing services provided by NSX-T Advanced Load Balancer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type ALBController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ALBControllerSpec   `json:"spec"`
	Status            ALBControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ALBControllerList contains a list of ALBControllers
type ALBControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ALBController `json:"items"`
}

// Repository type metadata.
var (
	ALBController_Kind             = "ALBController"
	ALBController_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ALBController_Kind}.String()
	ALBController_KindAPIVersion   = ALBController_Kind + "." + CRDGroupVersion.String()
	ALBController_GroupVersionKind = CRDGroupVersion.WithKind(ALBController_Kind)
)

func init() {
	SchemeBuilder.Register(&ALBController{}, &ALBControllerList{})
}
