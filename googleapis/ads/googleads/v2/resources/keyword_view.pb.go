// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/keyword_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A keyword view.
type KeywordView struct {
	// The resource name of the keyword view.
	// Keyword view resource names have the form:
	//
	// `customers/{customer_id}/keywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordView) Reset()         { *m = KeywordView{} }
func (m *KeywordView) String() string { return proto.CompactTextString(m) }
func (*KeywordView) ProtoMessage()    {}
func (*KeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f8be4cdb6fb03af, []int{0}
}

func (m *KeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordView.Unmarshal(m, b)
}
func (m *KeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordView.Marshal(b, m, deterministic)
}
func (m *KeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordView.Merge(m, src)
}
func (m *KeywordView) XXX_Size() int {
	return xxx_messageInfo_KeywordView.Size(m)
}
func (m *KeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordView proto.InternalMessageInfo

func (m *KeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*KeywordView)(nil), "google.ads.googleads.v2.resources.KeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/keyword_view.proto", fileDescriptor_7f8be4cdb6fb03af)
}

var fileDescriptor_7f8be4cdb6fb03af = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x69, 0x05, 0xc1, 0xaa, 0x20, 0xb3, 0x12, 0x71, 0xe1, 0x28, 0x03, 0xae, 0x12, 0x88,
	0xae, 0xe2, 0x2a, 0xb3, 0x19, 0x50, 0x90, 0x61, 0x16, 0x5d, 0x48, 0x61, 0x88, 0x93, 0x10, 0x82,
	0xd3, 0xbc, 0x92, 0x57, 0x5b, 0xbc, 0x8e, 0x4b, 0x8f, 0xe2, 0x51, 0xbc, 0x83, 0x20, 0x6d, 0x4c,
	0x70, 0xe5, 0xec, 0x7e, 0x92, 0xef, 0xff, 0xde, 0x4b, 0x8a, 0x5b, 0x03, 0x60, 0xb6, 0x9a, 0x4a,
	0x85, 0x34, 0xc4, 0x21, 0x75, 0x8c, 0x7a, 0x8d, 0xf0, 0xea, 0x37, 0x1a, 0xe9, 0x8b, 0x7e, 0xeb,
	0xc1, 0xab, 0x75, 0x67, 0x75, 0x4f, 0x1a, 0x0f, 0x2d, 0x4c, 0xa6, 0x01, 0x25, 0x52, 0x21, 0x49,
	0x2d, 0xd2, 0x31, 0x92, 0x5a, 0x67, 0xe7, 0x51, 0xdc, 0x58, 0x2a, 0x9d, 0x83, 0x56, 0xb6, 0x16,
	0x1c, 0x06, 0xc1, 0x25, 0x2b, 0x0e, 0x1f, 0x82, 0xb6, 0xb4, 0xba, 0x9f, 0x5c, 0x15, 0xc7, 0xb1,
	0xb9, 0x76, 0xb2, 0xd6, 0xa7, 0xd9, 0x45, 0x76, 0x7d, 0xb0, 0x3a, 0x8a, 0x87, 0x8f, 0xb2, 0xd6,
	0xf3, 0xef, 0xac, 0x98, 0x6d, 0xa0, 0x26, 0x3b, 0x67, 0xcf, 0x4f, 0xfe, 0xb8, 0x97, 0xc3, 0xbc,
	0x65, 0xf6, 0x74, 0xff, 0x5b, 0x33, 0xb0, 0x95, 0xce, 0x10, 0xf0, 0x86, 0x1a, 0xed, 0xc6, 0x6d,
	0xe2, 0xc3, 0x1b, 0x8b, 0xff, 0xfc, 0xc3, 0x5d, 0x4a, 0xef, 0xf9, 0xde, 0x42, 0x88, 0x8f, 0x7c,
	0xba, 0x08, 0x4a, 0xa1, 0x90, 0x84, 0x38, 0xa4, 0x92, 0x91, 0x55, 0x24, 0x3f, 0x23, 0x53, 0x09,
	0x85, 0x55, 0x62, 0xaa, 0x92, 0x55, 0x89, 0xf9, 0xca, 0x67, 0xe1, 0x82, 0x73, 0xa1, 0x90, 0xf3,
	0x44, 0x71, 0x5e, 0x32, 0xce, 0x13, 0xf7, 0xbc, 0x3f, 0x2e, 0x7b, 0xf3, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x61, 0x2f, 0x00, 0xf1, 0xb3, 0x01, 0x00, 0x00,
}
