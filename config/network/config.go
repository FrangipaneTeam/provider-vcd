// Package network contains VCD network configuration.
package network

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_network_direct
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_direct
	p.AddResourceConfigurator("vcd_network_direct", func(r *config.Resource) {
		// TODO: external_network ?

	})

	// vcd_network_isolated
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_isolated
	p.AddResourceConfigurator("vcd_network_isolated", func(r *config.Resource) {
	})
}
