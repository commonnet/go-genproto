// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/security_marks.proto

package securitycenter

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

// User specified security marks that are attached to the parent Cloud Security
// Command Center (Cloud SCC) resource. Security marks are scoped within a Cloud
// SCC organization -- they can be modified and viewed by all users who have
// proper permissions on the organization.
type SecurityMarks struct {
	// The relative resource name of the SecurityMarks. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Examples:
	// "organizations/123/assets/456/securityMarks"
	// "organizations/123/sources/456/findings/789/securityMarks".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mutable user specified security marks belonging to the parent resource.
	// Constraints are as follows:
	//   - Keys and values are treated as case insensitive
	//   - Keys must be between 1 - 256 characters (inclusive)
	//   - Keys must be letters, numbers, underscores, or dashes
	//   - Values have leading and trailing whitespace trimmed, remaining
	//     characters must be between 1 - 4096 characters (inclusive)
	Marks                map[string]string `protobuf:"bytes,2,rep,name=marks,proto3" json:"marks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SecurityMarks) Reset()         { *m = SecurityMarks{} }
func (m *SecurityMarks) String() string { return proto.CompactTextString(m) }
func (*SecurityMarks) ProtoMessage()    {}
func (*SecurityMarks) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bf2c7597a92919b, []int{0}
}

func (m *SecurityMarks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityMarks.Unmarshal(m, b)
}
func (m *SecurityMarks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityMarks.Marshal(b, m, deterministic)
}
func (m *SecurityMarks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityMarks.Merge(m, src)
}
func (m *SecurityMarks) XXX_Size() int {
	return xxx_messageInfo_SecurityMarks.Size(m)
}
func (m *SecurityMarks) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityMarks.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityMarks proto.InternalMessageInfo

func (m *SecurityMarks) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecurityMarks) GetMarks() map[string]string {
	if m != nil {
		return m.Marks
	}
	return nil
}

func init() {
	proto.RegisterType((*SecurityMarks)(nil), "google.cloud.securitycenter.v1.SecurityMarks")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.securitycenter.v1.SecurityMarks.MarksEntry")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/security_marks.proto", fileDescriptor_3bf2c7597a92919b)
}

var fileDescriptor_3bf2c7597a92919b = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x2f, 0x4e, 0x4d, 0x2e, 0x2d, 0xca, 0x2c,
	0xa9, 0x4c, 0x4e, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0x84, 0x8b, 0xc4, 0xe7, 0x26, 0x16,
	0x65, 0x17, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0x41, 0x34, 0xe9, 0x81, 0x35, 0xe9,
	0xa1, 0x6a, 0xd2, 0x2b, 0x33, 0x94, 0x92, 0x81, 0x1a, 0x9a, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97,
	0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x07, 0xd5, 0xad, 0xb4, 0x96, 0x91, 0x8b, 0x37, 0x18,
	0xaa, 0xc7, 0x17, 0x64, 0xaa, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0xe4, 0xc7, 0xc5, 0x0a, 0xb6, 0x52, 0x82, 0x49, 0x81, 0x59,
	0x83, 0xdb, 0xc8, 0x42, 0x0f, 0xbf, 0x9d, 0x7a, 0x28, 0x26, 0xea, 0x81, 0x49, 0xd7, 0xbc, 0x92,
	0xa2, 0xca, 0x20, 0x88, 0x31, 0x52, 0x16, 0x5c, 0x5c, 0x08, 0x41, 0x21, 0x01, 0x2e, 0xe6, 0xec,
	0xd4, 0x4a, 0xa8, 0x85, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04,
	0x13, 0x58, 0x0c, 0xc2, 0xb1, 0x62, 0xb2, 0x60, 0x74, 0xda, 0xc6, 0xc8, 0xa5, 0x94, 0x9c, 0x9f,
	0x4b, 0xc0, 0x01, 0x01, 0x8c, 0x51, 0x3e, 0x50, 0x15, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a,
	0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x4f, 0xeb, 0x43, 0xa4, 0x12, 0x0b, 0x32, 0x8b,
	0x71, 0x05, 0xb5, 0x35, 0xaa, 0xc8, 0x2a, 0x26, 0x39, 0x77, 0x88, 0x71, 0xce, 0x60, 0x0b, 0x61,
	0xfe, 0x73, 0x86, 0x58, 0x18, 0x66, 0x78, 0x0a, 0xa6, 0x20, 0x06, 0xac, 0x20, 0x06, 0x55, 0x41,
	0x4c, 0x98, 0x61, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x04,
	0x8b, 0x20, 0xe4, 0x01, 0x00, 0x00,
}
