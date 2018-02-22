// Package dhcp represents the VPP binary API of the 'dhcp' VPP module.
// DO NOT EDIT. Generated from '/usr/share/vpp/api/dhcp.api.json'
package dhcp

import "git.fd.io/govpp.git/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0xbe9dca5b

// DhcpServer represents the VPP binary API data type 'dhcp_server'.
// Generated from '/usr/share/vpp/api/dhcp.api.json', line 65:
//
//                "dhcp_server",
//                16
//
type DhcpServer struct {
	ServerVrfID uint32
	DhcpServer  []byte `struc:"[16]byte"`
}

func (*DhcpServer) GetTypeName() string {
	return "dhcp_server"
}
func (*DhcpServer) GetCrcString() string {
	return "f16506c4"
}

// DhcpProxyConfig represents the VPP binary API message 'dhcp_proxy_config'.
// Generated from '/usr/share/vpp/api/dhcp.api.json', line 78:
//
//            "dhcp_proxy_config_reply",
//            [
//
type DhcpProxyConfig struct {
	RxVrfID        uint32
	ServerVrfID    uint32
	IsIpv6         uint8
	IsAdd          uint8
	DhcpServer     []byte `struc:"[16]byte"`
	DhcpSrcAddress []byte `struc:"[16]byte"`
}

func (*DhcpProxyConfig) GetMessageName() string {
	return "dhcp_proxy_config"
}
func (*DhcpProxyConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpProxyConfig) GetCrcString() string {
	return "6af4b645"
}
func NewDhcpProxyConfig() api.Message {
	return &DhcpProxyConfig{}
}

// DhcpProxyConfigReply represents the VPP binary API message 'dhcp_proxy_config_reply'.
//
type DhcpProxyConfigReply struct {
	Retval int32
}

func (*DhcpProxyConfigReply) GetMessageName() string {
	return "dhcp_proxy_config_reply"
}
func (*DhcpProxyConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpProxyConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func NewDhcpProxyConfigReply() api.Message {
	return &DhcpProxyConfigReply{}
}

// DhcpProxySetVss represents the VPP binary API message 'dhcp_proxy_set_vss'.
//
type DhcpProxySetVss struct {
	TblID      uint32
	VssType    uint8
	VpnASCIIID []byte `struc:"[129]byte"`
	Oui        uint32
	VpnIndex   uint32
	IsIpv6     uint8
	IsAdd      uint8
}

func (*DhcpProxySetVss) GetMessageName() string {
	return "dhcp_proxy_set_vss"
}
func (*DhcpProxySetVss) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpProxySetVss) GetCrcString() string {
	return "606535aa"
}
func NewDhcpProxySetVss() api.Message {
	return &DhcpProxySetVss{}
}

// DhcpProxySetVssReply represents the VPP binary API message 'dhcp_proxy_set_vss_reply'.
//
type DhcpProxySetVssReply struct {
	Retval int32
}

func (*DhcpProxySetVssReply) GetMessageName() string {
	return "dhcp_proxy_set_vss_reply"
}
func (*DhcpProxySetVssReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpProxySetVssReply) GetCrcString() string {
	return "e8d4e804"
}
func NewDhcpProxySetVssReply() api.Message {
	return &DhcpProxySetVssReply{}
}

// DhcpClientConfig represents the VPP binary API message 'dhcp_client_config'.
//
type DhcpClientConfig struct {
	SwIfIndex        uint32
	Hostname         []byte `struc:"[64]byte"`
	ClientID         []byte `struc:"[64]byte"`
	IsAdd            uint8
	WantDhcpEvent    uint8
	SetBroadcastFlag uint8
	Pid              uint32
}

func (*DhcpClientConfig) GetMessageName() string {
	return "dhcp_client_config"
}
func (*DhcpClientConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpClientConfig) GetCrcString() string {
	return "652052aa"
}
func NewDhcpClientConfig() api.Message {
	return &DhcpClientConfig{}
}

// DhcpClientConfigReply represents the VPP binary API message 'dhcp_client_config_reply'.
//
type DhcpClientConfigReply struct {
	Retval int32
}

func (*DhcpClientConfigReply) GetMessageName() string {
	return "dhcp_client_config_reply"
}
func (*DhcpClientConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpClientConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func NewDhcpClientConfigReply() api.Message {
	return &DhcpClientConfigReply{}
}

// DhcpComplEvent represents the VPP binary API message 'dhcp_compl_event'.
//
type DhcpComplEvent struct {
	Pid           uint32
	Hostname      []byte `struc:"[64]byte"`
	IsIpv6        uint8
	MaskWidth     uint8
	HostAddress   []byte `struc:"[16]byte"`
	RouterAddress []byte `struc:"[16]byte"`
	HostMac       []byte `struc:"[6]byte"`
}

func (*DhcpComplEvent) GetMessageName() string {
	return "dhcp_compl_event"
}
func (*DhcpComplEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}
func (*DhcpComplEvent) GetCrcString() string {
	return "28c60f01"
}
func NewDhcpComplEvent() api.Message {
	return &DhcpComplEvent{}
}

// DhcpProxyDump represents the VPP binary API message 'dhcp_proxy_dump'.
//
type DhcpProxyDump struct {
	IsIP6 uint8
}

func (*DhcpProxyDump) GetMessageName() string {
	return "dhcp_proxy_dump"
}
func (*DhcpProxyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*DhcpProxyDump) GetCrcString() string {
	return "6fe91190"
}
func NewDhcpProxyDump() api.Message {
	return &DhcpProxyDump{}
}

// DhcpProxyDetails represents the VPP binary API message 'dhcp_proxy_details'.
//
type DhcpProxyDetails struct {
	RxVrfID        uint32
	VssOui         uint32
	VssFibID       uint32
	VssType        uint8
	VssVpnASCIIID  []byte `struc:"[129]byte"`
	IsIpv6         uint8
	DhcpSrcAddress []byte `struc:"[16]byte"`
	Count          uint8  `struc:"sizeof=Servers"`
	Servers        []DhcpServer
}

func (*DhcpProxyDetails) GetMessageName() string {
	return "dhcp_proxy_details"
}
func (*DhcpProxyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*DhcpProxyDetails) GetCrcString() string {
	return "a5f2ad84"
}
func NewDhcpProxyDetails() api.Message {
	return &DhcpProxyDetails{}
}
