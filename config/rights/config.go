// Package rights contains VCD rights configuration.
package rights

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_rights_bundle
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/rights_bundle
	p.AddResourceConfigurator("vcd_rights_bundle", func(r *config.Resource) {
	})

	// vcd_role
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/role
	p.AddResourceConfigurator("vcd_role", func(r *config.Resource) {
	})

}
