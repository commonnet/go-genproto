// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/change_status_error.proto

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

// Enum describing possible change status errors.
type ChangeStatusErrorEnum_ChangeStatusError int32

const (
	// Enum unspecified.
	ChangeStatusErrorEnum_UNSPECIFIED ChangeStatusErrorEnum_ChangeStatusError = 0
	// The received error code is not known in this version.
	ChangeStatusErrorEnum_UNKNOWN ChangeStatusErrorEnum_ChangeStatusError = 1
	// The requested start date is too old.
	ChangeStatusErrorEnum_START_DATE_TOO_OLD ChangeStatusErrorEnum_ChangeStatusError = 3
)

var ChangeStatusErrorEnum_ChangeStatusError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "START_DATE_TOO_OLD",
}

var ChangeStatusErrorEnum_ChangeStatusError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"START_DATE_TOO_OLD": 3,
}

func (x ChangeStatusErrorEnum_ChangeStatusError) String() string {
	return proto.EnumName(ChangeStatusErrorEnum_ChangeStatusError_name, int32(x))
}

func (ChangeStatusErrorEnum_ChangeStatusError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7c28c4196dfc2bd, []int{0, 0}
}

// Container for enum describing possible change status errors.
type ChangeStatusErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusErrorEnum) Reset()         { *m = ChangeStatusErrorEnum{} }
func (m *ChangeStatusErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusErrorEnum) ProtoMessage()    {}
func (*ChangeStatusErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7c28c4196dfc2bd, []int{0}
}

func (m *ChangeStatusErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusErrorEnum.Unmarshal(m, b)
}
func (m *ChangeStatusErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusErrorEnum.Marshal(b, m, deterministic)
}
func (m *ChangeStatusErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusErrorEnum.Merge(m, src)
}
func (m *ChangeStatusErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusErrorEnum.Size(m)
}
func (m *ChangeStatusErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.ChangeStatusErrorEnum_ChangeStatusError", ChangeStatusErrorEnum_ChangeStatusError_name, ChangeStatusErrorEnum_ChangeStatusError_value)
	proto.RegisterType((*ChangeStatusErrorEnum)(nil), "google.ads.googleads.v1.errors.ChangeStatusErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/change_status_error.proto", fileDescriptor_a7c28c4196dfc2bd)
}

var fileDescriptor_a7c28c4196dfc2bd = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x6a, 0xb3, 0x40,
	0x14, 0x85, 0x7f, 0x0d, 0xfc, 0x85, 0xc9, 0xa2, 0x56, 0x68, 0x16, 0xa5, 0x64, 0xe1, 0x03, 0xcc,
	0x20, 0xdd, 0x94, 0xe9, 0x6a, 0x12, 0x6d, 0x90, 0x16, 0x95, 0x6a, 0x2c, 0x14, 0x41, 0x26, 0x51,
	0xa6, 0x42, 0x32, 0x23, 0x8e, 0xc9, 0x03, 0x75, 0xd9, 0x47, 0xe9, 0xa3, 0x74, 0xd7, 0x37, 0x28,
	0xce, 0x54, 0x37, 0xa1, 0x5d, 0x79, 0xb8, 0x7e, 0xe7, 0xdc, 0x33, 0x17, 0xdc, 0x32, 0x21, 0xd8,
	0xae, 0x42, 0xb4, 0x94, 0x48, 0xcb, 0x5e, 0x1d, 0x5d, 0x54, 0xb5, 0xad, 0x68, 0x25, 0xda, 0xbe,
	0x52, 0xce, 0xaa, 0x42, 0x76, 0xb4, 0x3b, 0xc8, 0x42, 0x0d, 0x61, 0xd3, 0x8a, 0x4e, 0xd8, 0x73,
	0x8d, 0x43, 0x5a, 0x4a, 0x38, 0x3a, 0xe1, 0xd1, 0x85, 0xda, 0x79, 0x75, 0x3d, 0x24, 0x37, 0x35,
	0xa2, 0x9c, 0x8b, 0x8e, 0x76, 0xb5, 0xe0, 0x52, 0xbb, 0x9d, 0x0d, 0xb8, 0x5c, 0xaa, 0xe8, 0x44,
	0x25, 0xfb, 0xbd, 0xc7, 0xe7, 0x87, 0xbd, 0x13, 0x80, 0x8b, 0x93, 0x1f, 0xf6, 0x39, 0x98, 0xae,
	0xc3, 0x24, 0xf6, 0x97, 0xc1, 0x7d, 0xe0, 0x7b, 0xd6, 0x3f, 0x7b, 0x0a, 0xce, 0xd6, 0xe1, 0x43,
	0x18, 0x3d, 0x87, 0x96, 0x61, 0xcf, 0x80, 0x9d, 0xa4, 0xe4, 0x29, 0x2d, 0x3c, 0x92, 0xfa, 0x45,
	0x1a, 0x45, 0x45, 0xf4, 0xe8, 0x59, 0x93, 0xc5, 0x97, 0x01, 0x9c, 0xad, 0xd8, 0xc3, 0xbf, 0x8b,
	0x2e, 0x66, 0x27, 0xfb, 0xe2, 0xbe, 0x62, 0x6c, 0xbc, 0x78, 0x3f, 0x4e, 0x26, 0x76, 0x94, 0x33,
	0x28, 0x5a, 0x86, 0x58, 0xc5, 0xd5, 0x03, 0x86, 0x63, 0x35, 0xb5, 0xfc, 0xed, 0x76, 0x77, 0xfa,
	0xf3, 0x66, 0x4e, 0x56, 0x84, 0xbc, 0x9b, 0xf3, 0x95, 0x0e, 0x23, 0xa5, 0x84, 0x5a, 0xf6, 0x2a,
	0x73, 0xa1, 0x5a, 0x29, 0x3f, 0x06, 0x20, 0x27, 0xa5, 0xcc, 0x47, 0x20, 0xcf, 0xdc, 0x5c, 0x03,
	0x9f, 0xa6, 0xa3, 0xa7, 0x18, 0x93, 0x52, 0x62, 0x3c, 0x22, 0x18, 0x67, 0x2e, 0xc6, 0x1a, 0xda,
	0xfc, 0x57, 0xed, 0x6e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xa9, 0xfd, 0xb6, 0xd8, 0x01,
	0x00, 0x00,
}
