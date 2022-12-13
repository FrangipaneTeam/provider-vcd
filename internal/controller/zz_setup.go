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
	providerconfig "github.com/FrangipaneTeam/provider-vcd/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontrol.Setup,
		catalog.Setup,
		item.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
