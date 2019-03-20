// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/mobile_app_category_constant.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A mobile application category constant.
type MobileAppCategoryConstant struct {
	// The resource name of the mobile app category constant.
	// Mobile app category constant resource names have the form:
	//
	// `mobileAppCategoryConstants/{mobile_app_category_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the mobile app category constant.
	Id *wrappers.Int32Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Mobile app category name.
	Name                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileAppCategoryConstant) Reset()         { *m = MobileAppCategoryConstant{} }
func (m *MobileAppCategoryConstant) String() string { return proto.CompactTextString(m) }
func (*MobileAppCategoryConstant) ProtoMessage()    {}
func (*MobileAppCategoryConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d66e9d0b240fe19, []int{0}
}

func (m *MobileAppCategoryConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileAppCategoryConstant.Unmarshal(m, b)
}
func (m *MobileAppCategoryConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileAppCategoryConstant.Marshal(b, m, deterministic)
}
func (m *MobileAppCategoryConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileAppCategoryConstant.Merge(m, src)
}
func (m *MobileAppCategoryConstant) XXX_Size() int {
	return xxx_messageInfo_MobileAppCategoryConstant.Size(m)
}
func (m *MobileAppCategoryConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileAppCategoryConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileAppCategoryConstant proto.InternalMessageInfo

func (m *MobileAppCategoryConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileAppCategoryConstant) GetId() *wrappers.Int32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileAppCategoryConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileAppCategoryConstant)(nil), "google.ads.googleads.v1.resources.MobileAppCategoryConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/mobile_app_category_constant.proto", fileDescriptor_1d66e9d0b240fe19)
}

var fileDescriptor_1d66e9d0b240fe19 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xeb, 0x30,
	0x18, 0xc7, 0x69, 0x77, 0x38, 0x70, 0x7a, 0xf4, 0xa6, 0x57, 0x73, 0x8e, 0xb1, 0x29, 0x83, 0x81,
	0x90, 0xda, 0xed, 0x2e, 0x5e, 0x75, 0x13, 0x86, 0x82, 0x32, 0x26, 0xf4, 0x42, 0x0a, 0x25, 0x6b,
	0x63, 0x28, 0xb4, 0x49, 0x48, 0xb2, 0x89, 0xcf, 0xe0, 0x43, 0x08, 0x5e, 0xfa, 0x28, 0x3e, 0x8a,
	0x4f, 0x21, 0x4b, 0x9a, 0xdc, 0x88, 0x7a, 0xf7, 0xa7, 0xf9, 0xe5, 0xff, 0xfd, 0x9a, 0x2f, 0xb8,
	0x24, 0x8c, 0x91, 0x1a, 0x47, 0xa8, 0x94, 0x91, 0x89, 0xfb, 0xb4, 0x8b, 0x23, 0x81, 0x25, 0xdb,
	0x8a, 0x02, 0xcb, 0xa8, 0x61, 0x9b, 0xaa, 0xc6, 0x39, 0xe2, 0x3c, 0x2f, 0x90, 0xc2, 0x84, 0x89,
	0xa7, 0xbc, 0x60, 0x54, 0x2a, 0x44, 0x15, 0xe0, 0x82, 0x29, 0x16, 0x8e, 0xcc, 0x55, 0x80, 0x4a,
	0x09, 0x5c, 0x0b, 0xd8, 0xc5, 0xc0, 0xb5, 0xf4, 0x06, 0xed, 0x20, 0x7d, 0x61, 0xb3, 0x7d, 0x88,
	0x1e, 0x05, 0xe2, 0x1c, 0x0b, 0x69, 0x2a, 0x7a, 0x7d, 0x2b, 0xc2, 0xab, 0x08, 0x51, 0xca, 0x14,
	0x52, 0x15, 0xa3, 0xed, 0xe9, 0xc9, 0x8b, 0x17, 0x1c, 0xdd, 0x68, 0x8f, 0x84, 0xf3, 0x45, 0x6b,
	0xb1, 0x68, 0x25, 0xc2, 0xd3, 0xe0, 0xd0, 0x0e, 0xca, 0x29, 0x6a, 0x70, 0xd7, 0x1b, 0x7a, 0x93,
	0x7f, 0xeb, 0x03, 0xfb, 0xf1, 0x16, 0x35, 0x38, 0x3c, 0x0b, 0xfc, 0xaa, 0xec, 0xfa, 0x43, 0x6f,
	0xf2, 0x7f, 0x7a, 0xdc, 0x5a, 0x02, 0x6b, 0x03, 0xae, 0xa8, 0x9a, 0x4d, 0x53, 0x54, 0x6f, 0xf1,
	0xda, 0xaf, 0xca, 0xf0, 0x3c, 0xf8, 0xa3, 0x8b, 0x3a, 0x1a, 0xef, 0x7f, 0xc1, 0xef, 0x94, 0xa8,
	0x28, 0x31, 0xbc, 0x26, 0xe7, 0xcf, 0x7e, 0x30, 0x2e, 0x58, 0x03, 0x7e, 0x7d, 0x89, 0xf9, 0xe0,
	0xdb, 0x1f, 0x59, 0xed, 0xeb, 0x57, 0xde, 0xfd, 0x75, 0x5b, 0x42, 0x58, 0x8d, 0x28, 0x01, 0x4c,
	0x90, 0x88, 0x60, 0xaa, 0x87, 0xdb, 0x25, 0xf1, 0x4a, 0xfe, 0xb0, 0xb3, 0x0b, 0x97, 0x5e, 0xfd,
	0xce, 0x32, 0x49, 0xde, 0xfc, 0xd1, 0xd2, 0x54, 0x26, 0xa5, 0x04, 0x26, 0xee, 0x53, 0x1a, 0x83,
	0xb5, 0x25, 0xdf, 0x2d, 0x93, 0x25, 0xa5, 0xcc, 0x1c, 0x93, 0xa5, 0x71, 0xe6, 0x98, 0x0f, 0x7f,
	0x6c, 0x0e, 0x20, 0x4c, 0x4a, 0x09, 0xa1, 0xa3, 0x20, 0x4c, 0x63, 0x08, 0x1d, 0xb7, 0xf9, 0xab,
	0x65, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x03, 0x3b, 0x4a, 0x5f, 0x02, 0x00, 0x00,
}
