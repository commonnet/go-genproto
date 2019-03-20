// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/shared_set_error.proto

package errors

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

// Enum describing possible shared set errors.
type SharedSetErrorEnum_SharedSetError int32

const (
	// Enum unspecified.
	SharedSetErrorEnum_UNSPECIFIED SharedSetErrorEnum_SharedSetError = 0
	// The received error code is not known in this version.
	SharedSetErrorEnum_UNKNOWN SharedSetErrorEnum_SharedSetError = 1
	// The customer cannot create this type of shared set.
	SharedSetErrorEnum_CUSTOMER_CANNOT_CREATE_SHARED_SET_OF_THIS_TYPE SharedSetErrorEnum_SharedSetError = 2
	// A shared set with this name already exists.
	SharedSetErrorEnum_DUPLICATE_NAME SharedSetErrorEnum_SharedSetError = 3
	// Removed shared sets cannot be mutated.
	SharedSetErrorEnum_SHARED_SET_REMOVED SharedSetErrorEnum_SharedSetError = 4
	// The shared set cannot be removed because it is in use.
	SharedSetErrorEnum_SHARED_SET_IN_USE SharedSetErrorEnum_SharedSetError = 5
)

var SharedSetErrorEnum_SharedSetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CUSTOMER_CANNOT_CREATE_SHARED_SET_OF_THIS_TYPE",
	3: "DUPLICATE_NAME",
	4: "SHARED_SET_REMOVED",
	5: "SHARED_SET_IN_USE",
}

var SharedSetErrorEnum_SharedSetError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CUSTOMER_CANNOT_CREATE_SHARED_SET_OF_THIS_TYPE": 2,
	"DUPLICATE_NAME":     3,
	"SHARED_SET_REMOVED": 4,
	"SHARED_SET_IN_USE":  5,
}

func (x SharedSetErrorEnum_SharedSetError) String() string {
	return proto.EnumName(SharedSetErrorEnum_SharedSetError_name, int32(x))
}

func (SharedSetErrorEnum_SharedSetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6111e246eb270151, []int{0, 0}
}

// Container for enum describing possible shared set errors.
type SharedSetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSetErrorEnum) Reset()         { *m = SharedSetErrorEnum{} }
func (m *SharedSetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*SharedSetErrorEnum) ProtoMessage()    {}
func (*SharedSetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6111e246eb270151, []int{0}
}

func (m *SharedSetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetErrorEnum.Unmarshal(m, b)
}
func (m *SharedSetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetErrorEnum.Marshal(b, m, deterministic)
}
func (m *SharedSetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetErrorEnum.Merge(m, src)
}
func (m *SharedSetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_SharedSetErrorEnum.Size(m)
}
func (m *SharedSetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.SharedSetErrorEnum_SharedSetError", SharedSetErrorEnum_SharedSetError_name, SharedSetErrorEnum_SharedSetError_value)
	proto.RegisterType((*SharedSetErrorEnum)(nil), "google.ads.googleads.v0.errors.SharedSetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/shared_set_error.proto", fileDescriptor_6111e246eb270151)
}

var fileDescriptor_6111e246eb270151 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x4e, 0xe3, 0x30,
	0x00, 0x85, 0x27, 0xe9, 0xfc, 0x48, 0xae, 0xd4, 0xc9, 0x78, 0x34, 0xb3, 0xec, 0x22, 0x07, 0x70,
	0x22, 0x10, 0x1b, 0xb3, 0x72, 0x13, 0xb7, 0x8d, 0xa0, 0x4e, 0x14, 0x27, 0x41, 0xa0, 0x48, 0x56,
	0x20, 0x51, 0x40, 0x6a, 0xeb, 0x2a, 0x2e, 0x3d, 0x10, 0x12, 0x1b, 0xce, 0xc0, 0x09, 0x38, 0x0a,
	0x0b, 0xce, 0x80, 0x12, 0xd3, 0xaa, 0x5d, 0xc0, 0xca, 0x4f, 0x4f, 0xdf, 0xb3, 0xfd, 0x1e, 0x38,
	0xa9, 0xa5, 0xac, 0xe7, 0x95, 0x53, 0x94, 0xca, 0xd1, 0xb2, 0x55, 0x1b, 0xd7, 0xa9, 0x9a, 0x46,
	0x36, 0xca, 0x51, 0xb7, 0x45, 0x53, 0x95, 0x42, 0x55, 0x6b, 0xd1, 0x39, 0x68, 0xd5, 0xc8, 0xb5,
	0x84, 0x43, 0xcd, 0xa2, 0xa2, 0x54, 0x68, 0x17, 0x43, 0x1b, 0x17, 0xe9, 0x98, 0xfd, 0x6c, 0x00,
	0xc8, 0xbb, 0x28, 0xaf, 0xd6, 0xb4, 0xf5, 0xe8, 0xf2, 0x7e, 0x61, 0x3f, 0x1a, 0x60, 0x70, 0x68,
	0xc3, 0xdf, 0xa0, 0x9f, 0x32, 0x1e, 0x51, 0x2f, 0x18, 0x07, 0xd4, 0xb7, 0xbe, 0xc1, 0x3e, 0xf8,
	0x95, 0xb2, 0x33, 0x16, 0x5e, 0x30, 0xcb, 0x80, 0x47, 0x00, 0x79, 0x29, 0x4f, 0xc2, 0x19, 0x8d,
	0x85, 0x47, 0x18, 0x0b, 0x13, 0xe1, 0xc5, 0x94, 0x24, 0x54, 0xf0, 0x29, 0x89, 0xa9, 0x2f, 0x38,
	0x4d, 0x44, 0x38, 0x16, 0xc9, 0x34, 0xe0, 0x22, 0xb9, 0x8c, 0xa8, 0x65, 0x42, 0x08, 0x06, 0x7e,
	0x1a, 0x9d, 0x07, 0x5e, 0x8b, 0x31, 0x32, 0xa3, 0x56, 0x0f, 0xfe, 0x07, 0x70, 0x2f, 0x10, 0xd3,
	0x59, 0x98, 0x51, 0xdf, 0xfa, 0x0e, 0xff, 0x81, 0x3f, 0x7b, 0x7e, 0xc0, 0x44, 0xca, 0xa9, 0xf5,
	0x63, 0xf4, 0x66, 0x00, 0xfb, 0x46, 0x2e, 0xd0, 0xd7, 0x2d, 0x47, 0x7f, 0x0f, 0xbb, 0x44, 0xed,
	0x34, 0x91, 0x71, 0xe5, 0x7f, 0xc4, 0x6a, 0x39, 0x2f, 0x96, 0x35, 0x92, 0x4d, 0xed, 0xd4, 0xd5,
	0xb2, 0x1b, 0x6e, 0xbb, 0xf1, 0xea, 0x4e, 0x7d, 0x36, 0xf9, 0xa9, 0x3e, 0x1e, 0xcc, 0xde, 0x84,
	0x90, 0x27, 0x73, 0x38, 0xd1, 0x97, 0x91, 0x52, 0x21, 0x2d, 0x5b, 0x95, 0xb9, 0xa8, 0x7b, 0x52,
	0xbd, 0x6c, 0x81, 0x9c, 0x94, 0x2a, 0xdf, 0x01, 0x79, 0xe6, 0xe6, 0x1a, 0x78, 0x35, 0x6d, 0xed,
	0x62, 0x4c, 0x4a, 0x85, 0xf1, 0x0e, 0xc1, 0x38, 0x73, 0x31, 0xd6, 0xd0, 0xf5, 0xcf, 0xee, 0x77,
	0xc7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xcf, 0xc3, 0x65, 0x0f, 0x02, 0x00, 0x00,
}
