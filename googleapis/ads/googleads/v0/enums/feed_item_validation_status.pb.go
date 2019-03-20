// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_item_validation_status.proto

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

// The possible validation statuses of a feed item.
type FeedItemValidationStatusEnum_FeedItemValidationStatus int32

const (
	// No value has been specified.
	FeedItemValidationStatusEnum_UNSPECIFIED FeedItemValidationStatusEnum_FeedItemValidationStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemValidationStatusEnum_UNKNOWN FeedItemValidationStatusEnum_FeedItemValidationStatus = 1
	// Validation pending.
	FeedItemValidationStatusEnum_PENDING FeedItemValidationStatusEnum_FeedItemValidationStatus = 2
	// An error was found.
	FeedItemValidationStatusEnum_INVALID FeedItemValidationStatusEnum_FeedItemValidationStatus = 3
	// Feed item is semantically well-formed.
	FeedItemValidationStatusEnum_VALID FeedItemValidationStatusEnum_FeedItemValidationStatus = 4
)

var FeedItemValidationStatusEnum_FeedItemValidationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "INVALID",
	4: "VALID",
}

var FeedItemValidationStatusEnum_FeedItemValidationStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"INVALID":     3,
	"VALID":       4,
}

func (x FeedItemValidationStatusEnum_FeedItemValidationStatus) String() string {
	return proto.EnumName(FeedItemValidationStatusEnum_FeedItemValidationStatus_name, int32(x))
}

func (FeedItemValidationStatusEnum_FeedItemValidationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90b408c2fc9af860, []int{0, 0}
}

// Container for enum describing possible validation statuses of a feed item.
type FeedItemValidationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemValidationStatusEnum) Reset()         { *m = FeedItemValidationStatusEnum{} }
func (m *FeedItemValidationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemValidationStatusEnum) ProtoMessage()    {}
func (*FeedItemValidationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b408c2fc9af860, []int{0}
}

func (m *FeedItemValidationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemValidationStatusEnum.Unmarshal(m, b)
}
func (m *FeedItemValidationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemValidationStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemValidationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemValidationStatusEnum.Merge(m, src)
}
func (m *FeedItemValidationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemValidationStatusEnum.Size(m)
}
func (m *FeedItemValidationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemValidationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemValidationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedItemValidationStatusEnum_FeedItemValidationStatus", FeedItemValidationStatusEnum_FeedItemValidationStatus_name, FeedItemValidationStatusEnum_FeedItemValidationStatus_value)
	proto.RegisterType((*FeedItemValidationStatusEnum)(nil), "google.ads.googleads.v0.enums.FeedItemValidationStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_item_validation_status.proto", fileDescriptor_90b408c2fc9af860)
}

var fileDescriptor_90b408c2fc9af860 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xff, 0xa6, 0x3f, 0x20, 0xdc, 0x81, 0x28, 0x13, 0x03, 0x1d, 0xda, 0x07, 0x70, 0x22,
	0xb1, 0x99, 0x01, 0xb9, 0x34, 0xad, 0x2c, 0x90, 0x89, 0x54, 0x35, 0x48, 0x28, 0xa8, 0x32, 0xd8,
	0x58, 0x91, 0x9a, 0xb8, 0xaa, 0x9d, 0x6c, 0xbc, 0x0c, 0x23, 0x8f, 0xc2, 0xa3, 0x30, 0xf2, 0x04,
	0xc8, 0x36, 0xcd, 0x16, 0x96, 0xe8, 0xdc, 0x9c, 0x7b, 0x3f, 0x1d, 0x1f, 0x70, 0x2d, 0x95, 0x92,
	0x5b, 0x11, 0x33, 0xae, 0x63, 0x2f, 0xad, 0x6a, 0x93, 0x58, 0xd4, 0x4d, 0xa5, 0xe3, 0x57, 0x21,
	0xf8, 0xa6, 0x34, 0xa2, 0xda, 0xb4, 0x6c, 0x5b, 0x72, 0x66, 0x4a, 0x55, 0x6f, 0xb4, 0x61, 0xa6,
	0xd1, 0x70, 0xb7, 0x57, 0x46, 0x45, 0x63, 0x7f, 0x05, 0x19, 0xd7, 0xb0, 0x03, 0xc0, 0x36, 0x81,
	0x0e, 0x30, 0x7d, 0x03, 0x17, 0x0b, 0x21, 0x38, 0x31, 0xa2, 0xca, 0x3b, 0xc2, 0xca, 0x01, 0xd2,
	0xba, 0xa9, 0xa6, 0x4f, 0xe0, 0xbc, 0xcf, 0x8f, 0xce, 0xc0, 0x68, 0x4d, 0x57, 0x59, 0x7a, 0x43,
	0x16, 0x24, 0x9d, 0x87, 0xff, 0xa2, 0x11, 0x38, 0x59, 0xd3, 0x5b, 0x7a, 0xff, 0x40, 0xc3, 0x81,
	0x1d, 0xb2, 0x94, 0xce, 0x09, 0x5d, 0x86, 0x81, 0x1d, 0x08, 0xcd, 0xf1, 0x1d, 0x99, 0x87, 0xc3,
	0xe8, 0x14, 0x1c, 0x79, 0xf9, 0x7f, 0xf6, 0x3d, 0x00, 0x93, 0x17, 0x55, 0xc1, 0x3f, 0x43, 0xce,
	0xc6, 0x7d, 0x11, 0x32, 0xfb, 0xc4, 0x6c, 0xf0, 0x38, 0xfb, 0xbd, 0x97, 0x6a, 0xcb, 0x6a, 0x09,
	0xd5, 0x5e, 0xc6, 0x52, 0xd4, 0xae, 0x80, 0x43, 0x6b, 0xbb, 0x52, 0xf7, 0x94, 0x78, 0xe5, 0xbe,
	0xef, 0xc1, 0x70, 0x89, 0xf1, 0x47, 0x30, 0x5e, 0x7a, 0x14, 0xe6, 0x1a, 0x7a, 0x69, 0x55, 0x9e,
	0x40, 0x5b, 0x87, 0xfe, 0x3c, 0xf8, 0x05, 0xe6, 0xba, 0xe8, 0xfc, 0x22, 0x4f, 0x0a, 0xe7, 0x7f,
	0x05, 0x13, 0xff, 0x13, 0x21, 0xcc, 0x35, 0x42, 0xdd, 0x06, 0x42, 0x79, 0x82, 0x90, 0xdb, 0x79,
	0x3e, 0x76, 0xc1, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xa0, 0x37, 0xc9, 0xdc, 0x01,
	0x00, 0x00,
}
