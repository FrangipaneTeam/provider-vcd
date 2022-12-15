// Package nsxt contains VCD nsxt configuration.
package nsxt

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_nsxt_edgegateway
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway
	p.AddResourceConfigurator("vcd_nsxt_edgegateway", func(r *config.Resource) {
		// TODO: external_network ?
	})
}
