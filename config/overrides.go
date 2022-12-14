package config

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() config.ResourceOption { //nolint:gocyclo
	return func(r *config.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch k {
			// org is a organization name.
			case "org":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("org", "Org"),
				}
			// vdc is a virtual data center name.
			case "vdc":
				r.References[k] = config.Reference{
					Type: tools.GenerateType("vdc", "Vdc"),
				}
			}
		}
	}
}
