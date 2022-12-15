// Package library contains VCD library configuration.
package library

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_certificate_library
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/certificate_library
	p.AddResourceConfigurator("vcd_library_certificate", func(r *config.Resource) {
	})
}
