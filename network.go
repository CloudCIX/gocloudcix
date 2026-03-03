// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix

import (
	"github.com/CloudCIX/gocloudcix/option"
)

// NetworkService contains methods and other services that help with interacting
// with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options []option.RequestOption
	// Management of Network Firewall Rules
	//
	// This module provides API endpoints for managing network firewall rules in the
	// CloudCIX Compute platform. Each project can have ONE project firewall and ONE
	// geo firewall. These firewalls control traffic flow to and from your cloud
	// resources by defining rules that allow or block network traffic.
	//
	// Firewall Types:
	//
	//  1. Project Firewall (type: "project") - Fine-grained rules controlling traffic
	//     based on specific source/destination IP addresses, ports, and protocols. Can
	//     reference your member's IP Address Groups using '@groupname' syntax in
	//     source/destination fields (e.g., "source": "@office_networks").
	//
	//  2. Geo Firewall (type: "geo") - Country-based IP filtering using global IP
	//     Address Groups (member_id = 0) that contain IP ranges for specific
	//     countries/regions. References groups by numeric ID using the
	//     "ip_address_group_id" field (e.g., "ip_address_group_id": 123).
	//
	// IP Address Group Usage:
	//
	//   - Project Firewalls: Use "source": "@groupname" or "destination": "@groupname"
	//     (member groups only)
	//   - Geo Firewalls: Use "ip_address_group_id": 123 (global groups with member_id=0
	//     only)
	//   - Project firewalls cannot use global groups, and geo firewalls cannot use
	//     member groups
	//
	// CRITICAL: When updating firewall rules, you MUST include ALL rules you want to
	// keep. The update operation replaces the entire rule list - any rules not
	// included in the update will be permanently deleted.
	//
	// Available operations:
	//
	//   - List and filter firewall rules across your projects by type
	//   - Create a project's single firewall or geo firewall with complete rule
	//     definitions
	//   - Retrieve detailed information about individual firewall configurations
	//   - Update firewall rules (replaces ALL existing rules) or delete firewalls by
	//     changing their state
	//
	// Rule Direction: Each firewall rule specifies traffic direction using the
	// 'inbound' flag:
	//
	//   - inbound: true = Inbound rule (traffic coming INTO your project/network)
	//   - inbound: false = Outbound rule (traffic going OUT FROM your project/network)
	//     Rules default to outbound (inbound: false) if not specified.
	//
	// Each firewall includes its associated project, router, rule definitions, and
	// current state.
	Firewalls NetworkFirewallService
	// Management of Network IP Groups
	//
	// IP address groups organise sets of CIDR networks for use in firewall rules and
	// access control. Two types are available:
	//
	// - Geo Groups (type="geo"): Maintained by CloudCIX and accessible to all members
	//
	//   - Used for geo-filtering based on country IP ranges (e.g., 'Ireland', 'USA',
	//     'China')
	//   - Essential for creating geo firewalls that block/allow traffic from specific
	//     countries
	//   - To list country groups: GET /ip_address_groups?search[member_id]=0
	//   - Referenced in geo firewall rules by numeric ID: "ip_address_group_id": 123
	//
	//   - Project Groups (type="project"): Created and managed by individual members for
	//     their own use
	//   - Used for project firewalls with fine-grained access control
	//   - Examples: office networks, VPN endpoints, admin workstations
	//   - Referenced in project firewall rules using @groupname syntax: "source":
	//     "@office_networks"
	//
	// Usage in Firewall Rules:
	//
	// - Project Firewall: "source": "@office_networks" (uses project type groups only)
	// - Geo Firewall: "group_name": "@ie_v4" (uses geo type groups only)
	//
	// Examples:
	//
	//   - Block traffic from Ireland: Create geo firewall rule with group_name of
	//     Ireland group
	//   - Allow access from office: Create project firewall rule with source
	//     "@office_networks"
	//   - Compliance geo-blocking: Use global country groups referenced by ID in geo
	//     firewall rules
	IPGroups NetworkIPGroupService
	// Management of Virtual Network Routers
	//
	// This module provides API endpoints for managing virtual network routers in the
	// CloudCIX Compute platform. Each project can have one virtual router that
	// provides network connectivity and routing between your cloud resources and
	// external networks and one or more static routes. The router manages one or more
	// private networks (subnets) and handles traffic routing, NAT, and network
	// isolation for your project's virtual machines and containers.
	//
	// Network Router Type:
	//
	//  1. Project Router (type: "router") - Manage IP forwarding, and participate in
	//     routing decisions within isolated network topologies.
	//  2. Static Routes (type: "static_route") - Define a fixed route entry within the
	//     Project Router's routing table. It maps a destination network to a nexthop
	//     IP, enabling deterministic packet forwarding without dynamic updates.
	//
	// Available operations:
	//
	//   - List and filter virtual routers from all your projects
	//   - Create a project's router with one or more private network definitions (RFC
	//     1918 address ranges)
	//   - Retrieve detailed router configuration including networks, IP addresses, and
	//     state
	//   - Update router by adding networks, changing network names, or changing router
	//     state
	//
	// Network Management: When creating or adding networks, you only specify the IPv4
	// CIDR range and name. The system automatically generates VLAN IDs and IPv6 ranges
	// based on regional availability. When updating a router to add new networks, you
	// must include all existing networks (with their auto-generated VLAN and IPv6
	// properties) to preserve them, plus any new networks (with only IPv4 CIDR and
	// name specified). Existing network IPv4/IPv6 ranges and VLANs cannot be modified,
	// but network names can be updated by including the name field with existing
	// networks.
	//
	// Each router includes its associated project, public IP addresses (IPv4/IPv6),
	// private networks with VLANs, and current state. You can add additional private
	// networks to an existing router through the update operation.
	Routers NetworkRouterService
	// Management of Network VPNs Supported Types are:
	//
	// - site-to-site
	Vpns NetworkVpnService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r NetworkService) {
	r = NetworkService{}
	r.Options = opts
	r.Firewalls = NewNetworkFirewallService(opts...)
	r.IPGroups = NewNetworkIPGroupService(opts...)
	r.Routers = NewNetworkRouterService(opts...)
	r.Vpns = NewNetworkVpnService(opts...)
	return
}
