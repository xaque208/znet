// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/rpc.proto

package rpc

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

type SearchRequest struct {
	Hosts                string   `protobuf:"bytes,1,opt,name=hosts,proto3" json:"hosts,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Kernel               string   `protobuf:"bytes,3,opt,name=kernel,proto3" json:"kernel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9874a201429861e, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetHosts() string {
	if m != nil {
		return m.Hosts
	}
	return ""
}

func (m *SearchRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SearchRequest) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

type SearchResponse struct {
	Hosts                []*Host  `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9874a201429861e, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetHosts() []*Host {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type Host struct {
	CommonName           string   `protobuf:"bytes,1,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Platform             string   `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9874a201429861e, []int{2}
}

func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *Host) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Host) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Host) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*Host)(nil), "Host")
}

func init() { proto.RegisterFile("rpc/rpc.proto", fileDescriptor_d9874a201429861e) }

var fileDescriptor_d9874a201429861e = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x4d, 0x83, 0x9d, 0xd0, 0x0a, 0x83, 0x48, 0xa8, 0x07, 0x4b, 0x4e, 0xf5, 0x60,
	0x84, 0x0a, 0x3e, 0x83, 0x5e, 0x3c, 0xd4, 0x93, 0x27, 0x59, 0xd7, 0x91, 0x16, 0xb3, 0x3b, 0xeb,
	0xec, 0x28, 0xe4, 0xed, 0xa5, 0x9b, 0x58, 0x9b, 0xdb, 0x7e, 0xff, 0xc0, 0xec, 0x37, 0x3f, 0xcc,
	0x24, 0xd8, 0x5b, 0x09, 0xb6, 0x0e, 0xc2, 0xca, 0xd5, 0x0b, 0xcc, 0x9e, 0xc9, 0x88, 0xdd, 0x6e,
	0xe8, 0xeb, 0x9b, 0xa2, 0xe2, 0x39, 0x4c, 0xb6, 0x1c, 0x35, 0x96, 0xa3, 0xe5, 0x68, 0x35, 0xdd,
	0x74, 0x80, 0x0b, 0x38, 0x6d, 0xd8, 0x1a, 0xdd, 0xb1, 0x2f, 0x4f, 0xd2, 0xe0, 0xc0, 0x78, 0x01,
	0xf9, 0x27, 0x89, 0xa7, 0xa6, 0x1c, 0xa7, 0x49, 0x4f, 0xd5, 0x0d, 0xcc, 0xff, 0x56, 0xc7, 0xc0,
	0x3e, 0x12, 0x5e, 0xfe, 0xef, 0x1e, 0xaf, 0x8a, 0xf5, 0xa4, 0x7e, 0xe0, 0xa8, 0xfd, 0x17, 0x55,
	0x0b, 0xd9, 0x1e, 0xf1, 0x0a, 0x0a, 0xcb, 0xce, 0xb1, 0x7f, 0xf5, 0xc6, 0x51, 0xaf, 0x01, 0x5d,
	0xf4, 0x64, 0x1c, 0xe1, 0x12, 0x8a, 0x77, 0x8a, 0x56, 0x76, 0xe1, 0x48, 0xe7, 0x38, 0xda, 0xdb,
	0x86, 0xc6, 0xe8, 0x07, 0x8b, 0xeb, 0x9d, 0x0e, 0x8c, 0x08, 0x99, 0xb6, 0x81, 0xca, 0x2c, 0xe5,
	0xe9, 0xbd, 0xbe, 0x87, 0xe9, 0xa3, 0xff, 0x21, 0xaf, 0x2c, 0x2d, 0x5e, 0x43, 0xde, 0x69, 0xe3,
	0xbc, 0x1e, 0x54, 0xb3, 0x38, 0xab, 0x87, 0xf7, 0xbc, 0xe5, 0xa9, 0xc3, 0xbb, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0x12, 0x99, 0x2b, 0x54, 0x01, 0x00, 0x00,
}
