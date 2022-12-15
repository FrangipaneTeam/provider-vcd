// Package vdc contains VCD vdc configuration.
package vdc

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_vdc_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/vdc_group
	p.AddResourceConfigurator("vcd_vdc_group", func(r *config.Resource) {
		r.References["starting_vdc_id"] = config.Reference{
			Type: tools.GenerateType("org", "Vdc"),
		}
		r.References["participating_vdc_ids"] = config.Reference{
			Type: tools.GenerateType("org", "Vdc"),
		}
	})
}
