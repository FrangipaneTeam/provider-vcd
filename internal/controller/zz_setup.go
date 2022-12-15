/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accesscontrol "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/accesscontrol"
	catalog "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/catalog"
	item "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/item"
	media "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/media"
	vapptemplate "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/vapptemplate"
	edgegateway "github.com/FrangipaneTeam/provider-vcd/internal/controller/edgegateway/edgegateway"
	settings "github.com/FrangipaneTeam/provider-vcd/internal/controller/edgegateway/settings"
	vpn "github.com/FrangipaneTeam/provider-vcd/internal/controller/edgegateway/vpn"
	certificate "github.com/FrangipaneTeam/provider-vcd/internal/controller/library/certificate"
	direct "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/direct"
	externalnetwork "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/externalnetwork"
	externalnetworkv2 "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/externalnetworkv2"
	isolated "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/isolated"
	isolatedv2 "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/isolatedv2"
	routed "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/routed"
	routedv2 "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/routedv2"
	edgegatewaynsxt "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/edgegateway"
	group "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/group"
	ldap "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/ldap"
	user "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/user"
	vdc "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/vdc"
	vdcaccesscontrol "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/vdcaccesscontrol"
	providerconfig "github.com/FrangipaneTeam/provider-vcd/internal/controller/providerconfig"
	org "github.com/FrangipaneTeam/provider-vcd/internal/controller/vcd/org"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontrol.Setup,
		catalog.Setup,
		item.Setup,
		media.Setup,
		vapptemplate.Setup,
		edgegateway.Setup,
		settings.Setup,
		vpn.Setup,
		certificate.Setup,
		direct.Setup,
		externalnetwork.Setup,
		externalnetworkv2.Setup,
		isolated.Setup,
		isolatedv2.Setup,
		routed.Setup,
		routedv2.Setup,
		edgegatewaynsxt.Setup,
		group.Setup,
		ldap.Setup,
		user.Setup,
		vdc.Setup,
		vdcaccesscontrol.Setup,
		providerconfig.Setup,
		org.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
