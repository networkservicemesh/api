// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectioncontext.proto

package networkservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IpFamily_Family int32

const (
	IpFamily_IPV4 IpFamily_Family = 0
	IpFamily_IPV6 IpFamily_Family = 1
)

var IpFamily_Family_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}

var IpFamily_Family_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x IpFamily_Family) String() string {
	return proto.EnumName(IpFamily_Family_name, int32(x))
}

func (IpFamily_Family) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{2, 0}
}

type IpNeighbor struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	HardwareAddress      string   `protobuf:"bytes,2,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpNeighbor) Reset()         { *m = IpNeighbor{} }
func (m *IpNeighbor) String() string { return proto.CompactTextString(m) }
func (*IpNeighbor) ProtoMessage()    {}
func (*IpNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{0}
}

func (m *IpNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpNeighbor.Unmarshal(m, b)
}
func (m *IpNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpNeighbor.Marshal(b, m, deterministic)
}
func (m *IpNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpNeighbor.Merge(m, src)
}
func (m *IpNeighbor) XXX_Size() int {
	return xxx_messageInfo_IpNeighbor.Size(m)
}
func (m *IpNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IpNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IpNeighbor proto.InternalMessageInfo

func (m *IpNeighbor) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *IpNeighbor) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

type Route struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type IpFamily struct {
	Family               IpFamily_Family `protobuf:"varint,1,opt,name=family,proto3,enum=connectioncontext.IpFamily_Family" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IpFamily) Reset()         { *m = IpFamily{} }
func (m *IpFamily) String() string { return proto.CompactTextString(m) }
func (*IpFamily) ProtoMessage()    {}
func (*IpFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{2}
}

func (m *IpFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpFamily.Unmarshal(m, b)
}
func (m *IpFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpFamily.Marshal(b, m, deterministic)
}
func (m *IpFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpFamily.Merge(m, src)
}
func (m *IpFamily) XXX_Size() int {
	return xxx_messageInfo_IpFamily.Size(m)
}
func (m *IpFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_IpFamily.DiscardUnknown(m)
}

var xxx_messageInfo_IpFamily proto.InternalMessageInfo

func (m *IpFamily) GetFamily() IpFamily_Family {
	if m != nil {
		return m.Family
	}
	return IpFamily_IPV4
}

type ExtraPrefixRequest struct {
	AddrFamily           *IpFamily `protobuf:"bytes,1,opt,name=addr_family,json=addrFamily,proto3" json:"addr_family,omitempty"`
	PrefixLen            uint32    `protobuf:"varint,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	RequiredNumber       uint32    `protobuf:"varint,3,opt,name=required_number,json=requiredNumber,proto3" json:"required_number,omitempty"`
	RequestedNumber      uint32    `protobuf:"varint,4,opt,name=requested_number,json=requestedNumber,proto3" json:"requested_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExtraPrefixRequest) Reset()         { *m = ExtraPrefixRequest{} }
func (m *ExtraPrefixRequest) String() string { return proto.CompactTextString(m) }
func (*ExtraPrefixRequest) ProtoMessage()    {}
func (*ExtraPrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{3}
}

func (m *ExtraPrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraPrefixRequest.Unmarshal(m, b)
}
func (m *ExtraPrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraPrefixRequest.Marshal(b, m, deterministic)
}
func (m *ExtraPrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraPrefixRequest.Merge(m, src)
}
func (m *ExtraPrefixRequest) XXX_Size() int {
	return xxx_messageInfo_ExtraPrefixRequest.Size(m)
}
func (m *ExtraPrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraPrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraPrefixRequest proto.InternalMessageInfo

func (m *ExtraPrefixRequest) GetAddrFamily() *IpFamily {
	if m != nil {
		return m.AddrFamily
	}
	return nil
}

func (m *ExtraPrefixRequest) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequiredNumber() uint32 {
	if m != nil {
		return m.RequiredNumber
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequestedNumber() uint32 {
	if m != nil {
		return m.RequestedNumber
	}
	return 0
}

type IPContext struct {
	SrcIpAddr            string                `protobuf:"bytes,1,opt,name=src_ip_addr,json=srcIpAddr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr            string                `protobuf:"bytes,2,opt,name=dst_ip_addr,json=dstIpAddr,proto3" json:"dst_ip_addr,omitempty"`
	SrcIpRequired        bool                  `protobuf:"varint,3,opt,name=src_ip_required,json=srcIpRequired,proto3" json:"src_ip_required,omitempty"`
	DstIpRequired        bool                  `protobuf:"varint,4,opt,name=dst_ip_required,json=dstIpRequired,proto3" json:"dst_ip_required,omitempty"`
	SrcRoutes            []*Route              `protobuf:"bytes,5,rep,name=src_routes,json=srcRoutes,proto3" json:"src_routes,omitempty"`
	DstRoutes            []*Route              `protobuf:"bytes,6,rep,name=dst_routes,json=dstRoutes,proto3" json:"dst_routes,omitempty"`
	ExcludedPrefixes     []string              `protobuf:"bytes,7,rep,name=excluded_prefixes,json=excludedPrefixes,proto3" json:"excluded_prefixes,omitempty"`
	IpNeighbors          []*IpNeighbor         `protobuf:"bytes,8,rep,name=ip_neighbors,json=ipNeighbors,proto3" json:"ip_neighbors,omitempty"`
	ExtraPrefixRequest   []*ExtraPrefixRequest `protobuf:"bytes,9,rep,name=extra_prefix_request,json=extraPrefixRequest,proto3" json:"extra_prefix_request,omitempty"`
	ExtraPrefixes        []string              `protobuf:"bytes,10,rep,name=extra_prefixes,json=extraPrefixes,proto3" json:"extra_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IPContext) Reset()         { *m = IPContext{} }
func (m *IPContext) String() string { return proto.CompactTextString(m) }
func (*IPContext) ProtoMessage()    {}
func (*IPContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{4}
}

func (m *IPContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPContext.Unmarshal(m, b)
}
func (m *IPContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPContext.Marshal(b, m, deterministic)
}
func (m *IPContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPContext.Merge(m, src)
}
func (m *IPContext) XXX_Size() int {
	return xxx_messageInfo_IPContext.Size(m)
}
func (m *IPContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IPContext.DiscardUnknown(m)
}

var xxx_messageInfo_IPContext proto.InternalMessageInfo

func (m *IPContext) GetSrcIpAddr() string {
	if m != nil {
		return m.SrcIpAddr
	}
	return ""
}

func (m *IPContext) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *IPContext) GetSrcIpRequired() bool {
	if m != nil {
		return m.SrcIpRequired
	}
	return false
}

func (m *IPContext) GetDstIpRequired() bool {
	if m != nil {
		return m.DstIpRequired
	}
	return false
}

func (m *IPContext) GetSrcRoutes() []*Route {
	if m != nil {
		return m.SrcRoutes
	}
	return nil
}

func (m *IPContext) GetDstRoutes() []*Route {
	if m != nil {
		return m.DstRoutes
	}
	return nil
}

func (m *IPContext) GetExcludedPrefixes() []string {
	if m != nil {
		return m.ExcludedPrefixes
	}
	return nil
}

func (m *IPContext) GetIpNeighbors() []*IpNeighbor {
	if m != nil {
		return m.IpNeighbors
	}
	return nil
}

func (m *IPContext) GetExtraPrefixRequest() []*ExtraPrefixRequest {
	if m != nil {
		return m.ExtraPrefixRequest
	}
	return nil
}

func (m *IPContext) GetExtraPrefixes() []string {
	if m != nil {
		return m.ExtraPrefixes
	}
	return nil
}

type DNSConfig struct {
	// ips of DNS Servers for this DNSConfig.  Any given IP may be IPv4 or IPv6
	DnsServerIps []string `protobuf:"bytes,1,rep,name=dns_server_ips,json=dnsServerIps,proto3" json:"dns_server_ips,omitempty"`
	// domains for which this DNSConfig provides resolution.  If empty, all domains.
	SearchDomains        []string `protobuf:"bytes,2,rep,name=search_domains,json=searchDomains,proto3" json:"search_domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSConfig) Reset()         { *m = DNSConfig{} }
func (m *DNSConfig) String() string { return proto.CompactTextString(m) }
func (*DNSConfig) ProtoMessage()    {}
func (*DNSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{5}
}

func (m *DNSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSConfig.Unmarshal(m, b)
}
func (m *DNSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSConfig.Marshal(b, m, deterministic)
}
func (m *DNSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSConfig.Merge(m, src)
}
func (m *DNSConfig) XXX_Size() int {
	return xxx_messageInfo_DNSConfig.Size(m)
}
func (m *DNSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DNSConfig proto.InternalMessageInfo

func (m *DNSConfig) GetDnsServerIps() []string {
	if m != nil {
		return m.DnsServerIps
	}
	return nil
}

func (m *DNSConfig) GetSearchDomains() []string {
	if m != nil {
		return m.SearchDomains
	}
	return nil
}

type DNSContext struct {
	Configs              []*DNSConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DNSContext) Reset()         { *m = DNSContext{} }
func (m *DNSContext) String() string { return proto.CompactTextString(m) }
func (*DNSContext) ProtoMessage()    {}
func (*DNSContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{6}
}

func (m *DNSContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSContext.Unmarshal(m, b)
}
func (m *DNSContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSContext.Marshal(b, m, deterministic)
}
func (m *DNSContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSContext.Merge(m, src)
}
func (m *DNSContext) XXX_Size() int {
	return xxx_messageInfo_DNSContext.Size(m)
}
func (m *DNSContext) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSContext.DiscardUnknown(m)
}

var xxx_messageInfo_DNSContext proto.InternalMessageInfo

func (m *DNSContext) GetConfigs() []*DNSConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

type EthernetContext struct {
	SrcMac               string   `protobuf:"bytes,1,opt,name=src_mac,json=srcMac,proto3" json:"src_mac,omitempty"`
	DstMac               string   `protobuf:"bytes,2,opt,name=dst_mac,json=dstMac,proto3" json:"dst_mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthernetContext) Reset()         { *m = EthernetContext{} }
func (m *EthernetContext) String() string { return proto.CompactTextString(m) }
func (*EthernetContext) ProtoMessage()    {}
func (*EthernetContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{7}
}

func (m *EthernetContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthernetContext.Unmarshal(m, b)
}
func (m *EthernetContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthernetContext.Marshal(b, m, deterministic)
}
func (m *EthernetContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthernetContext.Merge(m, src)
}
func (m *EthernetContext) XXX_Size() int {
	return xxx_messageInfo_EthernetContext.Size(m)
}
func (m *EthernetContext) XXX_DiscardUnknown() {
	xxx_messageInfo_EthernetContext.DiscardUnknown(m)
}

var xxx_messageInfo_EthernetContext proto.InternalMessageInfo

func (m *EthernetContext) GetSrcMac() string {
	if m != nil {
		return m.SrcMac
	}
	return ""
}

func (m *EthernetContext) GetDstMac() string {
	if m != nil {
		return m.DstMac
	}
	return ""
}

type SRIOVContext struct {
	EndpointVfPciAddress string   `protobuf:"bytes,1,opt,name=endpoint_vf_pci_address,json=endpointVfPciAddress,proto3" json:"endpoint_vf_pci_address,omitempty"`
	ClientPfPciAddress   string   `protobuf:"bytes,2,opt,name=client_pf_pci_address,json=clientPfPciAddress,proto3" json:"client_pf_pci_address,omitempty"`
	ClientVfPciAddress   string   `protobuf:"bytes,3,opt,name=client_vf_pci_address,json=clientVfPciAddress,proto3" json:"client_vf_pci_address,omitempty"`
	IommuGroup           uint32   `protobuf:"varint,4,opt,name=iommu_group,json=iommuGroup,proto3" json:"iommu_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SRIOVContext) Reset()         { *m = SRIOVContext{} }
func (m *SRIOVContext) String() string { return proto.CompactTextString(m) }
func (*SRIOVContext) ProtoMessage()    {}
func (*SRIOVContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{8}
}

func (m *SRIOVContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SRIOVContext.Unmarshal(m, b)
}
func (m *SRIOVContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SRIOVContext.Marshal(b, m, deterministic)
}
func (m *SRIOVContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SRIOVContext.Merge(m, src)
}
func (m *SRIOVContext) XXX_Size() int {
	return xxx_messageInfo_SRIOVContext.Size(m)
}
func (m *SRIOVContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SRIOVContext.DiscardUnknown(m)
}

var xxx_messageInfo_SRIOVContext proto.InternalMessageInfo

func (m *SRIOVContext) GetEndpointVfPciAddress() string {
	if m != nil {
		return m.EndpointVfPciAddress
	}
	return ""
}

func (m *SRIOVContext) GetClientPfPciAddress() string {
	if m != nil {
		return m.ClientPfPciAddress
	}
	return ""
}

func (m *SRIOVContext) GetClientVfPciAddress() string {
	if m != nil {
		return m.ClientVfPciAddress
	}
	return ""
}

func (m *SRIOVContext) GetIommuGroup() uint32 {
	if m != nil {
		return m.IommuGroup
	}
	return 0
}

type ConnectionContext struct {
	IpContext            *IPContext        `protobuf:"bytes,1,opt,name=ip_context,json=ipContext,proto3" json:"ip_context,omitempty"`
	DnsContext           *DNSContext       `protobuf:"bytes,2,opt,name=dns_context,json=dnsContext,proto3" json:"dns_context,omitempty"`
	EthernetContext      *EthernetContext  `protobuf:"bytes,3,opt,name=ethernet_context,json=ethernetContext,proto3" json:"ethernet_context,omitempty"`
	SriovContext         *SRIOVContext     `protobuf:"bytes,4,opt,name=sriov_context,json=sriovContext,proto3" json:"sriov_context,omitempty"`
	ExtraContext         map[string]string `protobuf:"bytes,5,rep,name=extra_context,json=extraContext,proto3" json:"extra_context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConnectionContext) Reset()         { *m = ConnectionContext{} }
func (m *ConnectionContext) String() string { return proto.CompactTextString(m) }
func (*ConnectionContext) ProtoMessage()    {}
func (*ConnectionContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{9}
}

func (m *ConnectionContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionContext.Unmarshal(m, b)
}
func (m *ConnectionContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionContext.Marshal(b, m, deterministic)
}
func (m *ConnectionContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionContext.Merge(m, src)
}
func (m *ConnectionContext) XXX_Size() int {
	return xxx_messageInfo_ConnectionContext.Size(m)
}
func (m *ConnectionContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionContext.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionContext proto.InternalMessageInfo

func (m *ConnectionContext) GetIpContext() *IPContext {
	if m != nil {
		return m.IpContext
	}
	return nil
}

func (m *ConnectionContext) GetDnsContext() *DNSContext {
	if m != nil {
		return m.DnsContext
	}
	return nil
}

func (m *ConnectionContext) GetEthernetContext() *EthernetContext {
	if m != nil {
		return m.EthernetContext
	}
	return nil
}

func (m *ConnectionContext) GetSriovContext() *SRIOVContext {
	if m != nil {
		return m.SriovContext
	}
	return nil
}

func (m *ConnectionContext) GetExtraContext() map[string]string {
	if m != nil {
		return m.ExtraContext
	}
	return nil
}

func init() {
	proto.RegisterEnum("connectioncontext.IpFamily_Family", IpFamily_Family_name, IpFamily_Family_value)
	proto.RegisterType((*IpNeighbor)(nil), "connectioncontext.IpNeighbor")
	proto.RegisterType((*Route)(nil), "connectioncontext.Route")
	proto.RegisterType((*IpFamily)(nil), "connectioncontext.IpFamily")
	proto.RegisterType((*ExtraPrefixRequest)(nil), "connectioncontext.ExtraPrefixRequest")
	proto.RegisterType((*IPContext)(nil), "connectioncontext.IPContext")
	proto.RegisterType((*DNSConfig)(nil), "connectioncontext.DNSConfig")
	proto.RegisterType((*DNSContext)(nil), "connectioncontext.DNSContext")
	proto.RegisterType((*EthernetContext)(nil), "connectioncontext.EthernetContext")
	proto.RegisterType((*SRIOVContext)(nil), "connectioncontext.SRIOVContext")
	proto.RegisterType((*ConnectionContext)(nil), "connectioncontext.ConnectionContext")
	proto.RegisterMapType((map[string]string)(nil), "connectioncontext.ConnectionContext.ExtraContextEntry")
}

func init() { proto.RegisterFile("connectioncontext.proto", fileDescriptor_c30b3f1555e8b686) }

var fileDescriptor_c30b3f1555e8b686 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x4f, 0x2b, 0x45,
	0x14, 0xb6, 0x2d, 0x14, 0xf6, 0xb4, 0xb4, 0x65, 0x83, 0xb2, 0xd1, 0x8b, 0x90, 0x8d, 0x57, 0x31,
	0x26, 0x10, 0x51, 0xf1, 0x06, 0x8d, 0xbf, 0x00, 0x6f, 0x9a, 0x08, 0x36, 0x43, 0x82, 0x46, 0x1f,
	0x36, 0xcb, 0xec, 0x69, 0x3b, 0xa1, 0x9d, 0x9d, 0x3b, 0x33, 0xcb, 0x85, 0xbf, 0xcf, 0xf8, 0xe6,
	0x9f, 0xe4, 0x83, 0x99, 0x1f, 0xbb, 0x16, 0xba, 0x7a, 0x9f, 0x3a, 0xfb, 0x9d, 0xef, 0xfb, 0xce,
	0xcc, 0x39, 0x67, 0xa6, 0xb0, 0x4d, 0x73, 0xce, 0x91, 0x6a, 0x96, 0x73, 0x9a, 0x73, 0x8d, 0xf7,
	0xfa, 0x40, 0xc8, 0x5c, 0xe7, 0xe1, 0xe6, 0x52, 0x20, 0x7e, 0x09, 0x30, 0x14, 0x97, 0xc8, 0x26,
	0xd3, 0x9b, 0x5c, 0x86, 0x3d, 0x68, 0x32, 0x11, 0x35, 0xf6, 0x1a, 0xfb, 0x01, 0x69, 0x32, 0x11,
	0x7e, 0x0c, 0x83, 0x69, 0x2a, 0xb3, 0xd7, 0xa9, 0xc4, 0x24, 0xcd, 0x32, 0x89, 0x4a, 0x45, 0x4d,
	0x1b, 0xed, 0x97, 0xf8, 0xf7, 0x0e, 0x8e, 0x77, 0x61, 0x95, 0xe4, 0x85, 0xc6, 0xf0, 0x1d, 0x68,
	0x0b, 0x89, 0x63, 0x76, 0xef, 0x7d, 0xfc, 0x57, 0x9c, 0xc1, 0xfa, 0x50, 0xfc, 0x98, 0xce, 0xd9,
	0xec, 0x21, 0x3c, 0x81, 0xf6, 0xd8, 0xae, 0x2c, 0xa7, 0x77, 0x14, 0x1f, 0x2c, 0x6f, 0xb9, 0x24,
	0x1f, 0xb8, 0x1f, 0xe2, 0x15, 0xf1, 0x33, 0x68, 0x7b, 0x97, 0x75, 0x58, 0x19, 0x8e, 0xae, 0x3f,
	0x1f, 0xbc, 0xe5, 0x57, 0xc7, 0x83, 0x46, 0xfc, 0x67, 0x03, 0xc2, 0xf3, 0x7b, 0x2d, 0xd3, 0x91,
	0xcd, 0x4a, 0xf0, 0x55, 0x81, 0x4a, 0x87, 0x5f, 0x43, 0xc7, 0xec, 0x3f, 0x59, 0xc8, 0xda, 0x39,
	0x7a, 0xef, 0x7f, 0xb2, 0x12, 0x30, 0x7c, 0x9f, 0x68, 0x07, 0xc0, 0x1d, 0x22, 0x99, 0x21, 0xb7,
	0x05, 0xd8, 0x20, 0x81, 0x43, 0x7e, 0x42, 0x1e, 0x7e, 0x04, 0x7d, 0x89, 0xaf, 0x0a, 0x26, 0x31,
	0x4b, 0x78, 0x31, 0xbf, 0x41, 0x19, 0xb5, 0x2c, 0xa7, 0x57, 0xc2, 0x97, 0x16, 0x35, 0xe5, 0x94,
	0x6e, 0x43, 0xff, 0x32, 0x57, 0x2c, 0xb3, 0x5f, 0xe1, 0x8e, 0x1a, 0xff, 0xdd, 0x82, 0x60, 0x38,
	0x3a, 0x75, 0xbb, 0x0a, 0xdf, 0x87, 0x8e, 0x92, 0x34, 0x61, 0xc2, 0x76, 0xc1, 0x17, 0x36, 0x50,
	0x92, 0x0e, 0x85, 0xa9, 0xbf, 0x89, 0x67, 0x4a, 0x57, 0x71, 0xd7, 0xa2, 0x20, 0x53, 0xda, 0xc7,
	0x3f, 0x84, 0xbe, 0xd7, 0x97, 0x3b, 0xb2, 0x3b, 0x5c, 0x27, 0x1b, 0xd6, 0x83, 0x78, 0xd0, 0xf0,
	0xbc, 0x4f, 0xc5, 0x5b, 0x71, 0x3c, 0xeb, 0x55, 0xf1, 0xbe, 0x04, 0x30, 0x7e, 0xd2, 0x34, 0x5c,
	0x45, 0xab, 0x7b, 0xad, 0xfd, 0xce, 0x51, 0x54, 0x53, 0x4d, 0x3b, 0x11, 0x76, 0xa3, 0x76, 0xa5,
	0x8c, 0xd0, 0x24, 0xf0, 0xc2, 0xf6, 0x9b, 0x84, 0x99, 0xd2, 0x5e, 0xf8, 0x09, 0x6c, 0xe2, 0x3d,
	0x9d, 0x15, 0x19, 0x66, 0x89, 0xab, 0x3c, 0xaa, 0x68, 0x6d, 0xaf, 0xb5, 0x1f, 0x90, 0x41, 0x19,
	0x18, 0x79, 0x3c, 0xfc, 0x0e, 0xba, 0x4c, 0x24, 0xdc, 0x4f, 0xb5, 0x8a, 0xd6, 0x6d, 0x9e, 0x9d,
	0xda, 0x76, 0x97, 0xb3, 0x4f, 0x3a, 0xac, 0x5a, 0xab, 0xf0, 0x17, 0xd8, 0x42, 0x33, 0x45, 0x3e,
	0x57, 0xe2, 0xdb, 0x13, 0x05, 0xd6, 0xe9, 0x79, 0x8d, 0xd3, 0xf2, 0xd0, 0x91, 0x10, 0x97, 0x07,
	0xf1, 0x39, 0xf4, 0x16, 0x8d, 0x51, 0x45, 0x60, 0x0f, 0xb1, 0xb1, 0xc0, 0x45, 0x15, 0xff, 0x0a,
	0xc1, 0xd9, 0xe5, 0xd5, 0x69, 0xce, 0xc7, 0x6c, 0x12, 0x7e, 0x00, 0xbd, 0x8c, 0xab, 0x44, 0xa1,
	0xbc, 0x43, 0x99, 0x30, 0xa1, 0xa2, 0x86, 0xd5, 0x74, 0x33, 0xae, 0xae, 0x2c, 0x38, 0x14, 0xca,
	0x38, 0x2b, 0x4c, 0x25, 0x9d, 0x26, 0x59, 0x3e, 0x4f, 0x19, 0x37, 0x37, 0xd5, 0x3a, 0x3b, 0xf4,
	0xcc, 0x81, 0xf1, 0x19, 0x80, 0x73, 0xb6, 0x83, 0x75, 0x0c, 0x6b, 0xd4, 0x26, 0x71, 0x9e, 0x9d,
	0xa3, 0x67, 0x35, 0x47, 0xab, 0x76, 0x42, 0x4a, 0x72, 0x7c, 0x0a, 0xfd, 0x73, 0x3d, 0x45, 0xc9,
	0x51, 0x97, 0x56, 0xdb, 0xb0, 0x66, 0x66, 0x62, 0x9e, 0xd2, 0xf2, 0xe2, 0x2b, 0x49, 0x2f, 0x52,
	0x6a, 0x02, 0xa6, 0xe7, 0x26, 0xe0, 0x06, 0xb3, 0x9d, 0x29, 0x7d, 0x91, 0xd2, 0xf8, 0xaf, 0x06,
	0x74, 0xaf, 0xc8, 0xf0, 0xe7, 0xeb, 0xd2, 0xe2, 0x0b, 0xd8, 0x46, 0x9e, 0x89, 0x9c, 0x71, 0x9d,
	0xdc, 0x8d, 0x13, 0x41, 0x59, 0xf5, 0xea, 0x38, 0xcb, 0xad, 0x32, 0x7c, 0x3d, 0x1e, 0x51, 0xe6,
	0x9f, 0x9e, 0xf0, 0x53, 0x78, 0x9b, 0xce, 0x18, 0x72, 0x9d, 0x88, 0xc7, 0x22, 0x97, 0x2e, 0x74,
	0xc1, 0x51, 0xbd, 0xe4, 0x49, 0x9e, 0xd6, 0xa2, 0xe4, 0x51, 0x96, 0x5d, 0xe8, 0xb0, 0x7c, 0x3e,
	0x2f, 0x92, 0x89, 0xcc, 0x0b, 0xe1, 0xef, 0x2d, 0x58, 0xe8, 0xa5, 0x41, 0xe2, 0x3f, 0x5a, 0xb0,
	0x79, 0x5a, 0x15, 0xaf, 0x3c, 0xd3, 0x57, 0x00, 0x4c, 0x24, 0xbe, 0x94, 0xfe, 0xe1, 0xa9, 0x2b,
	0x72, 0x75, 0xd9, 0x49, 0xc0, 0x44, 0x29, 0xfe, 0x06, 0x3a, 0xa6, 0xf3, 0xa5, 0xba, 0x69, 0xd5,
	0x3b, 0xff, 0xd9, 0x22, 0x2b, 0x87, 0x8c, 0xab, 0x52, 0x7f, 0x01, 0x03, 0xf4, 0x6d, 0xaa, 0x4c,
	0x5a, 0xd6, 0xa4, 0xee, 0xc5, 0x7d, 0xd2, 0x51, 0xd2, 0xc7, 0x27, 0x2d, 0x3e, 0x83, 0x0d, 0x25,
	0x59, 0x7e, 0x57, 0x79, 0xad, 0x58, 0xaf, 0xdd, 0x1a, 0xaf, 0xc5, 0xbe, 0x92, 0xae, 0x55, 0x95,
	0x2e, 0xbf, 0x83, 0x1b, 0xf6, 0xca, 0xc5, 0xbd, 0x1f, 0xc7, 0x35, 0x2e, 0x4b, 0xe5, 0x74, 0xd7,
	0xcc, 0x7f, 0x9c, 0x73, 0x2d, 0x1f, 0x48, 0x17, 0x17, 0xa0, 0x77, 0xbf, 0x85, 0xcd, 0x25, 0x4a,
	0x38, 0x80, 0xd6, 0x2d, 0x3e, 0xf8, 0x19, 0x32, 0xcb, 0x70, 0x0b, 0x56, 0xef, 0xd2, 0x59, 0x81,
	0x7e, 0x44, 0xdc, 0xc7, 0x49, 0xf3, 0x45, 0xe3, 0x87, 0x93, 0xdf, 0x5e, 0x4c, 0x98, 0x9e, 0x16,
	0x37, 0x07, 0x34, 0x9f, 0x1f, 0x72, 0xd4, 0xaf, 0x73, 0x79, 0x6b, 0xae, 0x1e, 0xa3, 0x38, 0x47,
	0x35, 0x3d, 0x4c, 0x05, 0x3b, 0x14, 0xb7, 0x13, 0xfb, 0xfb, 0x38, 0x7c, 0xd3, 0xb6, 0x7f, 0xb3,
	0x9f, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xad, 0xee, 0x22, 0x71, 0x81, 0x07, 0x00, 0x00,
}
