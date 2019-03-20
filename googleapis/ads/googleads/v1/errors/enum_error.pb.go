// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/enum_error.proto

package errors

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

// Enum describing possible enum errors.
type EnumErrorEnum_EnumError int32

const (
	// Enum unspecified.
	EnumErrorEnum_UNSPECIFIED EnumErrorEnum_EnumError = 0
	// The received error code is not known in this version.
	EnumErrorEnum_UNKNOWN EnumErrorEnum_EnumError = 1
	// The enum value is not permitted.
	EnumErrorEnum_ENUM_VALUE_NOT_PERMITTED EnumErrorEnum_EnumError = 3
)

var EnumErrorEnum_EnumError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "ENUM_VALUE_NOT_PERMITTED",
}

var EnumErrorEnum_EnumError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"ENUM_VALUE_NOT_PERMITTED": 3,
}

func (x EnumErrorEnum_EnumError) String() string {
	return proto.EnumName(EnumErrorEnum_EnumError_name, int32(x))
}

func (EnumErrorEnum_EnumError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14e62fef9dcb7b60, []int{0, 0}
}

// Container for enum describing possible enum errors.
type EnumErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumErrorEnum) Reset()         { *m = EnumErrorEnum{} }
func (m *EnumErrorEnum) String() string { return proto.CompactTextString(m) }
func (*EnumErrorEnum) ProtoMessage()    {}
func (*EnumErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_14e62fef9dcb7b60, []int{0}
}

func (m *EnumErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumErrorEnum.Unmarshal(m, b)
}
func (m *EnumErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumErrorEnum.Marshal(b, m, deterministic)
}
func (m *EnumErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumErrorEnum.Merge(m, src)
}
func (m *EnumErrorEnum) XXX_Size() int {
	return xxx_messageInfo_EnumErrorEnum.Size(m)
}
func (m *EnumErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_EnumErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.EnumErrorEnum_EnumError", EnumErrorEnum_EnumError_name, EnumErrorEnum_EnumError_value)
	proto.RegisterType((*EnumErrorEnum)(nil), "google.ads.googleads.v1.errors.EnumErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/enum_error.proto", fileDescriptor_14e62fef9dcb7b60)
}

var fileDescriptor_14e62fef9dcb7b60 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0x84, 0x40,
	0x18, 0xc7, 0xd3, 0x85, 0xa2, 0x59, 0x2a, 0xf1, 0x14, 0xb1, 0xec, 0xc1, 0x07, 0x98, 0x41, 0xba,
	0x4d, 0xa7, 0xd9, 0x9c, 0x44, 0x6a, 0x67, 0xa5, 0xd4, 0x22, 0x04, 0xb1, 0x94, 0x41, 0x58, 0x67,
	0xc4, 0x71, 0xf7, 0x81, 0x3a, 0xf6, 0x28, 0xbd, 0x49, 0x3d, 0x45, 0xe8, 0xac, 0xde, 0xea, 0x34,
	0x3f, 0x86, 0xdf, 0xff, 0xfb, 0xfe, 0x7c, 0x00, 0x71, 0x29, 0xf9, 0xb6, 0x44, 0x79, 0xa1, 0x0e,
	0xd8, 0xd3, 0xde, 0x45, 0x65, 0xdb, 0xca, 0x56, 0xa1, 0x52, 0xec, 0xea, 0x6c, 0x60, 0xd8, 0xb4,
	0xb2, 0x93, 0xf6, 0x52, 0x5b, 0x30, 0x2f, 0x14, 0x9c, 0x02, 0x70, 0xef, 0x42, 0x1d, 0xb8, 0x5a,
	0x8c, 0x03, 0x9b, 0x0a, 0xe5, 0x42, 0xc8, 0x2e, 0xef, 0x2a, 0x29, 0x94, 0x4e, 0x3b, 0x2f, 0xe0,
	0x8c, 0x8a, 0x5d, 0x4d, 0x7b, 0xb7, 0x07, 0xc7, 0x07, 0xa7, 0xd3, 0x87, 0x7d, 0x01, 0xe6, 0x31,
	0x7b, 0x0a, 0xe9, 0x6d, 0x70, 0x17, 0x50, 0xcf, 0x3a, 0xb2, 0xe7, 0xe0, 0x24, 0x66, 0xf7, 0x6c,
	0xf3, 0xcc, 0x2c, 0xc3, 0x5e, 0x80, 0x4b, 0xca, 0xe2, 0x75, 0x96, 0x90, 0x87, 0x98, 0x66, 0x6c,
	0x13, 0x65, 0x21, 0x7d, 0x5c, 0x07, 0x51, 0x44, 0x3d, 0x6b, 0xb6, 0xfa, 0x36, 0x80, 0xf3, 0x2e,
	0x6b, 0xf8, 0x7f, 0xbd, 0xd5, 0xf9, 0xb4, 0x2d, 0xec, 0x0b, 0x85, 0xc6, 0xab, 0x77, 0x48, 0x70,
	0xb9, 0xcd, 0x05, 0x87, 0xb2, 0xe5, 0x88, 0x97, 0x62, 0xa8, 0x3b, 0x5e, 0xa4, 0xa9, 0xd4, 0x5f,
	0x07, 0xba, 0xd1, 0xcf, 0x87, 0x39, 0xf3, 0x09, 0xf9, 0x34, 0x97, 0xbe, 0x1e, 0x46, 0x0a, 0x05,
	0x35, 0xf6, 0x94, 0xb8, 0x70, 0x58, 0xa9, 0xbe, 0x46, 0x21, 0x25, 0x85, 0x4a, 0x27, 0x21, 0x4d,
	0xdc, 0x54, 0x0b, 0x3f, 0xa6, 0xa3, 0x7f, 0x31, 0x26, 0x85, 0xc2, 0x78, 0x52, 0x30, 0x4e, 0x5c,
	0x8c, 0xb5, 0xf4, 0x76, 0x3c, 0xb4, 0xbb, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x29, 0xb5, 0xb1,
	0xb9, 0xbd, 0x01, 0x00, 0x00,
}
