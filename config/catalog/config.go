// Package catalog contains VCD catalog configuration.
package catalog

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_catalog
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/catalog
	p.AddResourceConfigurator("vcd_catalog", func(r *config.Resource) {
	})
}
