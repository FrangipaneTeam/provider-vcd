// Package global contains VCD global configuration.
package global

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_global_role
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/global_role
	p.AddResourceConfigurator("vcd_global_role", func(r *config.Resource) {
	})
}
