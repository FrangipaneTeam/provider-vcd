// Package egw contains VCD edgegateway configuration.
package egw

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_edgegateway
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/edgegateway
	p.AddResourceConfigurator("vcd_edgegateway", func(r *config.Resource) {
	})

	// vcd_edgegateway_settings
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/edgegateway_settings
	p.AddResourceConfigurator("vcd_edgegateway_settings", func(r *config.Resource) {
	})
}
