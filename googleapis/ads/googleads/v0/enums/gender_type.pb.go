// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/gender_type.proto

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

// The type of demographic genders (e.g. female).
type GenderTypeEnum_GenderType int32

const (
	// Not specified.
	GenderTypeEnum_UNSPECIFIED GenderTypeEnum_GenderType = 0
	// Used for return value only. Represents value unknown in this version.
	GenderTypeEnum_UNKNOWN GenderTypeEnum_GenderType = 1
	// Male.
	GenderTypeEnum_MALE GenderTypeEnum_GenderType = 10
	// Female.
	GenderTypeEnum_FEMALE GenderTypeEnum_GenderType = 11
	// Undetermined gender.
	GenderTypeEnum_UNDETERMINED GenderTypeEnum_GenderType = 20
)

var GenderTypeEnum_GenderType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	10: "MALE",
	11: "FEMALE",
	20: "UNDETERMINED",
}

var GenderTypeEnum_GenderType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"MALE":         10,
	"FEMALE":       11,
	"UNDETERMINED": 20,
}

func (x GenderTypeEnum_GenderType) String() string {
	return proto.EnumName(GenderTypeEnum_GenderType_name, int32(x))
}

func (GenderTypeEnum_GenderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8b913088d490c7cb, []int{0, 0}
}

// Container for enum describing the type of demographic genders.
type GenderTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenderTypeEnum) Reset()         { *m = GenderTypeEnum{} }
func (m *GenderTypeEnum) String() string { return proto.CompactTextString(m) }
func (*GenderTypeEnum) ProtoMessage()    {}
func (*GenderTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b913088d490c7cb, []int{0}
}

func (m *GenderTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenderTypeEnum.Unmarshal(m, b)
}
func (m *GenderTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenderTypeEnum.Marshal(b, m, deterministic)
}
func (m *GenderTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenderTypeEnum.Merge(m, src)
}
func (m *GenderTypeEnum) XXX_Size() int {
	return xxx_messageInfo_GenderTypeEnum.Size(m)
}
func (m *GenderTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GenderTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GenderTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.GenderTypeEnum_GenderType", GenderTypeEnum_GenderType_name, GenderTypeEnum_GenderType_value)
	proto.RegisterType((*GenderTypeEnum)(nil), "google.ads.googleads.v0.enums.GenderTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/gender_type.proto", fileDescriptor_8b913088d490c7cb)
}

var fileDescriptor_8b913088d490c7cb = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0x86, 0x32, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc, 0xd2,
	0xdc, 0x62, 0xfd, 0xf4, 0xd4, 0xbc, 0x94, 0xd4, 0xa2, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x2a, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x06, 0xbd,
	0x32, 0x03, 0x3d, 0xb0, 0x06, 0xa5, 0x14, 0x2e, 0x3e, 0x77, 0xb0, 0x9e, 0x90, 0xca, 0x82, 0x54,
	0xd7, 0xbc, 0xd2, 0x5c, 0xa5, 0x20, 0x2e, 0x2e, 0x84, 0x88, 0x10, 0x3f, 0x17, 0x77, 0xa8, 0x5f,
	0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83, 0x10, 0x37, 0x17, 0x7b, 0xa8,
	0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0xa3, 0x10, 0x07, 0x17, 0x8b, 0xaf, 0xa3, 0x8f, 0xab,
	0x00, 0x97, 0x10, 0x17, 0x17, 0x9b, 0x9b, 0x2b, 0x98, 0xcd, 0x2d, 0x24, 0xc0, 0xc5, 0x13, 0xea,
	0xe7, 0xe2, 0x1a, 0xe2, 0x1a, 0xe4, 0xeb, 0xe9, 0xe7, 0xea, 0x22, 0x20, 0xe2, 0xf4, 0x84, 0x91,
	0x4b, 0x31, 0x39, 0x3f, 0x57, 0x0f, 0xaf, 0x5b, 0x9c, 0xf8, 0x11, 0xf6, 0x06, 0x80, 0xdc, 0x1e,
	0xc0, 0x18, 0xe5, 0x04, 0xd5, 0x91, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0x0e,
	0xf2, 0x22, 0xd8, 0x67, 0x30, 0xef, 0x17, 0x64, 0x16, 0xe3, 0x08, 0x0d, 0x6b, 0x30, 0xb9, 0x88,
	0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15, 0x93, 0xac, 0x3b, 0xc4, 0x28, 0xc7, 0x94, 0x62, 0x3d, 0x08,
	0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x03, 0xf9, 0xba, 0xf8, 0x14, 0x4c, 0x3e, 0xc6, 0x31, 0xa5, 0x38,
	0x06, 0x2e, 0x1f, 0x13, 0x66, 0x10, 0x03, 0x96, 0x7f, 0xc5, 0xa4, 0x08, 0x11, 0xb4, 0xb2, 0x72,
	0x4c, 0x29, 0xb6, 0xb2, 0x82, 0xab, 0xb0, 0xb2, 0x0a, 0x33, 0xb0, 0xb2, 0x02, 0xab, 0x49, 0x62,
	0x03, 0x3b, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x23, 0x1d, 0x86, 0x6f, 0xa5, 0x01, 0x00,
	0x00,
}
