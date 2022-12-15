/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/FrangipaneTeam/provider-vcd/apis/catalog/v1beta1"
	v1beta1edgegateway "github.com/FrangipaneTeam/provider-vcd/apis/edgegateway/v1beta1"
	v1beta1library "github.com/FrangipaneTeam/provider-vcd/apis/library/v1beta1"
	v1beta1network "github.com/FrangipaneTeam/provider-vcd/apis/network/v1beta1"
	v1beta1nsxt "github.com/FrangipaneTeam/provider-vcd/apis/nsxt/v1beta1"
	v1beta1org "github.com/FrangipaneTeam/provider-vcd/apis/org/v1beta1"
	v1alpha1 "github.com/FrangipaneTeam/provider-vcd/apis/v1alpha1"
	v1beta1apis "github.com/FrangipaneTeam/provider-vcd/apis/v1beta1"
	v1beta1vcd "github.com/FrangipaneTeam/provider-vcd/apis/vcd/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1edgegateway.SchemeBuilder.AddToScheme,
		v1beta1library.SchemeBuilder.AddToScheme,
		v1beta1network.SchemeBuilder.AddToScheme,
		v1beta1nsxt.SchemeBuilder.AddToScheme,
		v1beta1org.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
		v1beta1vcd.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
