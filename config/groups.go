package config

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/upjet/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "vcd_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
// It's written with data from TF Provider AWS repo service grouping in here:
// https://github.com/hashicorp/terraform-provider-aws/tree/main/internal/service
//
// At the end, all of them are based on grouping of the AWS Go SDK.
// The initial grouping is calculated based on folder grouping of AWS TF Provider
// which is based on Go SDK. Here is the script used to fetch that list:
// https://gist.github.com/muvaf/8d61365ffc1df7757864422ba16d7819
// ! Warning do not consider _v[0-9]
// ? Examples
// ? "flexibleengine_vpc_subnet_v1": ReplaceGroupWords("vpc", 0), => Output: Group: vpc, Kind: Subnet
// ? "flexibleengine_networking_subnet_v2": ReplaceGroupWords("vpc", 0), => Output: Group: vpc, Kind: NetWorkingSubnet
var GroupMap = map[string]GroupKindCalculator{
	// EGW
	"vcd_edgegateway": ReplaceGroupWords("edgegateway", 0), // Group: egw, Kind: EdgeGateway
	// "vcd_edgegateway_settings": ReplaceGroupWords("egw", 1), // Group: egw, Kind: Settings
	// "vcd_edgegateway_vpn":      ReplaceGroupWords("egw", 1), // Group: egw, Kind: Vpn

	"vcd_external_network":    ReplaceGroupWords("network", 0), // Group: network, Kind: ExternalNetwork
	"vcd_external_network_v2": ReplaceGroupWords("network", 0), // Group: network, Kind: ExternalNetworkV2

}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	// EGW
	"vcd_edgegateway": "EdgeGateway",

	// ORG
	"vcd_org": "Org",

	// NSXT
	"vcd_nsxt_edgegateway": "EdgeGateway",
}
