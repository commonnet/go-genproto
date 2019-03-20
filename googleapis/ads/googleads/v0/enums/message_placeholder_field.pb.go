// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/message_placeholder_field.proto

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

// Possible values for Message placeholder fields.
type MessagePlaceholderFieldEnum_MessagePlaceholderField int32

const (
	// Not specified.
	MessagePlaceholderFieldEnum_UNSPECIFIED MessagePlaceholderFieldEnum_MessagePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	MessagePlaceholderFieldEnum_UNKNOWN MessagePlaceholderFieldEnum_MessagePlaceholderField = 1
	// Data Type: STRING. The name of your business.
	MessagePlaceholderFieldEnum_BUSINESS_NAME MessagePlaceholderFieldEnum_MessagePlaceholderField = 2
	// Data Type: STRING. Country code of phone number.
	MessagePlaceholderFieldEnum_COUNTRY_CODE MessagePlaceholderFieldEnum_MessagePlaceholderField = 3
	// Data Type: STRING. A phone number that's capable of sending and receiving
	// text messages.
	MessagePlaceholderFieldEnum_PHONE_NUMBER MessagePlaceholderFieldEnum_MessagePlaceholderField = 4
	// Data Type: STRING. The text that will go in your click-to-message ad.
	MessagePlaceholderFieldEnum_MESSAGE_EXTENSION_TEXT MessagePlaceholderFieldEnum_MessagePlaceholderField = 5
	// Data Type: STRING. The message text automatically shows in people's
	// messaging apps when they tap to send you a message.
	MessagePlaceholderFieldEnum_MESSAGE_TEXT MessagePlaceholderFieldEnum_MessagePlaceholderField = 6
)

var MessagePlaceholderFieldEnum_MessagePlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BUSINESS_NAME",
	3: "COUNTRY_CODE",
	4: "PHONE_NUMBER",
	5: "MESSAGE_EXTENSION_TEXT",
	6: "MESSAGE_TEXT",
}

var MessagePlaceholderFieldEnum_MessagePlaceholderField_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"BUSINESS_NAME":          2,
	"COUNTRY_CODE":           3,
	"PHONE_NUMBER":           4,
	"MESSAGE_EXTENSION_TEXT": 5,
	"MESSAGE_TEXT":           6,
}

func (x MessagePlaceholderFieldEnum_MessagePlaceholderField) String() string {
	return proto.EnumName(MessagePlaceholderFieldEnum_MessagePlaceholderField_name, int32(x))
}

func (MessagePlaceholderFieldEnum_MessagePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_524e7aae95a8a6af, []int{0, 0}
}

// Values for Message placeholder fields.
type MessagePlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagePlaceholderFieldEnum) Reset()         { *m = MessagePlaceholderFieldEnum{} }
func (m *MessagePlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*MessagePlaceholderFieldEnum) ProtoMessage()    {}
func (*MessagePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_524e7aae95a8a6af, []int{0}
}

func (m *MessagePlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagePlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *MessagePlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagePlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *MessagePlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagePlaceholderFieldEnum.Merge(m, src)
}
func (m *MessagePlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_MessagePlaceholderFieldEnum.Size(m)
}
func (m *MessagePlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagePlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessagePlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.MessagePlaceholderFieldEnum_MessagePlaceholderField", MessagePlaceholderFieldEnum_MessagePlaceholderField_name, MessagePlaceholderFieldEnum_MessagePlaceholderField_value)
	proto.RegisterType((*MessagePlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.MessagePlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/message_placeholder_field.proto", fileDescriptor_524e7aae95a8a6af)
}

var fileDescriptor_524e7aae95a8a6af = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x4e, 0x83, 0x30,
	0x1c, 0x16, 0xa6, 0x33, 0xe9, 0x34, 0x22, 0x07, 0x4d, 0xd4, 0x1d, 0xb6, 0x07, 0x28, 0x24, 0xde,
	0x6a, 0x3c, 0xc0, 0xd6, 0x4d, 0x62, 0x28, 0x64, 0xc0, 0x9c, 0x86, 0x84, 0xe0, 0xa8, 0xb8, 0x04,
	0xe8, 0xb2, 0xba, 0x3d, 0x8d, 0x27, 0x8f, 0x3e, 0x83, 0x4f, 0xe0, 0xa3, 0x78, 0xf3, 0x0d, 0x4c,
	0x8b, 0x9b, 0xa7, 0x79, 0x21, 0x5f, 0xbe, 0x7f, 0xe1, 0xf7, 0x15, 0x5c, 0xe7, 0x8c, 0xe5, 0x05,
	0x35, 0xd2, 0x8c, 0x1b, 0x35, 0x14, 0x68, 0x65, 0x1a, 0xb4, 0x5a, 0x96, 0xdc, 0x28, 0x29, 0xe7,
	0x69, 0x4e, 0x93, 0x79, 0x91, 0x4e, 0xe9, 0x33, 0x2b, 0x32, 0xba, 0x48, 0x9e, 0x66, 0xb4, 0xc8,
	0xe0, 0x7c, 0xc1, 0x5e, 0x98, 0xde, 0xae, 0x33, 0x30, 0xcd, 0x38, 0xdc, 0xc4, 0xe1, 0xca, 0x84,
	0x32, 0xde, 0xfd, 0x50, 0xc0, 0xb9, 0x5b, 0x57, 0xf8, 0x7f, 0x0d, 0x03, 0x51, 0x80, 0xab, 0x65,
	0xd9, 0x7d, 0x55, 0xc0, 0xe9, 0x16, 0x5d, 0x3f, 0x02, 0xad, 0x88, 0x04, 0x3e, 0xee, 0x39, 0x03,
	0x07, 0xf7, 0xb5, 0x1d, 0xbd, 0x05, 0xf6, 0x23, 0x72, 0x4b, 0xbc, 0x3b, 0xa2, 0x29, 0xfa, 0x31,
	0x38, 0xb4, 0xa3, 0xc0, 0x21, 0x38, 0x08, 0x12, 0x62, 0xb9, 0x58, 0x53, 0x75, 0x0d, 0x1c, 0xf4,
	0xbc, 0x88, 0x84, 0xa3, 0xfb, 0xa4, 0xe7, 0xf5, 0xb1, 0xd6, 0x10, 0x8c, 0x7f, 0xe3, 0x11, 0x9c,
	0x90, 0xc8, 0xb5, 0xf1, 0x48, 0xdb, 0xd5, 0xcf, 0xc0, 0x89, 0x8b, 0x83, 0xc0, 0x1a, 0xe2, 0x04,
	0x4f, 0x42, 0x4c, 0x02, 0xc7, 0x23, 0x49, 0x88, 0x27, 0xa1, 0xb6, 0x27, 0xdc, 0x6b, 0x4d, 0x32,
	0x4d, 0xfb, 0x5b, 0x01, 0x9d, 0x29, 0x2b, 0xe1, 0xbf, 0x47, 0xda, 0x17, 0x5b, 0x2e, 0xf0, 0xc5,
	0x42, 0xbe, 0xf2, 0x60, 0xff, 0xc6, 0x73, 0x56, 0xa4, 0x55, 0x0e, 0xd9, 0x22, 0x37, 0x72, 0x5a,
	0xc9, 0xfd, 0xd6, 0x93, 0xcf, 0x67, 0x7c, 0xcb, 0x0b, 0x5c, 0xc9, 0xef, 0x9b, 0xda, 0x18, 0x5a,
	0xd6, 0xbb, 0xda, 0x1e, 0xd6, 0x55, 0x56, 0xc6, 0x61, 0x0d, 0x05, 0x1a, 0x9b, 0x50, 0xac, 0xc9,
	0x3f, 0xd7, 0x7a, 0x6c, 0x65, 0x3c, 0xde, 0xe8, 0xf1, 0xd8, 0x8c, 0xa5, 0xfe, 0xa5, 0x76, 0x6a,
	0x12, 0x21, 0x2b, 0xe3, 0x08, 0x6d, 0x1c, 0x08, 0x8d, 0x4d, 0x84, 0xa4, 0xe7, 0xb1, 0x29, 0x7f,
	0xec, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x47, 0x86, 0x29, 0x2c, 0x19, 0x02, 0x00, 0x00,
}
