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
	subscribedcatalog "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/subscribedcatalog"
	vapptemplate "github.com/FrangipaneTeam/provider-vcd/internal/controller/catalog/vapptemplate"
	edgegateway "github.com/FrangipaneTeam/provider-vcd/internal/controller/edgegateway/edgegateway"
	settings "github.com/FrangipaneTeam/provider-vcd/internal/controller/edgegateway/settings"
	vpn "github.com/FrangipaneTeam/provider-vcd/internal/controller/edgegateway/vpn"
	role "github.com/FrangipaneTeam/provider-vcd/internal/controller/global/role"
	appprofile "github.com/FrangipaneTeam/provider-vcd/internal/controller/lb/appprofile"
	apprule "github.com/FrangipaneTeam/provider-vcd/internal/controller/lb/apprule"
	serverpool "github.com/FrangipaneTeam/provider-vcd/internal/controller/lb/serverpool"
	servicemonitor "github.com/FrangipaneTeam/provider-vcd/internal/controller/lb/servicemonitor"
	virtualserver "github.com/FrangipaneTeam/provider-vcd/internal/controller/lb/virtualserver"
	direct "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/direct"
	externalnetwork "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/externalnetwork"
	externalnetworkv2 "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/externalnetworkv2"
	isolated "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/isolated"
	isolatedv2 "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/isolatedv2"
	routed "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/routed"
	routedv2 "github.com/FrangipaneTeam/provider-vcd/internal/controller/network/routedv2"
	albcloud "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albcloud"
	albcontroller "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albcontroller"
	albedgegatewayserviceenginegroup "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albedgegatewayserviceenginegroup"
	albpool "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albpool"
	albserviceenginegroup "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albserviceenginegroup"
	albsettings "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albsettings"
	albvirtualservice "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/albvirtualservice"
	appportprofile "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/appportprofile"
	distributedfirewall "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/distributedfirewall"
	dynamicsecuritygroup "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/dynamicsecuritygroup"
	edgegatewaynsxt "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/edgegateway"
	edgegatewaybgpconfiguration "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/edgegatewaybgpconfiguration"
	edgegatewaybgpipprefixlist "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/edgegatewaybgpipprefixlist"
	edgegatewaybgpneighbor "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/edgegatewaybgpneighbor"
	firewall "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/firewall"
	ipsecvpntunnel "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/ipsecvpntunnel"
	ipset "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/ipset"
	natrule "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/natrule"
	networkdhcp "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/networkdhcp"
	networkimported "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/networkimported"
	routeadvertisement "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/routeadvertisement"
	securitygroup "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxt/securitygroup"
	dhcprelay "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxv/dhcprelay"
	dnat "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxv/dnat"
	firewallrule "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxv/firewallrule"
	ipsetnsxv "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxv/ipset"
	snat "github.com/FrangipaneTeam/provider-vcd/internal/controller/nsxv/snat"
	group "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/group"
	ldap "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/ldap"
	user "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/user"
	vdc "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/vdc"
	vdcaccesscontrol "github.com/FrangipaneTeam/provider-vcd/internal/controller/org/vdcaccesscontrol"
	providerconfig "github.com/FrangipaneTeam/provider-vcd/internal/controller/providerconfig"
	bundle "github.com/FrangipaneTeam/provider-vcd/internal/controller/rights/bundle"
	rolerights "github.com/FrangipaneTeam/provider-vcd/internal/controller/rights/role"
	accesscontrolvapp "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/accesscontrol"
	firewallrules "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/firewallrules"
	independentdisk "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/independentdisk"
	natrules "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/natrules"
	network "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/network"
	orgnetwork "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/orgnetwork"
	staticrouting "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/staticrouting"
	vapp "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/vapp"
	vm "github.com/FrangipaneTeam/provider-vcd/internal/controller/vapp/vm"
	org "github.com/FrangipaneTeam/provider-vcd/internal/controller/vcd/org"
	groupvdc "github.com/FrangipaneTeam/provider-vcd/internal/controller/vdc/group"
	affinityrule "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/affinityrule"
	insertedmedia "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/insertedmedia"
	internaldisk "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/internaldisk"
	placementpolicy "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/placementpolicy"
	securitytag "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/securitytag"
	sizingpolicy "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/sizingpolicy"
	vmvm "github.com/FrangipaneTeam/provider-vcd/internal/controller/vm/vm"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontrol.Setup,
		catalog.Setup,
		item.Setup,
		media.Setup,
		subscribedcatalog.Setup,
		vapptemplate.Setup,
		edgegateway.Setup,
		settings.Setup,
		vpn.Setup,
		role.Setup,
		appprofile.Setup,
		apprule.Setup,
		serverpool.Setup,
		servicemonitor.Setup,
		virtualserver.Setup,
		direct.Setup,
		externalnetwork.Setup,
		externalnetworkv2.Setup,
		isolated.Setup,
		isolatedv2.Setup,
		routed.Setup,
		routedv2.Setup,
		albcloud.Setup,
		albcontroller.Setup,
		albedgegatewayserviceenginegroup.Setup,
		albpool.Setup,
		albserviceenginegroup.Setup,
		albsettings.Setup,
		albvirtualservice.Setup,
		appportprofile.Setup,
		distributedfirewall.Setup,
		dynamicsecuritygroup.Setup,
		edgegatewaynsxt.Setup,
		edgegatewaybgpconfiguration.Setup,
		edgegatewaybgpipprefixlist.Setup,
		edgegatewaybgpneighbor.Setup,
		firewall.Setup,
		ipsecvpntunnel.Setup,
		ipset.Setup,
		natrule.Setup,
		networkdhcp.Setup,
		networkimported.Setup,
		routeadvertisement.Setup,
		securitygroup.Setup,
		dhcprelay.Setup,
		dnat.Setup,
		firewallrule.Setup,
		ipsetnsxv.Setup,
		snat.Setup,
		group.Setup,
		ldap.Setup,
		user.Setup,
		vdc.Setup,
		vdcaccesscontrol.Setup,
		providerconfig.Setup,
		bundle.Setup,
		rolerights.Setup,
		accesscontrolvapp.Setup,
		firewallrules.Setup,
		independentdisk.Setup,
		natrules.Setup,
		network.Setup,
		orgnetwork.Setup,
		staticrouting.Setup,
		vapp.Setup,
		vm.Setup,
		org.Setup,
		groupvdc.Setup,
		affinityrule.Setup,
		insertedmedia.Setup,
		internaldisk.Setup,
		placementpolicy.Setup,
		securitytag.Setup,
		sizingpolicy.Setup,
		vmvm.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
