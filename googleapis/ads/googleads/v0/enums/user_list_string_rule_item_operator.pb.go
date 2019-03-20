// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/user_list_string_rule_item_operator.proto

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

// Enum describing possible user list string rule item operators.
type UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator int32

const (
	// Not specified.
	UserListStringRuleItemOperatorEnum_UNSPECIFIED UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListStringRuleItemOperatorEnum_UNKNOWN UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 1
	// Contains.
	UserListStringRuleItemOperatorEnum_CONTAINS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 2
	// Equals.
	UserListStringRuleItemOperatorEnum_EQUALS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 3
	// Starts with.
	UserListStringRuleItemOperatorEnum_STARTS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 4
	// Ends with.
	UserListStringRuleItemOperatorEnum_ENDS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 5
	// Not equals.
	UserListStringRuleItemOperatorEnum_NOT_EQUALS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 6
	// Not contains.
	UserListStringRuleItemOperatorEnum_NOT_CONTAINS UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 7
	// Not starts with.
	UserListStringRuleItemOperatorEnum_NOT_STARTS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 8
	// Not ends with.
	UserListStringRuleItemOperatorEnum_NOT_ENDS_WITH UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator = 9
)

var UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CONTAINS",
	3: "EQUALS",
	4: "STARTS_WITH",
	5: "ENDS_WITH",
	6: "NOT_EQUALS",
	7: "NOT_CONTAINS",
	8: "NOT_STARTS_WITH",
	9: "NOT_ENDS_WITH",
}

var UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"CONTAINS":        2,
	"EQUALS":          3,
	"STARTS_WITH":     4,
	"ENDS_WITH":       5,
	"NOT_EQUALS":      6,
	"NOT_CONTAINS":    7,
	"NOT_STARTS_WITH": 8,
	"NOT_ENDS_WITH":   9,
}

func (x UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) String() string {
	return proto.EnumName(UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name, int32(x))
}

func (UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_645f3732b1441d95, []int{0, 0}
}

// Supported rule operator for string type.
type UserListStringRuleItemOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListStringRuleItemOperatorEnum) Reset()         { *m = UserListStringRuleItemOperatorEnum{} }
func (m *UserListStringRuleItemOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListStringRuleItemOperatorEnum) ProtoMessage()    {}
func (*UserListStringRuleItemOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_645f3732b1441d95, []int{0}
}

func (m *UserListStringRuleItemOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListStringRuleItemOperatorEnum.Unmarshal(m, b)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListStringRuleItemOperatorEnum.Marshal(b, m, deterministic)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListStringRuleItemOperatorEnum.Merge(m, src)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListStringRuleItemOperatorEnum.Size(m)
}
func (m *UserListStringRuleItemOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListStringRuleItemOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListStringRuleItemOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator", UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_name, UserListStringRuleItemOperatorEnum_UserListStringRuleItemOperator_value)
	proto.RegisterType((*UserListStringRuleItemOperatorEnum)(nil), "google.ads.googleads.v0.enums.UserListStringRuleItemOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/user_list_string_rule_item_operator.proto", fileDescriptor_645f3732b1441d95)
}

var fileDescriptor_645f3732b1441d95 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0xee, 0x93, 0x40,
	0x14, 0xc5, 0x85, 0x6a, 0x3f, 0xa6, 0xad, 0xc5, 0x71, 0x5d, 0x93, 0xd6, 0xfd, 0x40, 0xe2, 0x6e,
	0x5c, 0x4d, 0x5b, 0xac, 0xc4, 0x66, 0xa8, 0x05, 0xda, 0xc4, 0x90, 0x10, 0x94, 0x09, 0x21, 0x01,
	0xa6, 0x99, 0x81, 0x3e, 0x90, 0x4b, 0x1f, 0xc3, 0xa5, 0x6f, 0xa2, 0x4f, 0xe0, 0xd2, 0x30, 0x50,
	0xe2, 0xe6, 0xdf, 0x0d, 0x39, 0xf7, 0xde, 0x73, 0x7f, 0x61, 0xce, 0x05, 0xfb, 0x94, 0xf3, 0x34,
	0x67, 0x66, 0x9c, 0x48, 0xb3, 0x95, 0x8d, 0xba, 0x59, 0x26, 0x2b, 0xeb, 0x42, 0x9a, 0xb5, 0x64,
	0x22, 0xca, 0x33, 0x59, 0x45, 0xb2, 0x12, 0x59, 0x99, 0x46, 0xa2, 0xce, 0x59, 0x94, 0x55, 0xac,
	0x88, 0xf8, 0x95, 0x89, 0xb8, 0xe2, 0x02, 0x5d, 0x05, 0xaf, 0x38, 0x5c, 0xb6, 0xdb, 0x28, 0x4e,
	0x24, 0xea, 0x41, 0xe8, 0x66, 0x21, 0x05, 0x5a, 0xff, 0xd6, 0xc0, 0x3a, 0x90, 0x4c, 0x1c, 0x32,
	0x59, 0x79, 0x0a, 0x75, 0xaa, 0x73, 0xe6, 0x54, 0xac, 0x70, 0x3b, 0x8e, 0x5d, 0xd6, 0xc5, 0xfa,
	0xa7, 0x06, 0xde, 0x3c, 0xb6, 0xc1, 0x05, 0x98, 0x06, 0xd4, 0x3b, 0xda, 0x5b, 0xe7, 0x83, 0x63,
	0xef, 0x8c, 0x67, 0x70, 0x0a, 0x46, 0x01, 0xfd, 0x44, 0xdd, 0x0b, 0x35, 0x34, 0x38, 0x03, 0xe3,
	0xad, 0x4b, 0x7d, 0xe2, 0x50, 0xcf, 0xd0, 0x21, 0x00, 0x43, 0xfb, 0x73, 0x40, 0x0e, 0x9e, 0x31,
	0x68, 0xf6, 0x3c, 0x9f, 0x9c, 0x7c, 0x2f, 0xba, 0x38, 0xfe, 0x47, 0xe3, 0x39, 0x9c, 0x83, 0x89,
	0x4d, 0x77, 0x5d, 0xf9, 0x02, 0xbe, 0x04, 0x80, 0xba, 0x7e, 0xd4, 0xf9, 0x87, 0xd0, 0x00, 0xb3,
	0xa6, 0xee, 0x69, 0x23, 0xf8, 0x1a, 0x2c, 0x9a, 0xce, 0xff, 0x94, 0x31, 0x7c, 0x05, 0xe6, 0x6a,
	0xad, 0x27, 0x4d, 0x36, 0x7f, 0x35, 0xb0, 0xfa, 0xc6, 0x0b, 0xf4, 0x30, 0x91, 0xcd, 0xdb, 0xc7,
	0xef, 0x3c, 0x36, 0xa9, 0x1e, 0xb5, 0x2f, 0x9b, 0x8e, 0x92, 0xf2, 0x3c, 0x2e, 0x53, 0xc4, 0x45,
	0x6a, 0xa6, 0xac, 0x54, 0x99, 0xdf, 0x0f, 0x76, 0xcd, 0xe4, 0x13, 0xf7, 0x7b, 0xaf, 0xbe, 0xdf,
	0xf5, 0xc1, 0x9e, 0x90, 0x1f, 0xfa, 0x72, 0xdf, 0xa2, 0x48, 0x22, 0x51, 0x2b, 0x1b, 0x75, 0xb6,
	0x50, 0x13, 0xbd, 0xfc, 0x75, 0x9f, 0x87, 0x24, 0x91, 0x61, 0x3f, 0x0f, 0xcf, 0x56, 0xa8, 0xe6,
	0x7f, 0xf4, 0x55, 0xdb, 0xc4, 0x98, 0x24, 0x12, 0xe3, 0xde, 0x81, 0xf1, 0xd9, 0xc2, 0x58, 0x79,
	0xbe, 0x0e, 0xd5, 0x8f, 0xbd, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x54, 0x4d, 0x3f, 0x86, 0x57,
	0x02, 0x00, 0x00,
}
