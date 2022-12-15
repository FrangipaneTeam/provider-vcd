// Package nsxv contains VCD nsxv configuration.
package nsxv

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_nsxv_dhcp_relay
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_dhcp_relay
	p.AddResourceConfigurator("vcd_nsxv_dhcp_relay", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}

		r.References["ip_sets"] = config.Reference{
			Type: "IPSet",
		}

		r.References["relay_agent.network_name"] = config.Reference{
			Type: tools.GenerateType("network", "Routed"),
		}
	})

	// vcd_nsxv_dnat
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_dnat
	p.AddResourceConfigurator("vcd_nsxv_dnat", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}
	})

	// vcd_nsxv_firewall_rule
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_firewall_rule
	p.AddResourceConfigurator("vcd_nsxv_firewall_rule", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}

		r.References["source.ipsets"] = config.Reference{
			Type: "IPSet",
		}
	})

	// vcd_nsxv_ip_set
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_ip_set
	p.AddResourceConfigurator("vcd_nsxv_ip_set", func(r *config.Resource) {

	})

	// vcd_nsxv_snat
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_snat
	p.AddResourceConfigurator("vcd_nsxv_snat", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}
	})

}
