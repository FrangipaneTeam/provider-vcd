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
	edgegateway "github.com/FrangipaneTeam/provider-vcd/internal/controller/egw/edgegateway"
	settings "github.com/FrangipaneTeam/provider-vcd/internal/controller/egw/settings"
	vpn "github.com/FrangipaneTeam/provider-vcd/internal/controller/egw/vpn"
	certificate "github.com/FrangipaneTeam/provider-vcd/internal/controller/library/certificate"
	direct "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/direct"
	providerconfig "github.com/FrangipaneTeam/provider-vcd/internal/controller/providerconfig"
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
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
