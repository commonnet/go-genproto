// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/conversion_action_type.proto

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

// Possible types of a conversion action.
type ConversionActionTypeEnum_ConversionActionType int32

const (
	// Not specified.
	ConversionActionTypeEnum_UNSPECIFIED ConversionActionTypeEnum_ConversionActionType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionTypeEnum_UNKNOWN ConversionActionTypeEnum_ConversionActionType = 1
	// Conversions that occur when a user clicks on an ad's call extension.
	ConversionActionTypeEnum_AD_CALL ConversionActionTypeEnum_ConversionActionType = 2
	// Conversions that occur when a user on a mobile device clicks a phone
	// number.
	ConversionActionTypeEnum_CLICK_TO_CALL ConversionActionTypeEnum_ConversionActionType = 3
	// Conversions that occur when a user downloads a mobile app from the Google
	// Play Store.
	ConversionActionTypeEnum_GOOGLE_PLAY_DOWNLOAD ConversionActionTypeEnum_ConversionActionType = 4
	// Conversions that occur when a user makes a purchase in an app through
	// Android billing.
	ConversionActionTypeEnum_GOOGLE_PLAY_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 5
	// Call conversions that are tracked by the advertiser and uploaded.
	ConversionActionTypeEnum_UPLOAD_CALLS ConversionActionTypeEnum_ConversionActionType = 6
	// Conversions that are tracked by the advertiser and uploaded with
	// attributed clicks.
	ConversionActionTypeEnum_UPLOAD_CLICKS ConversionActionTypeEnum_ConversionActionType = 7
	// Conversions that occur on a webpage.
	ConversionActionTypeEnum_WEBPAGE ConversionActionTypeEnum_ConversionActionType = 8
	// Conversions that occur when a user calls a dynamically-generated phone
	// number from an advertiser's website.
	ConversionActionTypeEnum_WEBSITE_CALL ConversionActionTypeEnum_ConversionActionType = 9
)

var ConversionActionTypeEnum_ConversionActionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_CALL",
	3: "CLICK_TO_CALL",
	4: "GOOGLE_PLAY_DOWNLOAD",
	5: "GOOGLE_PLAY_IN_APP_PURCHASE",
	6: "UPLOAD_CALLS",
	7: "UPLOAD_CLICKS",
	8: "WEBPAGE",
	9: "WEBSITE_CALL",
}

var ConversionActionTypeEnum_ConversionActionType_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"AD_CALL":                     2,
	"CLICK_TO_CALL":               3,
	"GOOGLE_PLAY_DOWNLOAD":        4,
	"GOOGLE_PLAY_IN_APP_PURCHASE": 5,
	"UPLOAD_CALLS":                6,
	"UPLOAD_CLICKS":               7,
	"WEBPAGE":                     8,
	"WEBSITE_CALL":                9,
}

func (x ConversionActionTypeEnum_ConversionActionType) String() string {
	return proto.EnumName(ConversionActionTypeEnum_ConversionActionType_name, int32(x))
}

func (ConversionActionTypeEnum_ConversionActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a2fff4614cf23a45, []int{0, 0}
}

// Container for enum describing possible types of a conversion action.
type ConversionActionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionTypeEnum) Reset()         { *m = ConversionActionTypeEnum{} }
func (m *ConversionActionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionTypeEnum) ProtoMessage()    {}
func (*ConversionActionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2fff4614cf23a45, []int{0}
}

func (m *ConversionActionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionTypeEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionTypeEnum.Merge(m, src)
}
func (m *ConversionActionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionTypeEnum.Size(m)
}
func (m *ConversionActionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.ConversionActionTypeEnum_ConversionActionType", ConversionActionTypeEnum_ConversionActionType_name, ConversionActionTypeEnum_ConversionActionType_value)
	proto.RegisterType((*ConversionActionTypeEnum)(nil), "google.ads.googleads.v0.enums.ConversionActionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/conversion_action_type.proto", fileDescriptor_a2fff4614cf23a45)
}

var fileDescriptor_a2fff4614cf23a45 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x86, 0x85, 0x6a, 0xab, 0x53, 0x8d, 0x23, 0xe9, 0xa2, 0xc6, 0x34, 0xa6, 0x7d, 0x80, 0x81,
	0xc4, 0xdd, 0xb8, 0x1a, 0x60, 0x44, 0x52, 0x02, 0x13, 0x29, 0x25, 0x1a, 0x12, 0x82, 0x85, 0x90,
	0x26, 0x2d, 0x43, 0x3a, 0x6d, 0x93, 0xbe, 0x8e, 0x4b, 0x1f, 0xc5, 0xad, 0x6f, 0xe1, 0xc2, 0xf8,
	0x08, 0x66, 0x86, 0x5b, 0x72, 0x17, 0xbd, 0x77, 0x03, 0xff, 0x9c, 0xff, 0x7c, 0xe7, 0xcc, 0x9c,
	0x03, 0x70, 0xcd, 0x79, 0xbd, 0xab, 0xcc, 0xa2, 0x14, 0x66, 0x27, 0xa5, 0x3a, 0x5b, 0x66, 0xd5,
	0x9c, 0xf6, 0xc2, 0xdc, 0xf0, 0xe6, 0x5c, 0x1d, 0xc4, 0x96, 0x37, 0x79, 0xb1, 0x39, 0xca, 0xdf,
	0xf1, 0xd2, 0x56, 0xa8, 0x3d, 0xf0, 0x23, 0x37, 0x66, 0x1d, 0x80, 0x8a, 0x52, 0xa0, 0x9e, 0x45,
	0x67, 0x0b, 0x29, 0x76, 0xf1, 0x4f, 0x03, 0x53, 0xa7, 0xe7, 0x89, 0xc2, 0x57, 0x97, 0xb6, 0xa2,
	0xcd, 0x69, 0xbf, 0xf8, 0xad, 0x81, 0xc9, 0x2d, 0xd3, 0x78, 0x0d, 0xc6, 0x49, 0x18, 0x33, 0xea,
	0xf8, 0x9f, 0x7c, 0xea, 0xc2, 0x27, 0xc6, 0x18, 0x8c, 0x92, 0x70, 0x19, 0x46, 0x69, 0x08, 0x35,
	0x79, 0x20, 0x6e, 0xee, 0x90, 0x20, 0x80, 0xba, 0xf1, 0x06, 0xbc, 0x72, 0x02, 0xdf, 0x59, 0xe6,
	0xab, 0xa8, 0x0b, 0x0d, 0x8c, 0x29, 0x98, 0x78, 0x51, 0xe4, 0x05, 0x34, 0x67, 0x01, 0xf9, 0x9a,
	0xbb, 0x51, 0x1a, 0x06, 0x11, 0x71, 0xe1, 0x53, 0xe3, 0x3d, 0x78, 0x77, 0xdf, 0xf1, 0xc3, 0x9c,
	0x30, 0x96, 0xb3, 0xe4, 0x8b, 0xf3, 0x99, 0xc4, 0x14, 0x3e, 0x33, 0x20, 0x78, 0x99, 0x30, 0x99,
	0xac, 0x6a, 0xc5, 0x70, 0x28, 0xeb, 0x5f, 0x23, 0xb2, 0x4d, 0x0c, 0x47, 0xb2, 0x7f, 0x4a, 0x6d,
	0x46, 0x3c, 0x0a, 0x9f, 0x4b, 0x22, 0xa5, 0x76, 0xec, 0xaf, 0x68, 0xd7, 0xfe, 0x85, 0xfd, 0x57,
	0x03, 0xf3, 0x0d, 0xdf, 0xa3, 0x47, 0x07, 0x63, 0xbf, 0xbd, 0xf5, 0x70, 0x26, 0x47, 0xca, 0xb4,
	0x6f, 0xf6, 0x1d, 0x5b, 0xf3, 0x5d, 0xd1, 0xd4, 0x88, 0x1f, 0x6a, 0xb3, 0xae, 0x1a, 0x35, 0xf0,
	0xeb, 0x82, 0xda, 0xad, 0x78, 0x60, 0x5f, 0x1f, 0xd5, 0xf7, 0x87, 0x3e, 0xf0, 0x08, 0xf9, 0xa9,
	0xcf, 0xbc, 0xae, 0x14, 0x29, 0x05, 0xea, 0xa4, 0x54, 0x6b, 0x0b, 0xc9, 0x0d, 0x88, 0x5f, 0x57,
	0x3f, 0x23, 0xa5, 0xc8, 0x7a, 0x3f, 0x5b, 0x5b, 0x99, 0xf2, 0xff, 0xe8, 0xf3, 0x2e, 0x88, 0x31,
	0x29, 0x05, 0xc6, 0x7d, 0x06, 0xc6, 0x6b, 0x0b, 0x63, 0x95, 0xf3, 0x7d, 0xa8, 0x2e, 0xf6, 0xe1,
	0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x71, 0x97, 0x07, 0x47, 0x02, 0x00, 0x00,
}
