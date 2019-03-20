// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/location_extension_targeting_criterion_field.proto

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

// Possible values for Location Extension Targeting criterion fields.
type LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField int32

const (
	// Not specified.
	LocationExtensionTargetingCriterionFieldEnum_UNSPECIFIED LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 0
	// Used for return value only. Represents value unknown in this version.
	LocationExtensionTargetingCriterionFieldEnum_UNKNOWN LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 1
	// Data Type: STRING. Line 1 of the business address.
	LocationExtensionTargetingCriterionFieldEnum_ADDRESS_LINE_1 LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 2
	// Data Type: STRING. Line 2 of the business address.
	LocationExtensionTargetingCriterionFieldEnum_ADDRESS_LINE_2 LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 3
	// Data Type: STRING. City of the business address.
	LocationExtensionTargetingCriterionFieldEnum_CITY LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 4
	// Data Type: STRING. Province of the business address.
	LocationExtensionTargetingCriterionFieldEnum_PROVINCE LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 5
	// Data Type: STRING. Postal code of the business address.
	LocationExtensionTargetingCriterionFieldEnum_POSTAL_CODE LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 6
	// Data Type: STRING. Country code of the business address.
	LocationExtensionTargetingCriterionFieldEnum_COUNTRY_CODE LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField = 7
)

var LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADDRESS_LINE_1",
	3: "ADDRESS_LINE_2",
	4: "CITY",
	5: "PROVINCE",
	6: "POSTAL_CODE",
	7: "COUNTRY_CODE",
}

var LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ADDRESS_LINE_1": 2,
	"ADDRESS_LINE_2": 3,
	"CITY":           4,
	"PROVINCE":       5,
	"POSTAL_CODE":    6,
	"COUNTRY_CODE":   7,
}

func (x LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField) String() string {
	return proto.EnumName(LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField_name, int32(x))
}

func (LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b18aad743d9f814, []int{0, 0}
}

// Values for Location Extension Targeting criterion fields.
type LocationExtensionTargetingCriterionFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationExtensionTargetingCriterionFieldEnum) Reset() {
	*m = LocationExtensionTargetingCriterionFieldEnum{}
}
func (m *LocationExtensionTargetingCriterionFieldEnum) String() string {
	return proto.CompactTextString(m)
}
func (*LocationExtensionTargetingCriterionFieldEnum) ProtoMessage() {}
func (*LocationExtensionTargetingCriterionFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b18aad743d9f814, []int{0}
}

func (m *LocationExtensionTargetingCriterionFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationExtensionTargetingCriterionFieldEnum.Unmarshal(m, b)
}
func (m *LocationExtensionTargetingCriterionFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationExtensionTargetingCriterionFieldEnum.Marshal(b, m, deterministic)
}
func (m *LocationExtensionTargetingCriterionFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationExtensionTargetingCriterionFieldEnum.Merge(m, src)
}
func (m *LocationExtensionTargetingCriterionFieldEnum) XXX_Size() int {
	return xxx_messageInfo_LocationExtensionTargetingCriterionFieldEnum.Size(m)
}
func (m *LocationExtensionTargetingCriterionFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationExtensionTargetingCriterionFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LocationExtensionTargetingCriterionFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField", LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField_name, LocationExtensionTargetingCriterionFieldEnum_LocationExtensionTargetingCriterionField_value)
	proto.RegisterType((*LocationExtensionTargetingCriterionFieldEnum)(nil), "google.ads.googleads.v1.enums.LocationExtensionTargetingCriterionFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/location_extension_targeting_criterion_field.proto", fileDescriptor_5b18aad743d9f814)
}

var fileDescriptor_5b18aad743d9f814 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0xae, 0x93, 0x40,
	0x18, 0x85, 0x85, 0x7b, 0xbd, 0xf7, 0x66, 0xda, 0xe8, 0x64, 0x96, 0xc6, 0x2e, 0xda, 0x95, 0x0b,
	0x1d, 0x82, 0xee, 0xc6, 0x15, 0x85, 0x69, 0x43, 0x6c, 0x80, 0x14, 0x8a, 0xa9, 0x21, 0x21, 0x58,
	0xc6, 0x09, 0x09, 0x9d, 0x69, 0x18, 0xda, 0xb8, 0xf6, 0x51, 0x5c, 0xea, 0x9b, 0xf8, 0x0c, 0x3e,
	0x81, 0x4f, 0x61, 0x06, 0x4a, 0x37, 0xc6, 0x9b, 0x6e, 0xc8, 0xc9, 0xf9, 0x67, 0xbe, 0xc3, 0x9c,
	0x1f, 0x44, 0x5c, 0x4a, 0x5e, 0x33, 0xab, 0x28, 0x95, 0xd5, 0x4b, 0xad, 0x4e, 0xb6, 0xc5, 0xc4,
	0x71, 0xaf, 0xac, 0x5a, 0xee, 0x8a, 0xb6, 0x92, 0x22, 0x67, 0x5f, 0x5b, 0x26, 0x94, 0x56, 0x6d,
	0xd1, 0x70, 0xd6, 0x56, 0x82, 0xe7, 0xbb, 0xa6, 0x6a, 0x59, 0xa3, 0xbd, 0x2f, 0x15, 0xab, 0x4b,
	0x7c, 0x68, 0x64, 0x2b, 0xd1, 0xa4, 0xc7, 0xe0, 0xa2, 0x54, 0xf8, 0x42, 0xc4, 0x27, 0x1b, 0x77,
	0xc4, 0x17, 0x2f, 0x87, 0xc0, 0x43, 0x65, 0x15, 0x42, 0xc8, 0xb6, 0xe3, 0xab, 0xfe, 0xf2, 0xec,
	0xb7, 0x01, 0x5e, 0xaf, 0xce, 0x99, 0x74, 0x88, 0x4c, 0x86, 0x44, 0x77, 0x08, 0x5c, 0xe8, 0x3c,
	0x2a, 0x8e, 0xfb, 0xd9, 0x4f, 0x03, 0xbc, 0xba, 0xf6, 0x02, 0x7a, 0x0e, 0x46, 0x9b, 0x20, 0x8e,
	0xa8, 0xeb, 0x2f, 0x7c, 0xea, 0xc1, 0x27, 0x68, 0x04, 0xee, 0x37, 0xc1, 0x87, 0x20, 0xfc, 0x18,
	0x40, 0x03, 0x21, 0xf0, 0xcc, 0xf1, 0xbc, 0x35, 0x8d, 0xe3, 0x7c, 0xe5, 0x07, 0x34, 0xb7, 0xa1,
	0xf9, 0x8f, 0xf7, 0x16, 0xde, 0xa0, 0x07, 0x70, 0xeb, 0xfa, 0xc9, 0x16, 0xde, 0xa2, 0x31, 0x78,
	0x88, 0xd6, 0x61, 0xea, 0x07, 0x2e, 0x85, 0x4f, 0x35, 0x3d, 0x0a, 0xe3, 0xc4, 0x59, 0xe5, 0x6e,
	0xe8, 0x51, 0x78, 0x87, 0x20, 0x18, 0xbb, 0xe1, 0x26, 0x48, 0xd6, 0xdb, 0xde, 0xb9, 0x9f, 0x7f,
	0x33, 0xc1, 0x74, 0x27, 0xf7, 0xf8, 0xd1, 0x8a, 0xe6, 0x6f, 0xae, 0x7d, 0x50, 0xa4, 0x3b, 0x8b,
	0x8c, 0x4f, 0xf3, 0x33, 0x8f, 0xcb, 0xba, 0x10, 0x1c, 0xcb, 0x86, 0x5b, 0x9c, 0x89, 0xae, 0xd1,
	0x61, 0xa9, 0x87, 0x4a, 0xfd, 0x67, 0xc7, 0xef, 0xbb, 0xef, 0x77, 0xf3, 0x66, 0xe9, 0x38, 0x3f,
	0xcc, 0xc9, 0xb2, 0x47, 0x39, 0xa5, 0xc2, 0xbd, 0xd4, 0x2a, 0xb5, 0xb1, 0x6e, 0x5b, 0xfd, 0x1a,
	0xe6, 0x99, 0x53, 0xaa, 0xec, 0x32, 0xcf, 0x52, 0x3b, 0xeb, 0xe6, 0x7f, 0xcc, 0x69, 0x6f, 0x12,
	0xe2, 0x94, 0x8a, 0x90, 0xcb, 0x09, 0x42, 0x52, 0x9b, 0x90, 0xee, 0xcc, 0xe7, 0xbb, 0xee, 0xc7,
	0xde, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x13, 0x06, 0x99, 0xed, 0x7b, 0x02, 0x00, 0x00,
}
