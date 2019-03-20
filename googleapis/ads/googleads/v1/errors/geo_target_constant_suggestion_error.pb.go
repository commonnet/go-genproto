// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/geo_target_constant_suggestion_error.proto

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

// Enum describing possible geo target constant suggestion errors.
type GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError int32

const (
	// Enum unspecified.
	GeoTargetConstantSuggestionErrorEnum_UNSPECIFIED GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 0
	// The received error code is not known in this version.
	GeoTargetConstantSuggestionErrorEnum_UNKNOWN GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 1
	// A location name cannot be greater than 300 characters.
	GeoTargetConstantSuggestionErrorEnum_LOCATION_NAME_SIZE_LIMIT GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 2
	// At most 25 location names can be specified in a SuggestGeoTargetConstants
	// method.
	GeoTargetConstantSuggestionErrorEnum_LOCATION_NAME_LIMIT GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 3
	// The country code is invalid.
	GeoTargetConstantSuggestionErrorEnum_INVALID_COUNTRY_CODE GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 4
	// Geo target constant resource names or location names must be provided in
	// the request.
	GeoTargetConstantSuggestionErrorEnum_REQUEST_PARAMETERS_UNSET GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 5
)

var GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LOCATION_NAME_SIZE_LIMIT",
	3: "LOCATION_NAME_LIMIT",
	4: "INVALID_COUNTRY_CODE",
	5: "REQUEST_PARAMETERS_UNSET",
}

var GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"LOCATION_NAME_SIZE_LIMIT": 2,
	"LOCATION_NAME_LIMIT":      3,
	"INVALID_COUNTRY_CODE":     4,
	"REQUEST_PARAMETERS_UNSET": 5,
}

func (x GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError) String() string {
	return proto.EnumName(GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_name, int32(x))
}

func (GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef7e3347ae4bfc6d, []int{0, 0}
}

// Container for enum describing possible geo target constant suggestion errors.
type GeoTargetConstantSuggestionErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoTargetConstantSuggestionErrorEnum) Reset()         { *m = GeoTargetConstantSuggestionErrorEnum{} }
func (m *GeoTargetConstantSuggestionErrorEnum) String() string { return proto.CompactTextString(m) }
func (*GeoTargetConstantSuggestionErrorEnum) ProtoMessage()    {}
func (*GeoTargetConstantSuggestionErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef7e3347ae4bfc6d, []int{0}
}

func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Unmarshal(m, b)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Marshal(b, m, deterministic)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Merge(m, src)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Size() int {
	return xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Size(m)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError", GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_name, GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_value)
	proto.RegisterType((*GeoTargetConstantSuggestionErrorEnum)(nil), "google.ads.googleads.v1.errors.GeoTargetConstantSuggestionErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/geo_target_constant_suggestion_error.proto", fileDescriptor_ef7e3347ae4bfc6d)
}

var fileDescriptor_ef7e3347ae4bfc6d = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0xaa, 0xd4, 0x30,
	0x14, 0xc6, 0x6d, 0xaf, 0x7f, 0x20, 0x77, 0x61, 0xa9, 0x82, 0x17, 0xb9, 0x0c, 0x52, 0x74, 0x9b,
	0x52, 0xdc, 0xc5, 0x55, 0xa6, 0x8d, 0x43, 0x70, 0x26, 0xad, 0xfd, 0x27, 0x0e, 0x85, 0x50, 0xa7,
	0x25, 0x0c, 0xcc, 0x24, 0x43, 0xd3, 0x99, 0xb5, 0xcf, 0xe2, 0xd2, 0x8d, 0xef, 0xe1, 0x13, 0xf8,
	0x0c, 0x3e, 0x85, 0xb4, 0x99, 0x29, 0xb8, 0xd0, 0x59, 0xe5, 0x23, 0xe7, 0x3b, 0xbf, 0x2f, 0x39,
	0x07, 0x50, 0xa1, 0x94, 0xd8, 0xb5, 0x7e, 0xdd, 0x68, 0xdf, 0xc8, 0x41, 0x9d, 0x02, 0xbf, 0xed,
	0x3a, 0xd5, 0x69, 0x5f, 0xb4, 0x8a, 0xf7, 0x75, 0x27, 0xda, 0x9e, 0x6f, 0x94, 0xd4, 0x7d, 0x2d,
	0x7b, 0xae, 0x8f, 0x42, 0xb4, 0xba, 0xdf, 0x2a, 0xc9, 0x47, 0x17, 0x3c, 0x74, 0xaa, 0x57, 0xee,
	0xcc, 0xf4, 0xc3, 0xba, 0xd1, 0x70, 0x42, 0xc1, 0x53, 0x00, 0x0d, 0xea, 0xe5, 0xfd, 0x25, 0xea,
	0xb0, 0xf5, 0x6b, 0x29, 0x55, 0x5f, 0x0f, 0x08, 0x6d, 0xba, 0xbd, 0x5f, 0x16, 0x78, 0xbd, 0x68,
	0x55, 0x3e, 0x66, 0x85, 0xe7, 0xa8, 0x6c, 0x4a, 0x22, 0x03, 0x83, 0xc8, 0xe3, 0xde, 0xfb, 0x61,
	0x81, 0x57, 0xd7, 0x8c, 0xee, 0x53, 0x70, 0x5b, 0xb0, 0x2c, 0x21, 0x21, 0x7d, 0x4f, 0x49, 0xe4,
	0x3c, 0x70, 0x6f, 0xc1, 0x93, 0x82, 0x7d, 0x60, 0xf1, 0x27, 0xe6, 0x58, 0xee, 0x3d, 0xb8, 0x5b,
	0xc6, 0x21, 0xce, 0x69, 0xcc, 0x38, 0xc3, 0x2b, 0xc2, 0x33, 0xba, 0x26, 0x7c, 0x49, 0x57, 0x34,
	0x77, 0x6c, 0xf7, 0x05, 0x78, 0xf6, 0x77, 0xd5, 0x14, 0x6e, 0xdc, 0x3b, 0xf0, 0x9c, 0xb2, 0x12,
	0x2f, 0x69, 0xc4, 0xc3, 0xb8, 0x60, 0x79, 0xfa, 0x99, 0x87, 0x71, 0x44, 0x9c, 0x87, 0x03, 0x30,
	0x25, 0x1f, 0x0b, 0x92, 0xe5, 0x3c, 0xc1, 0x29, 0x5e, 0x91, 0x9c, 0xa4, 0x19, 0x2f, 0x58, 0x46,
	0x72, 0xe7, 0xd1, 0xfc, 0xab, 0x0d, 0xbc, 0x8d, 0xda, 0xc3, 0xff, 0xcf, 0x67, 0xfe, 0xe6, 0xda,
	0xaf, 0x92, 0x61, 0x50, 0x89, 0xb5, 0x8e, 0xce, 0x20, 0xa1, 0x76, 0xb5, 0x14, 0x50, 0x75, 0xc2,
	0x17, 0xad, 0x1c, 0xc7, 0x78, 0xd9, 0xe1, 0x61, 0xab, 0xff, 0xb5, 0xd2, 0x77, 0xe6, 0xf8, 0x66,
	0xdf, 0x2c, 0x30, 0xfe, 0x6e, 0xcf, 0x16, 0x06, 0x86, 0x1b, 0x0d, 0x8d, 0x1c, 0x54, 0x19, 0xc0,
	0x31, 0x52, 0xff, 0xbc, 0x18, 0x2a, 0xdc, 0xe8, 0x6a, 0x32, 0x54, 0x65, 0x50, 0x19, 0xc3, 0x6f,
	0xdb, 0x33, 0xb7, 0x08, 0xe1, 0x46, 0x23, 0x34, 0x59, 0x10, 0x2a, 0x03, 0x84, 0x8c, 0xe9, 0xcb,
	0xe3, 0xf1, 0x75, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x82, 0x5e, 0xcb, 0x80, 0x6f, 0x02,
	0x00, 0x00,
}
