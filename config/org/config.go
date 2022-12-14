// Package org contains VCD organization configuration.
package org

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// vcd_org_vdc
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc
	p.AddResourceConfigurator("vcd_org_vdc", func(r *config.Resource) {

		// default_compute_policy_id - (Optional, v3.8+, VCD 10.2+) ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy.
		r.References["default_compute_policy_id"] = config.Reference{
			// TODO
		}

		// vm_sizing_policy_ids -  (Optional, v3.0+, VCD 10.2+) Set of IDs of VM Sizing policies that are assigned to this VDC. This field requires default_compute_policy_id to be configured together.
		r.References["vm_sizing_policy_ids"] = config.Reference{
			// TODO
		}

		// vm_placement_policy_ids - (Optional, v3.8+, VCD 10.2+) Set of IDs of VM Placement policies that are assigned to this VDC. This field requires default_compute_policy_id to be configured together.
		r.References["vm_placement_policy_ids"] = config.Reference{
			// TODO
		}

		// edge_cluster_id - (Optional, v3.8+, VCD 10.3+) An ID of NSX-T Edge Cluster which should provide vApp Networking Services or DHCP for isolated networks. Can be looked up using vcd_nsxt_edge_cluster data source.
		r.References["edge_cluster_id"] = config.Reference{
			// TODO
		}

		// network_pool_name - Reference to a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
		r.References["network_pool_name"] = config.Reference{
			// TODO
		}
	})

	// vcd_org_vdc_access_control
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc_access_control
	p.AddResourceConfigurator("vcd_org_vdc_access_control", func(r *config.Resource) {
		r.References["vdc"] = config.Reference{
			Type: "Vdc",
		}

		// user_id - (Optional) The ID of a user which we are sharing with. Required if group_id is not set.
		r.References["user_id"] = config.Reference{
			// TODO
		}

		// group_id - (Optional) The ID of a group which we are sharing with. Required if user_id is not set.
		r.References["group_id"] = config.Reference{
			// TODO
		}
	})

}
