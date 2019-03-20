// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/gender_view.proto

package resources

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

// A gender view.
type GenderView struct {
	// The resource name of the gender view.
	// Gender view resource names have the form:
	//
	// `customers/{customer_id}/genderViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenderView) Reset()         { *m = GenderView{} }
func (m *GenderView) String() string { return proto.CompactTextString(m) }
func (*GenderView) ProtoMessage()    {}
func (*GenderView) Descriptor() ([]byte, []int) {
	return fileDescriptor_1caf9d1d919debe3, []int{0}
}

func (m *GenderView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenderView.Unmarshal(m, b)
}
func (m *GenderView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenderView.Marshal(b, m, deterministic)
}
func (m *GenderView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenderView.Merge(m, src)
}
func (m *GenderView) XXX_Size() int {
	return xxx_messageInfo_GenderView.Size(m)
}
func (m *GenderView) XXX_DiscardUnknown() {
	xxx_messageInfo_GenderView.DiscardUnknown(m)
}

var xxx_messageInfo_GenderView proto.InternalMessageInfo

func (m *GenderView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GenderView)(nil), "google.ads.googleads.v1.resources.GenderView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/gender_view.proto", fileDescriptor_1caf9d1d919debe3)
}

var fileDescriptor_1caf9d1d919debe3 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x69, 0x05, 0xc1, 0xa0, 0x08, 0x37, 0x89, 0x38, 0x78, 0xca, 0x81, 0x53, 0x42, 0xb8,
	0x2d, 0x4e, 0xb9, 0xa5, 0xe0, 0x20, 0xc7, 0x0d, 0x1d, 0xa4, 0x70, 0xc4, 0x4b, 0x08, 0x81, 0x6b,
	0x5e, 0xc9, 0xab, 0xbd, 0xef, 0xe3, 0xe8, 0x47, 0xf1, 0xa3, 0xf8, 0x19, 0x1c, 0xa4, 0x8d, 0x89,
	0x9b, 0xb7, 0xfd, 0x69, 0x7f, 0xff, 0xdf, 0x7b, 0x2f, 0x64, 0x69, 0x01, 0xec, 0xde, 0x30, 0xa5,
	0x91, 0xc5, 0x38, 0xa6, 0x81, 0xb3, 0x60, 0x10, 0xde, 0xc2, 0xce, 0x20, 0xb3, 0xc6, 0x6b, 0x13,
	0xb6, 0x83, 0x33, 0x07, 0xda, 0x05, 0xe8, 0x61, 0x36, 0x8f, 0x24, 0x55, 0x1a, 0x69, 0x2e, 0xd1,
	0x81, 0xd3, 0x5c, 0xba, 0xbe, 0x49, 0xde, 0xce, 0x31, 0xe5, 0x3d, 0xf4, 0xaa, 0x77, 0xe0, 0x31,
	0x0a, 0xee, 0x38, 0x21, 0xd5, 0x64, 0xad, 0x9d, 0x39, 0xcc, 0xee, 0xc9, 0x45, 0x2a, 0x6e, 0xbd,
	0x6a, 0xcd, 0x55, 0x71, 0x5b, 0x3c, 0x9c, 0x6d, 0xce, 0xd3, 0xc7, 0x67, 0xd5, 0x9a, 0xd5, 0x77,
	0x41, 0x16, 0x3b, 0x68, 0xe9, 0xd1, 0xd1, 0xab, 0xcb, 0x3f, 0xf5, 0x7a, 0x9c, 0xb6, 0x2e, 0x5e,
	0x9e, 0x7e, 0x5b, 0x16, 0xf6, 0xca, 0x5b, 0x0a, 0xc1, 0x8e, 0x57, 0x4d, 0xbb, 0xa4, 0xab, 0x3b,
	0x87, 0xff, 0x3c, 0xc2, 0x63, 0x4e, 0xef, 0xe5, 0x49, 0x25, 0xe5, 0x47, 0x39, 0xaf, 0xa2, 0x52,
	0x6a, 0xa4, 0x31, 0x8e, 0xa9, 0xe6, 0x74, 0x93, 0xc8, 0xcf, 0xc4, 0x34, 0x52, 0x63, 0x93, 0x99,
	0xa6, 0xe6, 0x4d, 0x66, 0xbe, 0xca, 0x45, 0xfc, 0x21, 0x84, 0xd4, 0x28, 0x44, 0xa6, 0x84, 0xa8,
	0xb9, 0x10, 0x99, 0x7b, 0x3d, 0x9d, 0x96, 0x5d, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x37,
	0x36, 0xcd, 0xb0, 0x01, 0x00, 0x00,
}
