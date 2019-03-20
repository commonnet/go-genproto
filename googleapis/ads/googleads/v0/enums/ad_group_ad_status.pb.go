// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_group_ad_status.proto

package enums

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

// The possible statuses of an AdGroupAd.
type AdGroupAdStatusEnum_AdGroupAdStatus int32

const (
	// No value has been specified.
	AdGroupAdStatusEnum_UNSPECIFIED AdGroupAdStatusEnum_AdGroupAdStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupAdStatusEnum_UNKNOWN AdGroupAdStatusEnum_AdGroupAdStatus = 1
	// The ad group ad is enabled.
	AdGroupAdStatusEnum_ENABLED AdGroupAdStatusEnum_AdGroupAdStatus = 2
	// The ad group ad is paused.
	AdGroupAdStatusEnum_PAUSED AdGroupAdStatusEnum_AdGroupAdStatus = 3
	// The ad group ad is removed.
	AdGroupAdStatusEnum_REMOVED AdGroupAdStatusEnum_AdGroupAdStatus = 4
)

var AdGroupAdStatusEnum_AdGroupAdStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupAdStatusEnum_AdGroupAdStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupAdStatusEnum_AdGroupAdStatus) String() string {
	return proto.EnumName(AdGroupAdStatusEnum_AdGroupAdStatus_name, int32(x))
}

func (AdGroupAdStatusEnum_AdGroupAdStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_137acffc11908561, []int{0, 0}
}

// Container for enum describing possible statuses of an AdGroupAd.
type AdGroupAdStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdStatusEnum) Reset()         { *m = AdGroupAdStatusEnum{} }
func (m *AdGroupAdStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdStatusEnum) ProtoMessage()    {}
func (*AdGroupAdStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_137acffc11908561, []int{0}
}

func (m *AdGroupAdStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupAdStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupAdStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdStatusEnum.Merge(m, src)
}
func (m *AdGroupAdStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdStatusEnum.Size(m)
}
func (m *AdGroupAdStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupAdStatusEnum_AdGroupAdStatus", AdGroupAdStatusEnum_AdGroupAdStatus_name, AdGroupAdStatusEnum_AdGroupAdStatus_value)
	proto.RegisterType((*AdGroupAdStatusEnum)(nil), "google.ads.googleads.v0.enums.AdGroupAdStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_group_ad_status.proto", fileDescriptor_137acffc11908561)
}

var fileDescriptor_137acffc11908561 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x6d, 0x47, 0x46, 0xc8, 0x2c, 0xa6, 0x54, 0xb7, 0xb3, 0x98, 0x39, 0x40, 0x5a, 0x10,
	0x5c, 0xc4, 0x55, 0x6a, 0x63, 0x19, 0xd4, 0x4e, 0xb1, 0xb4, 0x82, 0x14, 0x4a, 0x34, 0x25, 0x08,
	0x6d, 0x53, 0x9a, 0x76, 0x0e, 0xe4, 0xd2, 0xa3, 0x78, 0x13, 0x3d, 0x85, 0x24, 0x71, 0xba, 0x18,
	0xd0, 0x4d, 0xf8, 0xf3, 0xfe, 0xf7, 0x25, 0xef, 0x7f, 0xe0, 0x8a, 0x0b, 0xc1, 0xeb, 0xca, 0xa3,
	0x4c, 0x7a, 0x46, 0x2a, 0xb5, 0xf7, 0xbd, 0xaa, 0x1d, 0x1b, 0xe9, 0x51, 0x56, 0xf2, 0x5e, 0x8c,
	0x5d, 0x49, 0x59, 0x29, 0x07, 0x3a, 0x8c, 0x12, 0x76, 0xbd, 0x18, 0x84, 0xbb, 0x32, 0xcd, 0x90,
	0x32, 0x09, 0x27, 0x0e, 0xee, 0x7d, 0xa8, 0xb9, 0x4d, 0x0d, 0xce, 0x31, 0x8b, 0x14, 0x89, 0x59,
	0xaa, 0x39, 0xd2, 0x8e, 0xcd, 0x26, 0x03, 0xcb, 0xa3, 0xb2, 0xbb, 0x04, 0x8b, 0x2c, 0x4e, 0x13,
	0x72, 0xb3, 0xbd, 0xdd, 0x92, 0xd0, 0x39, 0x71, 0x17, 0xe0, 0x2c, 0x8b, 0xef, 0xe2, 0xdd, 0x53,
	0xec, 0x58, 0xea, 0x42, 0x62, 0x1c, 0xdc, 0x93, 0xd0, 0xb1, 0x5d, 0x00, 0xe6, 0x09, 0xce, 0x52,
	0x12, 0x3a, 0x33, 0x65, 0x3c, 0x92, 0x87, 0x5d, 0x4e, 0x42, 0xe7, 0x34, 0xf8, 0xb2, 0xc0, 0xfa,
	0x55, 0x34, 0xf0, 0xdf, 0x99, 0x82, 0x8b, 0xa3, 0xaf, 0x13, 0x15, 0x24, 0xb1, 0x9e, 0x83, 0x5f,
	0x8c, 0x8b, 0x9a, 0xb6, 0x1c, 0x8a, 0x9e, 0x7b, 0xbc, 0x6a, 0x75, 0xcc, 0xc3, 0x4a, 0xba, 0x37,
	0xf9, 0xc7, 0x86, 0xae, 0xf5, 0xf9, 0x6e, 0xcf, 0x22, 0x8c, 0x3f, 0xec, 0x55, 0x64, 0x9e, 0xc2,
	0x4c, 0x42, 0x23, 0x95, 0xca, 0x7d, 0xa8, 0xd2, 0xcb, 0xcf, 0x83, 0x5f, 0x60, 0x26, 0x8b, 0xc9,
	0x2f, 0x72, 0xbf, 0xd0, 0xfe, 0xb7, 0xbd, 0x36, 0x45, 0x84, 0x30, 0x93, 0x08, 0x4d, 0x1d, 0x08,
	0xe5, 0x3e, 0x42, 0xba, 0xe7, 0x65, 0xae, 0x07, 0xbb, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x08,
	0x7f, 0x08, 0x9a, 0xb9, 0x01, 0x00, 0x00,
}
