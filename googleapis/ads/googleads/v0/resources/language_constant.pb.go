// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/language_constant.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A language.
type LanguageConstant struct {
	// The resource name of the language constant.
	// Language constant resource names have the form:
	//
	// `languageConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the language constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The language code, e.g. "en_US", "en_AU", "es", "fr", etc.
	Code *wrappers.StringValue `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The full name of the language in English, e.g., "English (US)", "Spanish",
	// etc.
	Name                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LanguageConstant) Reset()         { *m = LanguageConstant{} }
func (m *LanguageConstant) String() string { return proto.CompactTextString(m) }
func (*LanguageConstant) ProtoMessage()    {}
func (*LanguageConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_05acff603327c3aa, []int{0}
}

func (m *LanguageConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LanguageConstant.Unmarshal(m, b)
}
func (m *LanguageConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LanguageConstant.Marshal(b, m, deterministic)
}
func (m *LanguageConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LanguageConstant.Merge(m, src)
}
func (m *LanguageConstant) XXX_Size() int {
	return xxx_messageInfo_LanguageConstant.Size(m)
}
func (m *LanguageConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_LanguageConstant.DiscardUnknown(m)
}

var xxx_messageInfo_LanguageConstant proto.InternalMessageInfo

func (m *LanguageConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *LanguageConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LanguageConstant) GetCode() *wrappers.StringValue {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *LanguageConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*LanguageConstant)(nil), "google.ads.googleads.v0.resources.LanguageConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/language_constant.proto", fileDescriptor_05acff603327c3aa)
}

var fileDescriptor_05acff603327c3aa = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0x69, 0x36, 0x3e, 0xf8, 0xaa, 0x82, 0x14, 0x84, 0xa1, 0x22, 0x9b, 0x32, 0x18, 0x08,
	0x69, 0x51, 0x11, 0x8c, 0x57, 0x9d, 0x17, 0x43, 0x11, 0x19, 0x13, 0x7a, 0x21, 0x85, 0x91, 0x35,
	0x31, 0x14, 0xb6, 0xa4, 0x24, 0xed, 0xbc, 0xf7, 0x51, 0xbc, 0xf4, 0x49, 0xc4, 0x47, 0xf1, 0x29,
	0x24, 0x49, 0x93, 0x0b, 0x05, 0xf5, 0xee, 0x90, 0xfc, 0xce, 0x39, 0xc9, 0xff, 0x1f, 0x5e, 0x30,
	0x21, 0xd8, 0x92, 0xc6, 0x98, 0xa8, 0xd8, 0x4a, 0xad, 0xd6, 0x49, 0x2c, 0xa9, 0x12, 0x8d, 0x2c,
	0xa8, 0x8a, 0x97, 0x98, 0xb3, 0x06, 0x33, 0x3a, 0x2f, 0x04, 0x57, 0x35, 0xe6, 0x35, 0xac, 0xa4,
	0xa8, 0x45, 0x34, 0xb0, 0x3c, 0xc4, 0x44, 0x41, 0x6f, 0x85, 0xeb, 0x04, 0x7a, 0xeb, 0xee, 0x41,
	0x9b, 0x6e, 0x0c, 0x8b, 0xe6, 0x31, 0x7e, 0x92, 0xb8, 0xaa, 0xa8, 0x54, 0x36, 0xe2, 0xf0, 0x2d,
	0x08, 0xb7, 0x6f, 0xdb, 0xf8, 0xab, 0x36, 0x3d, 0x3a, 0x0a, 0xb7, 0x5c, 0xc2, 0x9c, 0xe3, 0x15,
	0xed, 0x05, 0xfd, 0x60, 0xf4, 0x7f, 0xb6, 0xe9, 0x0e, 0xef, 0xf0, 0x8a, 0x46, 0xc7, 0x21, 0x28,
	0x49, 0x0f, 0xf4, 0x83, 0xd1, 0xc6, 0xc9, 0x5e, 0x5b, 0x0f, 0x5d, 0x0d, 0xbc, 0xe6, 0xf5, 0xf9,
	0x59, 0x86, 0x97, 0x0d, 0x9d, 0x81, 0x92, 0x44, 0x49, 0xd8, 0x2d, 0x04, 0xa1, 0xbd, 0x8e, 0xc1,
	0xf7, 0xbf, 0xe1, 0xf7, 0xb5, 0x2c, 0x39, 0xb3, 0xbc, 0x21, 0xb5, 0xc3, 0x54, 0x77, 0xff, 0xe2,
	0xd0, 0xe4, 0xf8, 0x19, 0x84, 0xc3, 0x42, 0xac, 0xe0, 0xaf, 0x43, 0x19, 0xef, 0x7c, 0xfd, 0xf1,
	0x54, 0xa7, 0x4e, 0x83, 0x87, 0x9b, 0xd6, 0xcb, 0x84, 0x1e, 0x39, 0x14, 0x92, 0xc5, 0x8c, 0x72,
	0xd3, 0xe9, 0x76, 0x53, 0x95, 0xea, 0x87, 0x55, 0x5d, 0x7a, 0xf5, 0x02, 0x3a, 0x93, 0x34, 0x7d,
	0x05, 0x83, 0x89, 0x8d, 0x4c, 0x89, 0x82, 0x56, 0x6a, 0x95, 0x25, 0x70, 0xe6, 0xc8, 0x77, 0xc7,
	0xe4, 0x29, 0x51, 0xb9, 0x67, 0xf2, 0x2c, 0xc9, 0x3d, 0xf3, 0x01, 0x86, 0xf6, 0x02, 0xa1, 0x94,
	0x28, 0x84, 0x3c, 0x85, 0x50, 0x96, 0x20, 0xe4, 0xb9, 0xc5, 0x3f, 0xf3, 0xd8, 0xd3, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1f, 0x35, 0xd2, 0x7c, 0x56, 0x02, 0x00, 0x00,
}
