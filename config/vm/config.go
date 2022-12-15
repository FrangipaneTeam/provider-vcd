// Package vm contains VCD vm configuration.
package vm

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_vm
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vm
	p.AddResourceConfigurator("vcd_vm", func(r *config.Resource) {
		r.References["vapp_template_id"] = config.Reference{
			Type: tools.GenerateType("catalog", "VappTemplate"),
		}
	})

	// vcd_vm_affinity_rule
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vm_affinity_rule
	p.AddResourceConfigurator("vcd_vm_affinity_rule", func(r *config.Resource) {
		// r.References["vm_ids"] = config.Reference{
		// 	Type: tools.GenerateType("vapp", "VM"),
		// }
	})

	// vcd_vm_internal_disk
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vm_internal_disk
	p.AddResourceConfigurator("vcd_vm_internal_disk", func(r *config.Resource) {
		// TODO: References
		// r.References["depends_on"] = config.Reference{
		// 	Type: tools.GenerateType("vapp", "VM"),
		// }
	})

	// vcd_vm_sizing_policy
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vm_sizing_policy
	p.AddResourceConfigurator("vcd_vm_sizing_policy", func(r *config.Resource) {
	})

	// vm_vm_placement_policy
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vm_placement_policy
	p.AddResourceConfigurator("vcd_vm_placement_policy", func(r *config.Resource) {
	})
}
