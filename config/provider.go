/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/FrangipaneTeam/provider-vcd/config/catalog"
	"github.com/FrangipaneTeam/provider-vcd/config/edgegateway"
	"github.com/FrangipaneTeam/provider-vcd/config/global"
	"github.com/FrangipaneTeam/provider-vcd/config/library"
	"github.com/FrangipaneTeam/provider-vcd/config/network"
	"github.com/FrangipaneTeam/provider-vcd/config/nsxt"
	"github.com/FrangipaneTeam/provider-vcd/config/nsxv"
	"github.com/FrangipaneTeam/provider-vcd/config/vapp"
	"github.com/FrangipaneTeam/provider-vcd/config/vdc"
	"github.com/FrangipaneTeam/provider-vcd/config/vm"
	"github.com/FrangipaneTeam/provider-vcd/pkg/tools"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), tools.ResourcePrefix, tools.ModulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			defaultVersion(),
			GroupKindOverrides(),
			KindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		catalog.Configure,
		edgegateway.Configure,
		library.Configure,
		network.Configure,
		nsxt.Configure,
		global.Configure,
		nsxv.Configure,
		vm.Configure,
		vapp.Configure,
		vdc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
