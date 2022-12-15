// Package vapp contains VCD vapp configuration.
package vapp

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_vapp
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp
	p.AddResourceConfigurator("vcd_vapp", func(r *config.Resource) {
	})

	// vcd_vapp_access_control
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_access_control
	p.AddResourceConfigurator("vcd_vapp_access_control", func(r *config.Resource) {
		r.References["vapp_id"] = config.Reference{
			Type: "Vapp",
		}
	})

	// vcd_vapp_firewall_rules
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules
	p.AddResourceConfigurator("vcd_vapp_firewall_rules", func(r *config.Resource) {
		r.References["vapp_id"] = config.Reference{
			Type: "Vapp",
		}

		r.References["vapp_name"] = config.Reference{
			Type: "Vapp",
		}

		r.References["network_id"] = config.Reference{
			Type: "VappNetwork",
		}
	})

	// vcd_vapp_nat_rules
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_nat_rules
	p.AddResourceConfigurator("vcd_vapp_nat_rules", func(r *config.Resource) {
		r.References["vapp_id"] = config.Reference{
			Type: "Vapp",
		}

		r.References["vapp_name"] = config.Reference{
			Type: "Vapp",
		}

		r.References["network_id"] = config.Reference{
			Type: "VappNetwork",
		}

		r.References["rule.vm_id"] = config.Reference{
			Type: tools.GenerateType("vm", "VM"),
		}
	})

	// vcd_vapp_network
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network
	p.AddResourceConfigurator("vcd_vapp_network", func(r *config.Resource) {
	})

	// vcd_vapp_org_network
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_org_network
	p.AddResourceConfigurator("vcd_vapp_org_network", func(r *config.Resource) {
	})

	// vcd_vapp_static_routing
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_static_routing
	p.AddResourceConfigurator("vcd_vapp_static_routing", func(r *config.Resource) {
		r.References["vapp_id"] = config.Reference{
			Type: "Vapp",
		}

		r.References["vapp_name"] = config.Reference{
			Type: "Vapp",
		}

		r.References["network_id"] = config.Reference{
			Type: "VappNetwork",
		}
	})

	// vcd_vapp_vm
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_vm
	p.AddResourceConfigurator("vcd_vapp_vm", func(r *config.Resource) {
		r.References["vapp_name"] = config.Reference{
			Type: "Vapp",
		}

		r.References["vapp_template_id"] = config.Reference{
			Type: tools.GenerateType("catalog", "VappTemplate"),
		}

		// r.References["network.name"] = config.Reference{
		// 	Type: "OrgNetwork",
		// }
	})

}
