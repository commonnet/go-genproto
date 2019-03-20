// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/operations.proto

package automl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Metadata used across all long running operations returned by AutoML API.
type OperationMetadata struct {
	// Ouptut only. Details of specific operation. Even if this field is empty,
	// the presence allows to distinguish different types of operations.
	//
	// Types that are valid to be assigned to Details:
	//	*OperationMetadata_CreateModelDetails
	Details isOperationMetadata_Details `protobuf_oneof:"details"`
	// Output only. Progress of operation. Range: [0, 100].
	ProgressPercent int32 `protobuf:"varint,13,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	// Output only. Partial failures encountered.
	// E.g. single files that couldn't be read.
	// This field should never exceed 20 entries.
	// Status details field will contain standard GCP error details.
	PartialFailures []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
	// Output only. Time when the operation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time when the operation was updated for the last time.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OperationMetadata) Reset()         { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()    {}
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_07afd87baa9bfac3, []int{0}
}

func (m *OperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationMetadata.Unmarshal(m, b)
}
func (m *OperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationMetadata.Marshal(b, m, deterministic)
}
func (m *OperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationMetadata.Merge(m, src)
}
func (m *OperationMetadata) XXX_Size() int {
	return xxx_messageInfo_OperationMetadata.Size(m)
}
func (m *OperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OperationMetadata proto.InternalMessageInfo

type isOperationMetadata_Details interface {
	isOperationMetadata_Details()
}

type OperationMetadata_CreateModelDetails struct {
	CreateModelDetails *CreateModelOperationMetadata `protobuf:"bytes,10,opt,name=create_model_details,json=createModelDetails,proto3,oneof"`
}

func (*OperationMetadata_CreateModelDetails) isOperationMetadata_Details() {}

func (m *OperationMetadata) GetDetails() isOperationMetadata_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *OperationMetadata) GetCreateModelDetails() *CreateModelOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_CreateModelDetails); ok {
		return x.CreateModelDetails
	}
	return nil
}

func (m *OperationMetadata) GetProgressPercent() int32 {
	if m != nil {
		return m.ProgressPercent
	}
	return 0
}

func (m *OperationMetadata) GetPartialFailures() []*status.Status {
	if m != nil {
		return m.PartialFailures
	}
	return nil
}

func (m *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OperationMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OperationMetadata_CreateModelDetails)(nil),
	}
}

// Details of CreateModel operation.
type CreateModelOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateModelOperationMetadata) Reset()         { *m = CreateModelOperationMetadata{} }
func (m *CreateModelOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateModelOperationMetadata) ProtoMessage()    {}
func (*CreateModelOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_07afd87baa9bfac3, []int{1}
}

func (m *CreateModelOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModelOperationMetadata.Unmarshal(m, b)
}
func (m *CreateModelOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModelOperationMetadata.Marshal(b, m, deterministic)
}
func (m *CreateModelOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModelOperationMetadata.Merge(m, src)
}
func (m *CreateModelOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateModelOperationMetadata.Size(m)
}
func (m *CreateModelOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModelOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModelOperationMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.cloud.automl.v1beta1.OperationMetadata")
	proto.RegisterType((*CreateModelOperationMetadata)(nil), "google.cloud.automl.v1beta1.CreateModelOperationMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/operations.proto", fileDescriptor_07afd87baa9bfac3)
}

var fileDescriptor_07afd87baa9bfac3 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0xcd, 0xd6, 0x3f, 0x38, 0x8b, 0xb4, 0x06, 0xc1, 0xb0, 0x2d, 0x76, 0xe9, 0x8b, 0x2b,
	0xc8, 0x0c, 0x5d, 0x9f, 0xa4, 0xf8, 0xd0, 0x56, 0xd4, 0x97, 0xc5, 0x12, 0xc5, 0x07, 0x59, 0x08,
	0x77, 0x93, 0xbb, 0x21, 0x30, 0xc9, 0x0c, 0x33, 0x77, 0x16, 0x7c, 0xf7, 0xd3, 0xf9, 0x6d, 0xfc,
	0x06, 0x92, 0x99, 0x89, 0x8a, 0x96, 0xec, 0xe3, 0xdc, 0xfb, 0x3b, 0x27, 0xe7, 0x1e, 0xc2, 0x5e,
	0xd6, 0x4a, 0xd5, 0x12, 0x45, 0x29, 0x95, 0xab, 0x04, 0x38, 0x52, 0xad, 0x14, 0xbb, 0xf3, 0x0d,
	0x12, 0x9c, 0x0b, 0xa5, 0xd1, 0x00, 0x35, 0xaa, 0xb3, 0x5c, 0x1b, 0x45, 0x2a, 0x3d, 0x0e, 0x34,
	0xf7, 0x34, 0x0f, 0x34, 0x8f, 0xf4, 0xec, 0x24, 0x5a, 0x81, 0x6e, 0x04, 0x74, 0x9d, 0xa2, 0xbf,
	0xa5, 0xb3, 0xe7, 0x63, 0x1f, 0x6a, 0x55, 0x85, 0x32, 0x82, 0xcb, 0xbd, 0x60, 0x81, 0x3b, 0x90,
	0xce, 0xbb, 0x47, 0x4d, 0xcc, 0x25, 0xfc, 0x6b, 0xe3, 0xb6, 0x02, 0x5b, 0x4d, 0xdf, 0xe2, 0xf2,
	0xf4, 0xdf, 0x25, 0x35, 0x2d, 0x5a, 0x82, 0x56, 0x47, 0xe0, 0x69, 0x04, 0x8c, 0x2e, 0x85, 0x25,
	0x20, 0x17, 0x33, 0x9f, 0xfd, 0x9c, 0xb0, 0xc7, 0x1f, 0x87, 0x0e, 0x56, 0x48, 0x50, 0x01, 0x41,
	0xda, 0xb2, 0x27, 0xa5, 0x41, 0x20, 0x2c, 0x42, 0x9a, 0x0a, 0x09, 0x1a, 0x69, 0x33, 0x36, 0x4f,
	0x16, 0xd3, 0xe5, 0x6b, 0x3e, 0xd2, 0x11, 0xbf, 0xf6, 0xc2, 0x55, 0xaf, 0xfb, 0xcf, 0xf8, 0xc3,
	0x9d, 0x3c, 0x2d, 0xff, 0xec, 0xdf, 0x06, 0xdb, 0xf4, 0x05, 0x3b, 0xd2, 0x46, 0xd5, 0x06, 0xad,
	0x2d, 0x34, 0x9a, 0x12, 0x3b, 0xca, 0x1e, 0xcd, 0x93, 0xc5, 0xbd, 0xfc, 0x70, 0x98, 0xdf, 0x84,
	0x71, 0xfa, 0x86, 0x1d, 0x69, 0x30, 0xd4, 0x80, 0x2c, 0xb6, 0xd0, 0x48, 0x67, 0xd0, 0x66, 0x93,
	0xf9, 0xc1, 0x62, 0xba, 0x4c, 0x87, 0x54, 0x46, 0x97, 0xfc, 0x93, 0xbf, 0x31, 0x3f, 0x8c, 0xec,
	0xbb, 0x88, 0xa6, 0x17, 0x6c, 0x1a, 0x0f, 0xeb, 0x1b, 0xca, 0x0e, 0xfc, 0x3d, 0xb3, 0x41, 0x39,
	0xd4, 0xc7, 0x3f, 0x0f, 0xf5, 0xe5, 0x2c, 0xe0, 0xfd, 0xa0, 0x17, 0x3b, 0x5d, 0xfd, 0x16, 0xdf,
	0xdd, 0x2f, 0x0e, 0x78, 0x3f, 0xb8, 0x7a, 0xc8, 0x1e, 0xc4, 0x16, 0xcf, 0x9e, 0xb1, 0x93, 0xb1,
	0x92, 0xae, 0xbe, 0x27, 0xec, 0xb4, 0x54, 0xed, 0x58, 0xcb, 0x37, 0xc9, 0xd7, 0xcb, 0xb8, 0xae,
	0x95, 0x84, 0xae, 0xe6, 0xca, 0xd4, 0xa2, 0xc6, 0xce, 0xa7, 0x10, 0x61, 0x05, 0xba, 0xb1, 0xb7,
	0xfe, 0x63, 0x17, 0xe1, 0xf9, 0x63, 0x72, 0xfc, 0xde, 0x83, 0xeb, 0xeb, 0x1e, 0x5a, 0x5f, 0x3a,
	0x52, 0x2b, 0xb9, 0xfe, 0x12, 0xa0, 0xcd, 0x7d, 0xef, 0xf5, 0xea, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0x8b, 0x31, 0xeb, 0x40, 0x03, 0x00, 0x00,
}
