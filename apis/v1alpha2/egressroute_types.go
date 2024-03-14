/*
Copyright 2024 F5, Inc.

*/

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// EgressRoute provides a way to route Egress traffic. When combined with a Gateway
// listener, it can be used to forward traffic on the port specified by the
// listener to a set of backends specified by the EgressRoute.
type EgressRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of EgressRoute.
	Spec EgressRouteSpec `json:"spec"`

	// Status defines the current state of EgressRoute.
	Status EgressRouteStatus `json:"status,omitempty"`
}

// EgressRouteSpec defines the desired state of EgressRoute.
type EgressRouteSpec struct {
	CommonRouteSpec `json:",inline"`

	DualStackEnabled     `json:"dualStackEnabled,omitempty"`
	SnatpoolIps          `json:"snatpoolIps,omitempty"`
	DnsNat46Enabled      `json:"dnsNat46Enabled,omitempty"`
	DnsNat46PoolIps      `json:"dnsNat46PoolIps,omitempty"`
	DnsNat46Ipv4Subnet   `json:"dnsNat46Ipv4Subnet,omitempty"`
	DnsNat46SorryIp      `json:"dnsNat46SorryIp,omitempty"`
	Nat64Enabled         `json:"nat64Enabled,omitempty"`
	Nat64Ipv6Subnet      `json:"nat64Ipv6Subnet,omitempty"`
	DebugLogEnabled      `json:"debugLogEnabled,omitempty"`
	DnsCacheName         `json:"dnsCacheName,omitempty"`
	MaxTmmReplicas       `json:"maxTmmReplicas,omitempty"`
	MaxReservedStaticIps `json:"maxReservedStaticIps,omitempty"`
	DnsRateLimit         `json:"dnsRateLimit,omitempty"`
	VlanList             `json:"vlanList,omitempty"`
	DisableListedVlans   `json:"disableListedVlans,omitempty"`
	MaxDNS46TTL          `json:"maxDNS46TTL,omitempty"`
}

// EgressRouteStatus defines the observed state of EgressRoute.
type EgressRouteStatus struct {
	RouteStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// EgressRouteList contains a list of EgressRoute
type EgressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EgressRoute `json:"items"`
}

// EXAMPLE comments and validation
// DualStackEnabled refers to Dual Stack Support. It must be a RFC 1123 label.
//
// This validation is based off of the corresponding Kubernetes validation:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/util/validation/validation.go#L187
//
// This is used for Namespace name validation here:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/api/validation/generic.go#L63
//
// Valid values include:
//
// * "true", "false"
//
// Invalid values include:
//
// * "example.com" - "." is an invalid character
//
// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=63

// DualStackEnabled refers to Dual Stack Support.
//
// Valid values include:
//
// * "true", "false"
type DualStackEnabled bool

// SnatpoolIps refers to Dual Stack Support. It must be a RFC 1123 label.
//
// This validation is based off of the corresponding Kubernetes validation:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/util/validation/validation.go#L187
//
// This is used for Namespace name validation here:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/api/validation/generic.go#L63
//
// Valid values include:
//
// * "mycoolsnatpoolname", "false"
type SnatpoolIps []string

// DnsNat46Enabled refers to DNS NAT ipv4 to ipv6 Enabled.
//
// Valid values include:
//
// * "true", "false"
type DnsNat46Enabled bool

// DnsNat46PoolIps refers to DNS NAT ipv4 to ipv6 Pool IPs.
//
// Valid values include:
//
// * {"10.1.1.1", "10.2.2.2"}
type DnsNat46PoolIps []string

// DnsNat46Ipv4Subnet refers to DNS NAT ipv4 to ipv6 Subnet.
// Default value: "10.2.2.0/24"
//
// Valid values include:
//
// * "10.2.2.0/24"
type DnsNat46Ipv4Subnet string

// DnsNat46SorryIp refers to DNS NAT ipv4 to ipv6 Sorry IP.
// Default value: "192.168.1.1"
//
// Valid values include:
//
// * "192.168.1.1"
type DnsNat46SorryIp string

// Nat64Enabled refers to NAT ipv6 to ipv4 Enabled.
//
// Valid values include:
//
// * "true", "false"
type Nat64Enabled bool

// Nat64Ipv6Subnet refers to NAT ipv6 Subnet.
// Default: "64:ff9b::/96"
//
// Valid values include:
//
// * "64:ff9b::/96"
type Nat64Ipv6Subnet string

// DebugLogEnabled Boolean for allowing user to enable/disable
// egress iRule debug logging
//
// Valid values include:
//
// * "true", "false"
type DebugLogEnabled bool

// DnsCacheName
//
// Valid values include:
//
// * "true", "false"
type DnsCacheName string

// MaxTmmReplicas The maximum number of allowed operational tmm pods.
// This should equal the number of SelfIps.
//
// Default: 1
//
// Valid values include:
// * 1
type MaxTmmReplicas int

// MaxReservedStaticIps The maximum IPv4 addresses reserved from the
// given CIDR Block "dnsNat46Ipv4Subnet" for static IPv4<-->IPv6
// mapping. These reserved IPs will be from the beginning of the
// CIDR Block. Rest of the IPs from the given CIDR Block will be
// distributed across TMM Pods based on the value in "maxTmmReplicas".
//
// Default: 0
//
// Valid values include:
// * 0
type MaxReservedStaticIps int

// DnsRateLimit Rate Limit is for DNS Virtual Server on a per pod
// basis. To disable the rate limit set value to 0.
//
// Default: 0
// Min: 0
// Max: 4294967295
// Valid values include:
// * 0 - 4294967295
type DnsRateLimit int

// VlanList String array to list Vlans to restrict to or disable with
// DisableListedVlan
//
// Default: 0
// Min: 0
// Max: 4294967295
// Valid values include:
// * 0 - 4294967295
type VlanList []string

// DisableListedVlans Boolean for allowing user to enable/disable
// egress iRule debug logging
//
// Default: true
// Valid values include:
//
// * "true", "false"
type DisableListedVlans bool

// MaxDNS46TTL The ceiling for DNS46 translation TTL,
// AAAA TTLs greater than max is set to max.
//
// Default: 120
// Min: 1
// Max: 2147483647
// Valid values include:
// * 0 - 2147483647
type MaxDNS46TTL int
