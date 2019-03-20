// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/legacy/audit_data.proto

package legacy

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

// Admin Console legacy audit log.
type AuditData struct {
	// Text description of the admin event.
	// This is the "Event" column in Admin Console's Admin Logs.
	EventMessage string `protobuf:"bytes,1,opt,name=event_message,json=eventMessage,proto3" json:"event_message,omitempty"`
	// Arbitrary event data.
	// This is the "Result" column in Admin Console's Admin Logs.
	EventData            map[string]string `protobuf:"bytes,2,rep,name=event_data,json=eventData,proto3" json:"event_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuditData) Reset()         { *m = AuditData{} }
func (m *AuditData) String() string { return proto.CompactTextString(m) }
func (*AuditData) ProtoMessage()    {}
func (*AuditData) Descriptor() ([]byte, []int) {
	return fileDescriptor_74c360c1976d6377, []int{0}
}

func (m *AuditData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditData.Unmarshal(m, b)
}
func (m *AuditData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditData.Marshal(b, m, deterministic)
}
func (m *AuditData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditData.Merge(m, src)
}
func (m *AuditData) XXX_Size() int {
	return xxx_messageInfo_AuditData.Size(m)
}
func (m *AuditData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditData.DiscardUnknown(m)
}

var xxx_messageInfo_AuditData proto.InternalMessageInfo

func (m *AuditData) GetEventMessage() string {
	if m != nil {
		return m.EventMessage
	}
	return ""
}

func (m *AuditData) GetEventData() map[string]string {
	if m != nil {
		return m.EventData
	}
	return nil
}

func init() {
	proto.RegisterType((*AuditData)(nil), "google.appengine.legacy.AuditData")
	proto.RegisterMapType((map[string]string)(nil), "google.appengine.legacy.AuditData.EventDataEntry")
}

func init() {
	proto.RegisterFile("google/appengine/legacy/audit_data.proto", fileDescriptor_74c360c1976d6377)
}

var fileDescriptor_74c360c1976d6377 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xc9, 0x16, 0x85, 0x1d, 0xb5, 0x48, 0x10, 0x5c, 0xf4, 0x52, 0xf4, 0xb2, 0xa7, 0x04,
	0xf5, 0x22, 0xfe, 0x39, 0x58, 0xec, 0x51, 0x58, 0x7a, 0xf4, 0x52, 0xc6, 0x76, 0x18, 0x16, 0xb7,
	0x49, 0xd8, 0x4d, 0x0b, 0xfb, 0xed, 0xfc, 0x68, 0x92, 0xa4, 0x2e, 0x48, 0xe9, 0x29, 0x33, 0x6f,
	0x7e, 0x99, 0x79, 0x3c, 0x28, 0xd9, 0x5a, 0x6e, 0x48, 0xa3, 0x73, 0x64, 0xb8, 0x36, 0xa4, 0x1b,
	0x62, 0x5c, 0xf6, 0x1a, 0x37, 0xab, 0xda, 0x2f, 0x56, 0xe8, 0x51, 0xb9, 0xd6, 0x7a, 0x2b, 0x2f,
	0x13, 0xa9, 0x06, 0x52, 0x25, 0xf2, 0xe6, 0x47, 0x40, 0xfe, 0x16, 0xe8, 0x77, 0xf4, 0x28, 0x6f,
	0xe1, 0x8c, 0xb6, 0x64, 0xfc, 0x62, 0x4d, 0x5d, 0x87, 0x4c, 0x85, 0x98, 0x88, 0x32, 0x9f, 0x9f,
	0x46, 0xf1, 0x23, 0x69, 0xb2, 0x02, 0x48, 0x50, 0xd8, 0x5f, 0x64, 0x93, 0x51, 0x79, 0x72, 0x7f,
	0xa7, 0x0e, 0x1c, 0x50, 0xc3, 0x72, 0x35, 0x0b, 0x9f, 0x42, 0x35, 0x33, 0xbe, 0xed, 0xe7, 0x39,
	0xfd, 0xf5, 0x57, 0x2f, 0x30, 0xfe, 0x3f, 0x94, 0xe7, 0x30, 0xfa, 0xa6, 0x7e, 0x77, 0x3e, 0x94,
	0xf2, 0x02, 0x8e, 0xb6, 0xd8, 0x6c, 0xa8, 0xc8, 0xa2, 0x96, 0x9a, 0xa7, 0xec, 0x51, 0x4c, 0x0d,
	0x5c, 0x2f, 0xed, 0xfa, 0x90, 0x81, 0xe9, 0x78, 0x70, 0x50, 0x85, 0x28, 0x2a, 0xf1, 0xf9, 0xba,
	0x43, 0xd9, 0x36, 0x68, 0x58, 0xd9, 0x96, 0x35, 0x93, 0x89, 0x41, 0xe9, 0x34, 0x42, 0x57, 0x77,
	0x7b, 0xa9, 0x3e, 0xa7, 0xe7, 0xeb, 0x38, 0x92, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0x5d, 0x14, 0xaa, 0x7e, 0x01, 0x00, 0x00,
}
