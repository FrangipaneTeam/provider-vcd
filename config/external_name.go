// Package config contains external name configurations for this provider.
/*
Copyright 2022 Upbound Inc.
*/
package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"vcd_catalog":                TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.name }}"),
	"vcd_catalog_access_control": config.TemplatedStringAsIdentifier("catalog_id", "{{ .parameters.org }}.{{ .externam_name }}"),
	"vcd_catalog_item":           TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.catalog }}.{{ .parameters.name }}"),
	"vcd_catalog_media":          TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.catalog }}.{{ .parameters.name }}"),
	"vcd_catalog_vapp_template":  TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.catalog }}.{{ .parameters.name }}"),

	"vcd_library_certificate": config.TemplatedStringAsIdentifier("alias", "{{ .parameters.org }}.{{ .external_name }}"),

	"vcd_edgegateway":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .external_name }}"),
	"vcd_edgegateway_settings": config.TemplatedStringAsIdentifier("edge_gateway_name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}
