// Package config contains external name configurations for this provider.
/*
Copyright 2022 Upbound Inc.
*/
package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"vcd_catalog":                TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.name }}"),
	"vcd_catalog_access_control": config.TemplatedStringAsIdentifier("catalog_id", "{{ .parameters.org }}.{{ .externam_name }}"),
	"vcd_catalog_item":           TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.catalog }}.{{ .parameters.name }}"),
	"vcd_catalog_media":          TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.catalog }}.{{ .parameters.name }}"),
	"vcd_catalog_vapp_template":  TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.catalog }}.{{ .parameters.name }}"),

	"vcd_library_certificate": config.TemplatedStringAsIdentifier("alias", "{{ .parameters.org }}.{{ .external_name }}"),

	"vcd_edgegateway":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .external_name }}"),
	"vcd_edgegateway_settings": config.TemplatedStringAsIdentifier("edge_gateway_name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),
	"vcd_edgegateway_vpn":      config.NameAsIdentifier,

	// vcd_org
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org
	"vcd_org": config.NameAsIdentifier,

	// vcd_org_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_group
	"vcd_org_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .external_name }}"),

	// vcd_org_ldap
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_ldap
	"vcd_org_ldap": TemplatedStringAsIdentifierWithNoName("{{ .parameters.org_id }}"),

	// vcd_org_user
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_user
	"vcd_org_user": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .external_name }}"),

	// vcd_org_vdc
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc
	"vcd_org_vdc": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .external_name }}"),

	// vcd_org_vdc_access_control
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc_access_control
	"vcd_org_vdc_access_control": config.TemplatedStringAsIdentifier("vdc", "{{ .parameters.org }}.{{ .external_name }}"),

	// vcd_network_direct
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_direct
	"vcd_network_direct": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_network_isolated
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_isolated
	"vcd_network_isolated": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_network_isolated_v2
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_isolated_v2
	"vcd_network_isolated_v2": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_network_routed
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_routed
	"vcd_network_routed": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_network_routed_v2
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/network_routed_v2
	"vcd_network_routed_v2": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_external_network
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/external_network
	"vcd_external_network": config.NameAsIdentifier,

	// vcd_external_network_v2
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/external_network_v2
	"vcd_external_network_v2": config.NameAsIdentifier,

	// vcd_nsxt_edgegateway
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway
	"vcd_nsxt_edgegateway": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_alb_cloud
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_cloud
	"vcd_nsxt_alb_cloud": config.NameAsIdentifier,

	// vcd_nsxt_alb_controller
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller
	"vcd_nsxt_alb_controller": config.NameAsIdentifier,

	// vcd_nsxt_alb_edgegateway_service_engine_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_edgegateway_service_engine_group
	"vcd_nsxt_alb_edgegateway_service_engine_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_alb_settings
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings
	"vcd_nsxt_alb_settings": config.NameAsIdentifier,

	// vcd_nsxt_alb_pool
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_pool
	"vcd_nsxt_alb_pool": config.NameAsIdentifier,

	// vcd_nsxt_alb_service_engine_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group
	"vcd_nsxt_alb_service_engine_group": config.NameAsIdentifier,

	// vcd_nsxt_alb_virtual_service
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_virtual_service
	"vcd_nsxt_alb_virtual_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_app_port_profile
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_app_port_profile
	"vcd_nsxt_app_port_profile": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_distributed_firewall
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall
	"vcd_nsxt_distributed_firewall": TemplatedStringAsIdentifierWithNoName("{{ .parameters.org }}.{{ .parameters.vdc }}"),

	// vcd_nsxt_dynamic_security_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_dynamic_security_group
	"vcd_nsxt_dynamic_security_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_edgegateway_bgp_configuration
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_configuration
	"vcd_nsxt_edgegateway_bgp_configuration": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_edgegateway_bgp_ip_prefix_list
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list
	"vcd_nsxt_edgegateway_bgp_ip_prefix_list": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_edgegateway_bgp_neighbor
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_neighbor
	"vcd_nsxt_edgegateway_bgp_neighbor": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_firewall
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_firewall
	"vcd_nsxt_firewall": config.TemplatedStringAsIdentifier("edge_gateway", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_ip_set
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_ip_set
	"vcd_nsxt_ip_set": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_ipsec_vpn_tunnel
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_ipsec_vpn_tunnel
	"vcd_nsxt_ipsec_vpn_tunnel": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_nat_rule
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_nat_rule
	"vcd_nsxt_nat_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_nsxt_network_dhcp
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_network_dhcp
	"vcd_nsxt_network_dhcp": config.TemplatedStringAsIdentifier("org_network_id", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_network_imported
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_network_imported
	"vcd_nsxt_network_imported": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_route_advertisement
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_route_advertisement
	"vcd_nsxt_route_advertisement": config.TemplatedStringAsIdentifier("edge_gateway_id", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxt_security_group
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_security_group
	"vcd_nsxt_security_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .edge_gateway_id }}.{{ .external_name }}"),

	// vcd_global_role
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/global_role
	"vcd_global_role": config.NameAsIdentifier,

	// vcd_nsxv_dhcp_relay
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_dhcp_relay
	"vcd_nsxv_dhcp_relay": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .external_name }}"),

	// vcd_nsxv_dnat
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_dnat
	"vcd_nsxv_dnat": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .parameters.edge_gateway }}.{{ .external_name }}"),

	// vcd_nsxv_firewall_rule
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_firewall_rule
	"vcd_nsxv_firewall_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .parameters.edge_gateway }}.{{ .external_name }}"),

	// vcd_nsxv_snat
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_snat
	"vcd_nsxv_snat": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .parameters.edge_gateway }}.{{ .external_name }}"),

	// vcd_nsxv_ip_set
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_ip_set
	"vcd_nsxv_ip_set": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org }}.{{ .parameters.vdc }}.{{ .parameters.edge_gateway }}.{{ .external_name }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}
