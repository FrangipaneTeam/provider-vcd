// Package network contains VCD network configuration.
package network

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_network_direct
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_direct
	p.AddResourceConfigurator("vcd_network_direct", func(r *config.Resource) {
		r.References["external_network"] = config.Reference{
			Type: "ExternalNetworkV2",
		}

	})

	// vcd_network_isolated
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_isolated
	p.AddResourceConfigurator("vcd_network_isolated", func(r *config.Resource) {
		r.References["owner_id"] = config.Reference{
			Type: tools.GenerateType("org", "Group"),
		}
	})

	// vcd_network_isolated_v2
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_isolated_v2
	p.AddResourceConfigurator("vcd_network_isolated_v2", func(r *config.Resource) {
	})

	// vcd_network_routed
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_routed
	p.AddResourceConfigurator("vcd_network_routed", func(r *config.Resource) {
	})

	// vcd_network_routed_v2
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_routed_v2
	p.AddResourceConfigurator("vcd_network_routed_v2", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: tools.GenerateType("nsxt", "EdgeGateway"),
		}
	})

	// vcd_external_network
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/external_network
	p.AddResourceConfigurator("vcd_external_network", func(r *config.Resource) {
	})

	// vcd_external_network_v2
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/external_network_v2
	p.AddResourceConfigurator("vcd_external_network_v2", func(r *config.Resource) {
	})
}
