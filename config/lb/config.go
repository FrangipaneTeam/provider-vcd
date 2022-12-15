// Package lb contains VCD lb configuration.
package lb

import (
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// vcd_lb_app_profile
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_app_profile
	p.AddResourceConfigurator("vcd_lb_app_profile", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}
	})

	// vcd_lb_app_rule
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_app_rule
	p.AddResourceConfigurator("vcd_lb_app_rule", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}
	})

	// vcd_lb_server_pool
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_server_pool
	p.AddResourceConfigurator("vcd_lb_server_pool", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}
	})

	// vcd_lb_service_monitor
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_service_monitor
	p.AddResourceConfigurator("vcd_lb_service_monitor", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}
	})

	// vcd_lb_virtual_server
	// https://registry.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_virtual_server
	p.AddResourceConfigurator("vcd_lb_virtual_server", func(r *config.Resource) {
		r.References["edge_gateway"] = config.Reference{
			Type: tools.GenerateType("edgegateway", "EdgeGateway"),
		}

		r.References["app_profile_id"] = config.Reference{
			Type: "AppProfile",
		}

		r.References["server_pool_id"] = config.Reference{
			Type: "ServerPool",
		}

		r.References["app_rule_ids"] = config.Reference{
			Type: "AppRule",
		}

		r.References["monitor_id"] = config.Reference{
			Type: "ServiceMonitor",
		}
	})
}
