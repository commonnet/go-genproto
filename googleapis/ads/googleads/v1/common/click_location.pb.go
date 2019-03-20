// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/click_location.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Location criteria associated with a click.
type ClickLocation struct {
	// The city location criterion associated with the impression.
	City *wrappers.StringValue `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	// The country location criterion associated with the impression.
	Country *wrappers.StringValue `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	// The metro location criterion associated with the impression.
	Metro *wrappers.StringValue `protobuf:"bytes,3,opt,name=metro,proto3" json:"metro,omitempty"`
	// The most specific location criterion associated with the impression.
	MostSpecific *wrappers.StringValue `protobuf:"bytes,4,opt,name=most_specific,json=mostSpecific,proto3" json:"most_specific,omitempty"`
	// The region location criterion associated with the impression.
	Region               *wrappers.StringValue `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClickLocation) Reset()         { *m = ClickLocation{} }
func (m *ClickLocation) String() string { return proto.CompactTextString(m) }
func (*ClickLocation) ProtoMessage()    {}
func (*ClickLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_68f6b7e96a137729, []int{0}
}

func (m *ClickLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickLocation.Unmarshal(m, b)
}
func (m *ClickLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickLocation.Marshal(b, m, deterministic)
}
func (m *ClickLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickLocation.Merge(m, src)
}
func (m *ClickLocation) XXX_Size() int {
	return xxx_messageInfo_ClickLocation.Size(m)
}
func (m *ClickLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickLocation.DiscardUnknown(m)
}

var xxx_messageInfo_ClickLocation proto.InternalMessageInfo

func (m *ClickLocation) GetCity() *wrappers.StringValue {
	if m != nil {
		return m.City
	}
	return nil
}

func (m *ClickLocation) GetCountry() *wrappers.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *ClickLocation) GetMetro() *wrappers.StringValue {
	if m != nil {
		return m.Metro
	}
	return nil
}

func (m *ClickLocation) GetMostSpecific() *wrappers.StringValue {
	if m != nil {
		return m.MostSpecific
	}
	return nil
}

func (m *ClickLocation) GetRegion() *wrappers.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func init() {
	proto.RegisterType((*ClickLocation)(nil), "google.ads.googleads.v1.common.ClickLocation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/click_location.proto", fileDescriptor_68f6b7e96a137729)
}

var fileDescriptor_68f6b7e96a137729 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6a, 0xeb, 0x30,
	0x14, 0xc6, 0xb1, 0xf3, 0xe7, 0x82, 0xee, 0xcd, 0xe2, 0xc9, 0x84, 0x10, 0x2e, 0x99, 0xee, 0x24,
	0x5d, 0x27, 0xa5, 0x83, 0x3a, 0x39, 0x29, 0x64, 0xe9, 0x10, 0x1a, 0xf0, 0x50, 0x0c, 0xc1, 0x91,
	0x1d, 0x21, 0x6a, 0xeb, 0x18, 0x49, 0x49, 0xc9, 0xeb, 0x74, 0xe8, 0xd0, 0x47, 0xe9, 0xa3, 0x14,
	0xfa, 0x0e, 0xc5, 0x96, 0x1d, 0xe8, 0xd0, 0xe2, 0xc9, 0x07, 0xeb, 0xf7, 0xfb, 0x3e, 0x63, 0x1d,
	0xb4, 0xe0, 0x00, 0x3c, 0xcf, 0x48, 0x92, 0x6a, 0x62, 0xc7, 0x6a, 0x3a, 0x05, 0x84, 0x41, 0x51,
	0x80, 0x24, 0x2c, 0x17, 0xec, 0x71, 0x97, 0x03, 0x4b, 0x8c, 0x00, 0x89, 0x4b, 0x05, 0x06, 0xbc,
	0xa9, 0x25, 0x71, 0x92, 0x6a, 0x7c, 0x91, 0xf0, 0x29, 0xc0, 0x56, 0x1a, 0x37, 0xe7, 0xa4, 0xa6,
	0xf7, 0xc7, 0x03, 0x79, 0x52, 0x49, 0x59, 0x66, 0x4a, 0x5b, 0x7f, 0x3c, 0x69, 0x4b, 0x4b, 0x41,
	0x12, 0x29, 0xc1, 0xd4, 0xe1, 0xcd, 0xe9, 0xec, 0xc5, 0x45, 0xa3, 0x55, 0x55, 0x7b, 0xd7, 0xb4,
	0x7a, 0xff, 0x51, 0x9f, 0x09, 0x73, 0xf6, 0x9d, 0xbf, 0xce, 0xbf, 0xdf, 0xf3, 0x49, 0xd3, 0x89,
	0xdb, 0x78, 0xbc, 0x35, 0x4a, 0x48, 0x1e, 0x25, 0xf9, 0x31, 0xbb, 0xaf, 0x49, 0xef, 0x1a, 0xfd,
	0x62, 0x70, 0x94, 0x46, 0x9d, 0x7d, 0xb7, 0x83, 0xd4, 0xc2, 0xde, 0x1c, 0x0d, 0x8a, 0xcc, 0x28,
	0xf0, 0x7b, 0x1d, 0x2c, 0x8b, 0x7a, 0x21, 0x1a, 0x15, 0xa0, 0xcd, 0x4e, 0x97, 0x19, 0x13, 0x07,
	0xc1, 0xfc, 0x7e, 0x07, 0xf7, 0x4f, 0xa5, 0x6c, 0x1b, 0xc3, 0xbb, 0x42, 0x43, 0x95, 0x71, 0x01,
	0xd2, 0x1f, 0x74, 0x70, 0x1b, 0x76, 0xf9, 0xe1, 0xa0, 0x19, 0x83, 0x02, 0xff, 0x7c, 0x1b, 0x4b,
	0xef, 0xcb, 0xcf, 0xdc, 0x54, 0x89, 0x1b, 0xe7, 0xe1, 0xb6, 0xb1, 0x38, 0xe4, 0x89, 0xe4, 0x18,
	0x14, 0x27, 0x3c, 0x93, 0x75, 0x5f, 0xbb, 0x08, 0xa5, 0xd0, 0xdf, 0xed, 0xc5, 0x8d, 0x7d, 0x3c,
	0xbb, 0xbd, 0x75, 0x18, 0xbe, 0xba, 0xd3, 0xb5, 0x0d, 0x0b, 0x53, 0x8d, 0xed, 0x58, 0x4d, 0x51,
	0x80, 0x57, 0x35, 0xf6, 0xd6, 0x02, 0x71, 0x98, 0xea, 0xf8, 0x02, 0xc4, 0x51, 0x10, 0x5b, 0xe0,
	0xdd, 0x9d, 0xd9, 0xb7, 0x94, 0x86, 0xa9, 0xa6, 0xf4, 0x82, 0x50, 0x1a, 0x05, 0x94, 0x5a, 0x68,
	0x3f, 0xac, 0xbf, 0x6e, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff, 0x40, 0xe0, 0xe9, 0xb4, 0xb4, 0x02,
	0x00, 0x00,
}
