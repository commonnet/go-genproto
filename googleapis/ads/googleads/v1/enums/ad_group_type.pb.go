// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/ad_group_type.proto

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

// Enum listing the possible types of an ad group.
type AdGroupTypeEnum_AdGroupType int32

const (
	// The type has not been specified.
	AdGroupTypeEnum_UNSPECIFIED AdGroupTypeEnum_AdGroupType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupTypeEnum_UNKNOWN AdGroupTypeEnum_AdGroupType = 1
	// The default ad group type for Search campaigns.
	AdGroupTypeEnum_SEARCH_STANDARD AdGroupTypeEnum_AdGroupType = 2
	// The default ad group type for Display campaigns.
	AdGroupTypeEnum_DISPLAY_STANDARD AdGroupTypeEnum_AdGroupType = 3
	// The ad group type for Shopping campaigns serving standard product ads.
	AdGroupTypeEnum_SHOPPING_PRODUCT_ADS AdGroupTypeEnum_AdGroupType = 4
	// The default ad group type for Hotel campaigns.
	AdGroupTypeEnum_HOTEL_ADS AdGroupTypeEnum_AdGroupType = 6
	// The type for ad groups in Smart Shopping campaigns.
	AdGroupTypeEnum_SHOPPING_SMART_ADS AdGroupTypeEnum_AdGroupType = 7
	// Short unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_BUMPER AdGroupTypeEnum_AdGroupType = 8
	// TrueView (skippable) in-stream video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_STREAM AdGroupTypeEnum_AdGroupType = 9
	// TrueView in-display video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_DISPLAY AdGroupTypeEnum_AdGroupType = 10
	// Unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_NON_SKIPPABLE_IN_STREAM AdGroupTypeEnum_AdGroupType = 11
	// Outstream video ads.
	AdGroupTypeEnum_VIDEO_OUTSTREAM AdGroupTypeEnum_AdGroupType = 12
)

var AdGroupTypeEnum_AdGroupType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_STANDARD",
	3:  "DISPLAY_STANDARD",
	4:  "SHOPPING_PRODUCT_ADS",
	6:  "HOTEL_ADS",
	7:  "SHOPPING_SMART_ADS",
	8:  "VIDEO_BUMPER",
	9:  "VIDEO_TRUE_VIEW_IN_STREAM",
	10: "VIDEO_TRUE_VIEW_IN_DISPLAY",
	11: "VIDEO_NON_SKIPPABLE_IN_STREAM",
	12: "VIDEO_OUTSTREAM",
}

var AdGroupTypeEnum_AdGroupType_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"SEARCH_STANDARD":               2,
	"DISPLAY_STANDARD":              3,
	"SHOPPING_PRODUCT_ADS":          4,
	"HOTEL_ADS":                     6,
	"SHOPPING_SMART_ADS":            7,
	"VIDEO_BUMPER":                  8,
	"VIDEO_TRUE_VIEW_IN_STREAM":     9,
	"VIDEO_TRUE_VIEW_IN_DISPLAY":    10,
	"VIDEO_NON_SKIPPABLE_IN_STREAM": 11,
	"VIDEO_OUTSTREAM":               12,
}

func (x AdGroupTypeEnum_AdGroupType) String() string {
	return proto.EnumName(AdGroupTypeEnum_AdGroupType_name, int32(x))
}

func (AdGroupTypeEnum_AdGroupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_07bc19b8ed4d56ed, []int{0, 0}
}

// Defines types of an ad group, specific to a particular campaign channel
// type. This type drives validations that restrict which entities can be
// added to the ad group.
type AdGroupTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupTypeEnum) Reset()         { *m = AdGroupTypeEnum{} }
func (m *AdGroupTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupTypeEnum) ProtoMessage()    {}
func (*AdGroupTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_07bc19b8ed4d56ed, []int{0}
}

func (m *AdGroupTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupTypeEnum.Unmarshal(m, b)
}
func (m *AdGroupTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupTypeEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupTypeEnum.Merge(m, src)
}
func (m *AdGroupTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupTypeEnum.Size(m)
}
func (m *AdGroupTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdGroupTypeEnum_AdGroupType", AdGroupTypeEnum_AdGroupType_name, AdGroupTypeEnum_AdGroupType_value)
	proto.RegisterType((*AdGroupTypeEnum)(nil), "google.ads.googleads.v1.enums.AdGroupTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/ad_group_type.proto", fileDescriptor_07bc19b8ed4d56ed)
}

var fileDescriptor_07bc19b8ed4d56ed = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0xa7, 0x19, 0xda, 0x98, 0x3b, 0x54, 0xcb, 0x4c, 0x08, 0x26, 0x8a, 0xd4, 0x3d, 0x80, 0xa3,
	0x88, 0x5b, 0x38, 0x39, 0x8d, 0x69, 0xad, 0xb5, 0x8e, 0x15, 0x27, 0x99, 0x40, 0x95, 0xac, 0x40,
	0xa2, 0xa8, 0xd2, 0x1a, 0x47, 0x75, 0x3b, 0x69, 0x6f, 0x83, 0x38, 0x72, 0xe6, 0x29, 0x78, 0x0f,
	0x2e, 0x3c, 0x05, 0x4a, 0xdc, 0x96, 0x1e, 0x60, 0x97, 0xe8, 0xcb, 0xef, 0x9f, 0x3e, 0x7f, 0x3f,
	0xe0, 0x55, 0x5a, 0x57, 0x77, 0xa5, 0x9b, 0x17, 0xc6, 0xb5, 0x63, 0x3b, 0xdd, 0x7b, 0x6e, 0x59,
	0x6f, 0x57, 0xc6, 0xcd, 0x0b, 0x55, 0xad, 0xf5, 0xb6, 0x51, 0x9b, 0x87, 0xa6, 0xc4, 0xcd, 0x5a,
	0x6f, 0x34, 0x1a, 0x5a, 0x1d, 0xce, 0x0b, 0x83, 0x0f, 0x16, 0x7c, 0xef, 0xe1, 0xce, 0x72, 0xf5,
	0x66, 0x9f, 0xd8, 0x2c, 0xdd, 0xbc, 0xae, 0xf5, 0x26, 0xdf, 0x2c, 0x75, 0x6d, 0xac, 0xf9, 0xfa,
	0x87, 0x03, 0x06, 0xa4, 0x98, 0xb4, 0x99, 0xc9, 0x43, 0x53, 0xd2, 0x7a, 0xbb, 0xba, 0xfe, 0xea,
	0x80, 0xfe, 0x11, 0x86, 0x06, 0xa0, 0x9f, 0x72, 0x29, 0xe8, 0x98, 0x7d, 0x60, 0x34, 0x84, 0x4f,
	0x50, 0x1f, 0x9c, 0xa5, 0xfc, 0x86, 0x47, 0xb7, 0x1c, 0xf6, 0xd0, 0x0b, 0x30, 0x90, 0x94, 0xc4,
	0xe3, 0xa9, 0x92, 0x09, 0xe1, 0x21, 0x89, 0x43, 0xe8, 0xa0, 0x4b, 0x00, 0x43, 0x26, 0xc5, 0x8c,
	0x7c, 0xfc, 0x8b, 0x9e, 0xa0, 0x57, 0xe0, 0x52, 0x4e, 0x23, 0x21, 0x18, 0x9f, 0x28, 0x11, 0x47,
	0x61, 0x3a, 0x4e, 0x14, 0x09, 0x25, 0x7c, 0x8a, 0x9e, 0x83, 0xf3, 0x69, 0x94, 0xd0, 0x59, 0xf7,
	0x7b, 0x8a, 0x5e, 0x02, 0x74, 0x10, 0xca, 0x39, 0x89, 0xad, 0xec, 0x0c, 0x41, 0x70, 0x91, 0xb1,
	0x90, 0x46, 0x2a, 0x48, 0xe7, 0x82, 0xc6, 0xf0, 0x19, 0x1a, 0x82, 0xd7, 0x16, 0x49, 0xe2, 0x94,
	0xaa, 0x8c, 0xd1, 0x5b, 0xc5, 0xb8, 0x92, 0x49, 0x4c, 0xc9, 0x1c, 0x9e, 0xa3, 0xb7, 0xe0, 0xea,
	0x1f, 0xf4, 0x6e, 0x35, 0x08, 0xd0, 0x08, 0x0c, 0x2d, 0xcf, 0x23, 0xae, 0xe4, 0x0d, 0x13, 0x82,
	0x04, 0x33, 0x7a, 0x14, 0xd1, 0x6f, 0xdf, 0x67, 0x25, 0x51, 0x9a, 0xec, 0xc0, 0x8b, 0xe0, 0x57,
	0x0f, 0x8c, 0xbe, 0xe8, 0x15, 0x7e, 0xf4, 0xf4, 0x01, 0x3c, 0xba, 0xa2, 0x68, 0xcf, 0x2d, 0x7a,
	0x9f, 0x82, 0x9d, 0xa5, 0xd2, 0x77, 0x79, 0x5d, 0x61, 0xbd, 0xae, 0xdc, 0xaa, 0xac, 0xbb, 0x32,
	0xf6, 0x85, 0x37, 0x4b, 0xf3, 0x9f, 0xfe, 0xdf, 0x77, 0xdf, 0x6f, 0xce, 0xc9, 0x84, 0x90, 0xef,
	0xce, 0x70, 0x62, 0xa3, 0x48, 0x61, 0xb0, 0x1d, 0xdb, 0x29, 0xf3, 0x70, 0xdb, 0xa2, 0xf9, 0xb9,
	0xe7, 0x17, 0xa4, 0x30, 0x8b, 0x03, 0xbf, 0xc8, 0xbc, 0x45, 0xc7, 0xff, 0x76, 0x46, 0x16, 0xf4,
	0x7d, 0x52, 0x18, 0xdf, 0x3f, 0x28, 0x7c, 0x3f, 0xf3, 0x7c, 0xbf, 0xd3, 0x7c, 0x3e, 0xed, 0x16,
	0x7b, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x38, 0x10, 0x8e, 0x8a, 0x97, 0x02, 0x00, 0x00,
}
