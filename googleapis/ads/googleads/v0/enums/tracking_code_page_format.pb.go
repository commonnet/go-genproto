// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/tracking_code_page_format.proto

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

// The format of the web page where the tracking tag and snippet will be
// installed.
type TrackingCodePageFormatEnum_TrackingCodePageFormat int32

const (
	// Not specified.
	TrackingCodePageFormatEnum_UNSPECIFIED TrackingCodePageFormatEnum_TrackingCodePageFormat = 0
	// Used for return value only. Represents value unknown in this version.
	TrackingCodePageFormatEnum_UNKNOWN TrackingCodePageFormatEnum_TrackingCodePageFormat = 1
	// Standard HTML page format.
	TrackingCodePageFormatEnum_HTML TrackingCodePageFormatEnum_TrackingCodePageFormat = 2
	// Google AMP page format.
	TrackingCodePageFormatEnum_AMP TrackingCodePageFormatEnum_TrackingCodePageFormat = 3
)

var TrackingCodePageFormatEnum_TrackingCodePageFormat_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "HTML",
	3: "AMP",
}

var TrackingCodePageFormatEnum_TrackingCodePageFormat_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"HTML":        2,
	"AMP":         3,
}

func (x TrackingCodePageFormatEnum_TrackingCodePageFormat) String() string {
	return proto.EnumName(TrackingCodePageFormatEnum_TrackingCodePageFormat_name, int32(x))
}

func (TrackingCodePageFormatEnum_TrackingCodePageFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_848931c9545ebcbb, []int{0, 0}
}

// Container for enum describing the format of the web page where the tracking
// tag and snippet will be installed.
type TrackingCodePageFormatEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingCodePageFormatEnum) Reset()         { *m = TrackingCodePageFormatEnum{} }
func (m *TrackingCodePageFormatEnum) String() string { return proto.CompactTextString(m) }
func (*TrackingCodePageFormatEnum) ProtoMessage()    {}
func (*TrackingCodePageFormatEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_848931c9545ebcbb, []int{0}
}

func (m *TrackingCodePageFormatEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingCodePageFormatEnum.Unmarshal(m, b)
}
func (m *TrackingCodePageFormatEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingCodePageFormatEnum.Marshal(b, m, deterministic)
}
func (m *TrackingCodePageFormatEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingCodePageFormatEnum.Merge(m, src)
}
func (m *TrackingCodePageFormatEnum) XXX_Size() int {
	return xxx_messageInfo_TrackingCodePageFormatEnum.Size(m)
}
func (m *TrackingCodePageFormatEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingCodePageFormatEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingCodePageFormatEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.TrackingCodePageFormatEnum_TrackingCodePageFormat", TrackingCodePageFormatEnum_TrackingCodePageFormat_name, TrackingCodePageFormatEnum_TrackingCodePageFormat_value)
	proto.RegisterType((*TrackingCodePageFormatEnum)(nil), "google.ads.googleads.v0.enums.TrackingCodePageFormatEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/tracking_code_page_format.proto", fileDescriptor_848931c9545ebcbb)
}

var fileDescriptor_848931c9545ebcbb = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x92, 0xa2, 0xc4, 0xe4, 0xec, 0xcc, 0xbc, 0xf4, 0xf8, 0xe4, 0xfc, 0x94,
	0xd4, 0xf8, 0x82, 0xc4, 0xf4, 0xd4, 0xf8, 0xb4, 0xfc, 0xa2, 0xdc, 0xc4, 0x12, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x1e, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x76, 0xbd, 0x32,
	0x03, 0x3d, 0xb0, 0x76, 0xa5, 0x74, 0x2e, 0xa9, 0x10, 0xa8, 0x09, 0xce, 0xf9, 0x29, 0xa9, 0x01,
	0x89, 0xe9, 0xa9, 0x6e, 0x60, 0xed, 0xae, 0x79, 0xa5, 0xb9, 0x4a, 0x9e, 0x5c, 0x62, 0xd8, 0x65,
	0x85, 0xf8, 0xb9, 0xb8, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04,
	0x18, 0x84, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x85, 0x38,
	0xb8, 0x58, 0x3c, 0x42, 0x7c, 0x7d, 0x04, 0x98, 0x84, 0xd8, 0xb9, 0x98, 0x1d, 0x7d, 0x03, 0x04,
	0x98, 0x9d, 0x3e, 0x30, 0x72, 0x29, 0x26, 0xe7, 0xe7, 0xea, 0xe1, 0x75, 0x8e, 0x93, 0x34, 0x76,
	0xeb, 0x02, 0x40, 0x5e, 0x09, 0x60, 0x8c, 0x72, 0x82, 0xea, 0x4e, 0xcf, 0xcf, 0x49, 0xcc, 0x4b,
	0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f, 0xcd, 0x03, 0x7b, 0x14, 0x16, 0x36, 0x05, 0x99, 0xc5,
	0x38, 0x82, 0xca, 0x1a, 0x4c, 0x2e, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5, 0x24, 0xeb, 0x0e,
	0x31, 0xca, 0x31, 0xa5, 0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4, 0x40, 0x1e, 0x2f, 0x3e,
	0x05, 0x93, 0x8f, 0x71, 0x4c, 0x29, 0x8e, 0x81, 0xcb, 0xc7, 0x84, 0x19, 0xc4, 0x80, 0xe5, 0x5f,
	0x31, 0x29, 0x42, 0x04, 0xad, 0xac, 0x1c, 0x53, 0x8a, 0xad, 0xac, 0xe0, 0x2a, 0xac, 0xac, 0xc2,
	0x0c, 0xac, 0xac, 0xc0, 0x6a, 0x92, 0xd8, 0xc0, 0x0e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xc8, 0xf5, 0xd1, 0x82, 0xc2, 0x01, 0x00, 0x00,
}
