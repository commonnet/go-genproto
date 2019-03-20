// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/keyword_plan_campaign_error.proto

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

// Enum describing possible errors from applying a keyword plan campaign.
type KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError int32

const (
	// Enum unspecified.
	KeywordPlanCampaignErrorEnum_UNSPECIFIED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 0
	// The received error code is not known in this version.
	KeywordPlanCampaignErrorEnum_UNKNOWN KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 1
	// A keyword plan campaign name is missing, empty, longer than allowed limit
	// or contains invalid chars.
	KeywordPlanCampaignErrorEnum_INVALID_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 2
	// A keyword plan campaign contains one or more untargetable languages.
	KeywordPlanCampaignErrorEnum_INVALID_LANGUAGES KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 3
	// A keyword plan campaign contains one or more invalid geo targets.
	KeywordPlanCampaignErrorEnum_INVALID_GEOS KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 4
	// The keyword plan campaign name is duplicate to an existing keyword plan
	// campaign name or other keyword plan campaign name in the request.
	KeywordPlanCampaignErrorEnum_DUPLICATE_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 5
	// The number of geo targets in the keyword plan campaign exceeds limits.
	KeywordPlanCampaignErrorEnum_MAX_GEOS_EXCEEDED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 6
)

var KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_NAME",
	3: "INVALID_LANGUAGES",
	4: "INVALID_GEOS",
	5: "DUPLICATE_NAME",
	6: "MAX_GEOS_EXCEEDED",
}

var KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"INVALID_NAME":      2,
	"INVALID_LANGUAGES": 3,
	"INVALID_GEOS":      4,
	"DUPLICATE_NAME":    5,
	"MAX_GEOS_EXCEEDED": 6,
}

func (x KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) String() string {
	return proto.EnumName(KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name, int32(x))
}

func (KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ae42b98134de2d2, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// campaign.
type KeywordPlanCampaignErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanCampaignErrorEnum) Reset()         { *m = KeywordPlanCampaignErrorEnum{} }
func (m *KeywordPlanCampaignErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaignErrorEnum) ProtoMessage()    {}
func (*KeywordPlanCampaignErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ae42b98134de2d2, []int{0}
}

func (m *KeywordPlanCampaignErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaignErrorEnum.Merge(m, src)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Size(m)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaignErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaignErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError", KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name, KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value)
	proto.RegisterType((*KeywordPlanCampaignErrorEnum)(nil), "google.ads.googleads.v0.errors.KeywordPlanCampaignErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/keyword_plan_campaign_error.proto", fileDescriptor_9ae42b98134de2d2)
}

var fileDescriptor_9ae42b98134de2d2 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xed, 0xfe, 0x4e, 0xc8, 0x44, 0x6b, 0x41, 0xf0, 0xa0, 0x3b, 0xf4, 0x03, 0xa4,
	0x05, 0x6f, 0xf1, 0x62, 0xd6, 0xc6, 0x52, 0xb6, 0x65, 0x85, 0xd9, 0x3a, 0xa4, 0x50, 0xe2, 0x5a,
	0xc2, 0xb0, 0x6b, 0x4a, 0xa3, 0x13, 0xbf, 0x8c, 0x07, 0x8f, 0x7e, 0x09, 0xef, 0x7e, 0x14, 0x3f,
	0x81, 0x47, 0x69, 0xb3, 0x0d, 0x2f, 0xf3, 0x94, 0x87, 0xf7, 0xfd, 0xbd, 0x4f, 0xf2, 0x3e, 0x01,
	0x57, 0x5c, 0x08, 0x5e, 0xe4, 0x36, 0xcb, 0xa4, 0xad, 0x64, 0xa3, 0x56, 0x8e, 0x9d, 0xd7, 0xb5,
	0xa8, 0xa5, 0xfd, 0x90, 0xbf, 0x3c, 0x8b, 0x3a, 0x4b, 0xab, 0x82, 0x95, 0xe9, 0x9c, 0x2d, 0x2b,
	0xb6, 0xe0, 0x65, 0xda, 0x36, 0x61, 0x55, 0x8b, 0x47, 0x61, 0xf6, 0xd5, 0x18, 0x64, 0x99, 0x84,
	0x5b, 0x07, 0xb8, 0x72, 0xa0, 0x72, 0xb0, 0x3e, 0x34, 0x70, 0x36, 0x54, 0x2e, 0x61, 0xc1, 0x4a,
	0x77, 0xed, 0x41, 0x9a, 0x2e, 0x29, 0x9f, 0x96, 0xd6, 0xab, 0x06, 0x4e, 0x77, 0x01, 0xe6, 0x11,
	0xe8, 0x45, 0x74, 0x1a, 0x12, 0x37, 0xb8, 0x0e, 0x88, 0x67, 0xfc, 0x33, 0x7b, 0x60, 0x3f, 0xa2,
	0x43, 0x3a, 0xb9, 0xa5, 0x86, 0x66, 0x1a, 0xe0, 0x20, 0xa0, 0x31, 0x1e, 0x05, 0x5e, 0x4a, 0xf1,
	0x98, 0x18, 0xba, 0x79, 0x02, 0x8e, 0x37, 0x95, 0x11, 0xa6, 0x7e, 0x84, 0x7d, 0x32, 0x35, 0x3a,
	0xbf, 0x41, 0x9f, 0x4c, 0xa6, 0xc6, 0x7f, 0xd3, 0x04, 0x87, 0x5e, 0x14, 0x8e, 0x02, 0x17, 0xdf,
	0x10, 0x35, 0xbc, 0xd7, 0x0c, 0x8f, 0xf1, 0xac, 0x25, 0x52, 0x32, 0x73, 0x09, 0xf1, 0x88, 0x67,
	0x74, 0x07, 0xdf, 0x1a, 0xb0, 0xe6, 0x62, 0x09, 0xff, 0x5e, 0x74, 0x70, 0xbe, 0x6b, 0x89, 0xb0,
	0xc9, 0x29, 0xd4, 0xee, 0xbc, 0xb5, 0x01, 0x17, 0x05, 0x2b, 0x39, 0x14, 0x35, 0xb7, 0x79, 0x5e,
	0xb6, 0x29, 0x6e, 0xb2, 0xaf, 0x16, 0x72, 0xd7, 0x57, 0x5c, 0xaa, 0xe3, 0x4d, 0xef, 0xf8, 0x18,
	0xbf, 0xeb, 0x7d, 0x5f, 0x99, 0xe1, 0x4c, 0x42, 0x25, 0x1b, 0x15, 0x3b, 0xb0, 0xbd, 0x52, 0x7e,
	0x6e, 0x80, 0x04, 0x67, 0x32, 0xd9, 0x02, 0x49, 0xec, 0x24, 0x0a, 0xf8, 0xd2, 0x2d, 0x55, 0x45,
	0x08, 0x67, 0x12, 0xa1, 0x2d, 0x82, 0x50, 0xec, 0x20, 0xa4, 0xa0, 0xfb, 0x6e, 0xfb, 0xba, 0x8b,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x7c, 0xec, 0x11, 0x27, 0x02, 0x00, 0x00,
}
