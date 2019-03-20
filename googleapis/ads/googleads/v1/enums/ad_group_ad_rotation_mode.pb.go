// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/ad_group_ad_rotation_mode.proto

package enums

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// The possible ad rotation modes of an ad group.
type AdGroupAdRotationModeEnum_AdGroupAdRotationMode int32

const (
	// The ad rotation mode has not been specified.
	AdGroupAdRotationModeEnum_UNSPECIFIED AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupAdRotationModeEnum_UNKNOWN AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 1
	// Optimize ad group ads based on clicks or conversions.
	AdGroupAdRotationModeEnum_OPTIMIZE AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 2
	// Rotate evenly forever.
	AdGroupAdRotationModeEnum_ROTATE_FOREVER AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 3
)

var AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "ROTATE_FOREVER",
}

var AdGroupAdRotationModeEnum_AdGroupAdRotationMode_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"OPTIMIZE":       2,
	"ROTATE_FOREVER": 3,
}

func (x AdGroupAdRotationModeEnum_AdGroupAdRotationMode) String() string {
	return proto.EnumName(AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name, int32(x))
}

func (AdGroupAdRotationModeEnum_AdGroupAdRotationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01a65cb679b1721f, []int{0, 0}
}

// Container for enum describing possible ad rotation modes of ads within an
// ad group.
type AdGroupAdRotationModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdRotationModeEnum) Reset()         { *m = AdGroupAdRotationModeEnum{} }
func (m *AdGroupAdRotationModeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdRotationModeEnum) ProtoMessage()    {}
func (*AdGroupAdRotationModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a65cb679b1721f, []int{0}
}

func (m *AdGroupAdRotationModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Unmarshal(m, b)
}
func (m *AdGroupAdRotationModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupAdRotationModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdRotationModeEnum.Merge(m, src)
}
func (m *AdGroupAdRotationModeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Size(m)
}
func (m *AdGroupAdRotationModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdRotationModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdRotationModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode", AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name, AdGroupAdRotationModeEnum_AdGroupAdRotationMode_value)
	proto.RegisterType((*AdGroupAdRotationModeEnum)(nil), "google.ads.googleads.v1.enums.AdGroupAdRotationModeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/ad_group_ad_rotation_mode.proto", fileDescriptor_01a65cb679b1721f)
}

var fileDescriptor_01a65cb679b1721f = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x18, 0xc5, 0xff, 0xeb, 0xe0, 0xaf, 0x64, 0xa2, 0xa5, 0xe0, 0x85, 0xc3, 0x5d, 0x6c, 0x0f, 0x90,
	0x52, 0xbc, 0x8b, 0x78, 0x91, 0x69, 0x36, 0x8a, 0xac, 0x2d, 0x75, 0xeb, 0x60, 0x14, 0x4a, 0x34,
	0x25, 0x0c, 0xd6, 0x7c, 0xa5, 0xe9, 0xf6, 0x40, 0x5e, 0xfa, 0x28, 0x3e, 0x8a, 0x37, 0xbe, 0x82,
	0xb4, 0xd9, 0x76, 0x35, 0xbd, 0x09, 0x87, 0x9c, 0xef, 0x77, 0xf2, 0xe5, 0xa0, 0x07, 0x09, 0x20,
	0x37, 0xb9, 0xcb, 0x85, 0x76, 0x8d, 0x6c, 0xd4, 0xce, 0x73, 0x73, 0xb5, 0x2d, 0xb4, 0xcb, 0x45,
	0x26, 0x2b, 0xd8, 0x96, 0x19, 0x17, 0x59, 0x05, 0x35, 0xaf, 0xd7, 0xa0, 0xb2, 0x02, 0x44, 0x8e,
	0xcb, 0x0a, 0x6a, 0x70, 0x06, 0x86, 0xc1, 0x5c, 0x68, 0x7c, 0xc4, 0xf1, 0xce, 0xc3, 0x2d, 0xde,
	0xbf, 0x3d, 0xa4, 0x97, 0x6b, 0x97, 0x2b, 0xb5, 0x0f, 0xd0, 0x06, 0x1e, 0xd5, 0xe8, 0x86, 0x8a,
	0x69, 0x13, 0x4f, 0x45, 0xbc, 0xf7, 0x66, 0x20, 0x72, 0xa6, 0xb6, 0xc5, 0x68, 0x89, 0xae, 0x4f,
	0x9a, 0xce, 0x15, 0xea, 0x2d, 0x82, 0x97, 0x88, 0x3d, 0xfa, 0x13, 0x9f, 0x3d, 0xd9, 0xff, 0x9c,
	0x1e, 0x3a, 0x5b, 0x04, 0xcf, 0x41, 0xb8, 0x0c, 0xec, 0x8e, 0x73, 0x81, 0xce, 0xc3, 0x68, 0xee,
	0xcf, 0xfc, 0x15, 0xb3, 0x2d, 0xc7, 0x41, 0x97, 0x71, 0x38, 0xa7, 0x73, 0x96, 0x4d, 0xc2, 0x98,
	0x25, 0x2c, 0xb6, 0xbb, 0xe3, 0xef, 0x0e, 0x1a, 0xbe, 0x41, 0x81, 0xff, 0xdc, 0x7c, 0xdc, 0x3f,
	0xf9, 0x78, 0xd4, 0xec, 0x1d, 0x75, 0x56, 0xe3, 0x3d, 0x2c, 0x61, 0xc3, 0x95, 0xc4, 0x50, 0x49,
	0x57, 0xe6, 0xaa, 0xfd, 0xd5, 0xa1, 0xc5, 0x72, 0xad, 0x7f, 0x29, 0xf5, 0xbe, 0x3d, 0xdf, 0xad,
	0xee, 0x94, 0xd2, 0x0f, 0x6b, 0x30, 0x35, 0x51, 0x54, 0x68, 0x6c, 0x64, 0xa3, 0x12, 0x0f, 0x37,
	0x2d, 0xe8, 0xcf, 0x83, 0x9f, 0x52, 0xa1, 0xd3, 0xa3, 0x9f, 0x26, 0x5e, 0xda, 0xfa, 0x5f, 0xd6,
	0xd0, 0x5c, 0x12, 0x42, 0x85, 0x26, 0xe4, 0x38, 0x41, 0x48, 0xe2, 0x11, 0xd2, 0xce, 0xbc, 0xfe,
	0x6f, 0x17, 0xbb, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x22, 0xa6, 0xd8, 0x5e, 0xec, 0x01, 0x00,
	0x00,
}
