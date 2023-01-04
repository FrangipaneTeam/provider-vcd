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

	// vcd_catalog_media
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/catalog_media
	p.AddResourceConfigurator("vcd_catalog_media", func(r *config.Resource) {
		// catalog_id
		r.References["catalog"] = config.Reference{
			Type: "Catalog",
		}
		// temporary workaround for remove show_upload_progress
		// need to update the terraform provider documentation
		r.ExternalName.OmittedFields = []string{"show_upload_progress"}
	})

	// vcd_catalog_vapp_template
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/catalog_vapp_template
	p.AddResourceConfigurator("vcd_catalog_vapp_template", func(r *config.Resource) {
		// catalog_id
		// r.UseAsync = true
		r.References["catalog_id"] = config.Reference{
			Type: "Catalog",
		}

		// temporary workaround for remove show_upload_progress
		// need to update the terraform provider documentation
		r.ExternalName.OmittedFields = []string{"show_upload_progress"}
	})

	// vcd_subscribed_catalog
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/subscribed_catalog
	p.AddResourceConfigurator("vcd_subscribed_catalog", func(r *config.Resource) {
	})
}
