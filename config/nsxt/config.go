// Package nsxt contains VCD nsxt configuration.
package nsxt

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_nsxt_edgegateway
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway
	p.AddResourceConfigurator("vcd_nsxt_edgegateway", func(r *config.Resource) {
		r.References["owner_id"] = config.Reference{
			Type: tools.GenerateType("org", "Group"),
		}
		// r.References["external_network_id"] = config.Reference{
		// 	Type: tools.GenerateType("network", "ExternalNetworkV2"),
		// }
	})

	// vcd_nsxt_alb_cloud
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_cloud
	p.AddResourceConfigurator("vcd_nsxt_alb_cloud", func(r *config.Resource) {
		r.References["controller_id"] = config.Reference{
			Type: "ALBController",
		}
	})

	// vcd_nsxt_alb_controller
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller
	p.AddResourceConfigurator("vcd_nsxt_alb_controller", func(r *config.Resource) {
	})

	// vcd_nsxt_alb_edgegateway_service_engine_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_edgegateway_service_engine_group
	p.AddResourceConfigurator("vcd_nsxt_alb_edgegateway_service_engine_group", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
		r.References["service_engine_group_id"] = config.Reference{
			Type: "ALBServiceEngineGroup",
		}
	})

	// vcd_nsxt_alb_settings
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings
	p.AddResourceConfigurator("vcd_nsxt_alb_settings", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_alb_pool
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_pool
	p.AddResourceConfigurator("vcd_nsxt_alb_pool", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_alb_service_engine_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group
	p.AddResourceConfigurator("vcd_nsxt_alb_service_engine_group", func(r *config.Resource) {
		r.References["alb_cloud_id"] = config.Reference{
			Type: "ALBCloud",
		}
	})

	// vcd_nsxt_alb_virtual_service
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_virtual_service
	p.AddResourceConfigurator("vcd_nsxt_alb_virtual_service", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
		r.References["pool_id"] = config.Reference{
			Type: "ALBPool",
		}
		r.References["service_engine_group_id"] = config.Reference{
			Type: "ALBEdgeGatewayServiceEngineGroup",
		}
		r.References["virtual_ip_address"] = config.Reference{
			Type: "EdgeGateway",
		}
		r.References["ca_certificate_id"] = config.Reference{
			Type: tools.GenerateType("library", "Certificate"),
		}
	})

	// vcd_nsxt_app_port_profile
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_app_port_profile
	p.AddResourceConfigurator("vcd_nsxt_app_port_profile", func(r *config.Resource) {

	})

	// vcd_nsxt_distributed_firewall
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall
	p.AddResourceConfigurator("vcd_nsxt_distributed_firewall", func(r *config.Resource) {
		r.References["vdc_group_id"] = config.Reference{
			Type: tools.GenerateType("org", "Group"),
		}
		r.References["rule.app_port_profile_ids"] = config.Reference{
			Type: "AppPortProfile",
		}
		r.References["rule.source_ids"] = config.Reference{
			Type: "IPSet",
		}
	})

	// vcd_nsxt_dynamic_security_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_dynamic_security_group
	p.AddResourceConfigurator("vcd_nsxt_dynamic_security_group", func(r *config.Resource) {
		r.References["vdc_group_id"] = config.Reference{
			Type: tools.GenerateType("org", "Group"),
		}
	})

	// vcd_nsxt_edgegateway_bgp_configuration
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_configuration
	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_configuration", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_edgegateway_bgp_ip_prefix_list
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list
	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_ip_prefix_list", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_edgegateway_bgp_neighbor
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_neighbor
	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_neighbor", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}

		r.References["in_filter_ip_prefix_list_id"] = config.Reference{
			Type: "EdgeGatewayBGPIPPrefixList",
		}

		r.References["out_filter_ip_prefix_list_id"] = config.Reference{
			Type: "EdgeGatewayBGPIPPrefixList",
		}

	})

	// vcd_nsxt_firewall
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_firewall
	p.AddResourceConfigurator("vcd_nsxt_firewall", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_ip_set
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_ip_set
	p.AddResourceConfigurator("vcd_nsxt_ip_set", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_ipsec_vpn_tunnel
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_ipsec_vpn_tunnel
	p.AddResourceConfigurator("vcd_nsxt_ipsec_vpn_tunnel", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
		r.References["local_ip_address"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_nat_rule
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_nat_rule
	p.AddResourceConfigurator("vcd_nsxt_nat_rule", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
		r.References["external_address"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_network_dhcp
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_network_dhcp
	p.AddResourceConfigurator("vcd_nsxt_network_dhcp", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}

		r.References["org_network_id"] = config.Reference{
			Type: tools.GenerateType("network", "RoutedV2"),
		}
	})

	// vcd_nsxt_network_imported
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_network_imported
	p.AddResourceConfigurator("vcd_nsxt_network_imported", func(r *config.Resource) {
		r.References["owner_id"] = config.Reference{
			Type: tools.GenerateType("org", "Group"),
		}
	})

	// vcd_nsxt_route_advertisement
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_route_advertisement
	p.AddResourceConfigurator("vcd_nsxt_route_advertisement", func(r *config.Resource) {
		r.References["owner_id"] = config.Reference{
			Type: tools.GenerateType("org", "Group"),
		}

		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}
	})

	// vcd_nsxt_security_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_security_group
	p.AddResourceConfigurator("vcd_nsxt_security_group", func(r *config.Resource) {
		r.References["edge_gateway_id"] = config.Reference{
			Type: "EdgeGateway",
		}

		// r.References["member_org_network_ids"] = config.Reference{
		// 	Type: tools.GenerateType("network", "RoutedV2"),
		// }

	})
}
