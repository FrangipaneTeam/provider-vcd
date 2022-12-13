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
		r.ShortGroup = "catalog"
	})

	// vcd_catalog_access_control
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/catalog_access_control
	p.AddResourceConfigurator("vcd_catalog_access_control", func(r *config.Resource) {
		// catalog_id
		r.References["catalog_id"] = config.Reference{
			Type: "Catalog",
		}
	})

	// vcd_catalog_item
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/catalog_item
	p.AddResourceConfigurator("vcd_catalog_item", func(r *config.Resource) {

	})
}
