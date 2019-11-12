// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/functions/v1beta2/operations.proto

package functions

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A type of an operation.
type OperationType int32

const (
	// Unknown operation type.
	OperationType_OPERATION_UNSPECIFIED OperationType = 0
	// Triggered by CreateFunction call
	OperationType_CREATE_FUNCTION OperationType = 1
	// Triggered by UpdateFunction call
	OperationType_UPDATE_FUNCTION OperationType = 2
	// Triggered by DeleteFunction call.
	OperationType_DELETE_FUNCTION OperationType = 3
)

var OperationType_name = map[int32]string{
	0: "OPERATION_UNSPECIFIED",
	1: "CREATE_FUNCTION",
	2: "UPDATE_FUNCTION",
	3: "DELETE_FUNCTION",
}

var OperationType_value = map[string]int32{
	"OPERATION_UNSPECIFIED": 0,
	"CREATE_FUNCTION":       1,
	"UPDATE_FUNCTION":       2,
	"DELETE_FUNCTION":       3,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25e43fb62395b1f8, []int{0}
}

// Metadata describing an [Operation][google.longrunning.Operation]
type OperationMetadataV1Beta2 struct {
	// Target of the operation - for example
	// projects/project-1/locations/region-1/functions/function-1
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Type of operation.
	Type OperationType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.functions.v1beta2.OperationType" json:"type,omitempty"`
	// The original request that started the operation.
	Request *any.Any `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// Version id of the function created or updated by an API call.
	// This field is only populated for Create and Update operations.
	VersionId int64 `protobuf:"varint,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// The last update timestamp of the operation.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OperationMetadataV1Beta2) Reset()         { *m = OperationMetadataV1Beta2{} }
func (m *OperationMetadataV1Beta2) String() string { return proto.CompactTextString(m) }
func (*OperationMetadataV1Beta2) ProtoMessage()    {}
func (*OperationMetadataV1Beta2) Descriptor() ([]byte, []int) {
	return fileDescriptor_25e43fb62395b1f8, []int{0}
}

func (m *OperationMetadataV1Beta2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationMetadataV1Beta2.Unmarshal(m, b)
}
func (m *OperationMetadataV1Beta2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationMetadataV1Beta2.Marshal(b, m, deterministic)
}
func (m *OperationMetadataV1Beta2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationMetadataV1Beta2.Merge(m, src)
}
func (m *OperationMetadataV1Beta2) XXX_Size() int {
	return xxx_messageInfo_OperationMetadataV1Beta2.Size(m)
}
func (m *OperationMetadataV1Beta2) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationMetadataV1Beta2.DiscardUnknown(m)
}

var xxx_messageInfo_OperationMetadataV1Beta2 proto.InternalMessageInfo

func (m *OperationMetadataV1Beta2) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *OperationMetadataV1Beta2) GetType() OperationType {
	if m != nil {
		return m.Type
	}
	return OperationType_OPERATION_UNSPECIFIED
}

func (m *OperationMetadataV1Beta2) GetRequest() *any.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *OperationMetadataV1Beta2) GetVersionId() int64 {
	if m != nil {
		return m.VersionId
	}
	return 0
}

func (m *OperationMetadataV1Beta2) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.functions.v1beta2.OperationType", OperationType_name, OperationType_value)
	proto.RegisterType((*OperationMetadataV1Beta2)(nil), "google.cloud.functions.v1beta2.OperationMetadataV1Beta2")
}

func init() {
	proto.RegisterFile("google/cloud/functions/v1beta2/operations.proto", fileDescriptor_25e43fb62395b1f8)
}

var fileDescriptor_25e43fb62395b1f8 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0xcb, 0xda, 0x30,
	0x1c, 0xc7, 0x17, 0x75, 0x0e, 0x23, 0xdb, 0x24, 0x7b, 0xa1, 0xca, 0x5e, 0x8a, 0xa7, 0x32, 0x58,
	0x82, 0xee, 0xe8, 0xa9, 0x6a, 0x1d, 0x85, 0x4d, 0x4b, 0x57, 0x77, 0xd8, 0xa5, 0x44, 0x1b, 0x4b,
	0x41, 0x93, 0xac, 0x4d, 0x85, 0x1e, 0x77, 0xdc, 0x7f, 0x3d, 0x5a, 0xd3, 0xa2, 0x7b, 0x78, 0x9e,
	0xe7, 0xd8, 0x4f, 0xbf, 0x9f, 0x6f, 0x7e, 0xbf, 0x10, 0x48, 0x62, 0x21, 0xe2, 0x23, 0x23, 0xfb,
	0xa3, 0xc8, 0x23, 0x72, 0xc8, 0xf9, 0x5e, 0x25, 0x82, 0x67, 0xe4, 0x3c, 0xd9, 0x31, 0x45, 0xa7,
	0x44, 0x48, 0x96, 0xd2, 0x0a, 0x61, 0x99, 0x0a, 0x25, 0xd0, 0x87, 0x8b, 0x80, 0x2b, 0x01, 0x37,
	0x02, 0xd6, 0xc2, 0x68, 0xa8, 0x0b, 0xab, 0xf4, 0x2e, 0x3f, 0x10, 0xca, 0x8b, 0x8b, 0x3a, 0xfa,
	0xf8, 0xff, 0x2f, 0x95, 0x9c, 0x58, 0xa6, 0xe8, 0x49, 0xea, 0xc0, 0x3b, 0x1d, 0xa0, 0x32, 0x21,
	0x94, 0x73, 0xa1, 0xae, 0x4f, 0x1e, 0xff, 0x69, 0x41, 0x63, 0x53, 0x8f, 0xf3, 0x9d, 0x29, 0x1a,
	0x51, 0x45, 0x7f, 0x4e, 0xe6, 0xe5, 0xb1, 0xe8, 0x2d, 0xec, 0x2a, 0x9a, 0xc6, 0x4c, 0x19, 0xc0,
	0x04, 0x56, 0xcf, 0xd7, 0x5f, 0xc8, 0x86, 0x1d, 0x55, 0x48, 0x66, 0xb4, 0x4c, 0x60, 0xbd, 0x98,
	0x7e, 0xc6, 0x0f, 0x4f, 0x8f, 0x9b, 0xfe, 0xa0, 0x90, 0xcc, 0xaf, 0x54, 0x84, 0xe1, 0xb3, 0x94,
	0xfd, 0xce, 0x59, 0xa6, 0x8c, 0xb6, 0x09, 0xac, 0xfe, 0xf4, 0x75, 0xdd, 0x52, 0x2f, 0x82, 0x6d,
	0x5e, 0xf8, 0x75, 0x08, 0xbd, 0x87, 0xf0, 0xcc, 0xd2, 0x2c, 0x11, 0x3c, 0x4c, 0x22, 0xa3, 0x63,
	0x02, 0xab, 0xed, 0xf7, 0x34, 0x71, 0x23, 0x34, 0x83, 0xfd, 0x5c, 0x46, 0x54, 0xb1, 0xb0, 0x5c,
	0xdf, 0x78, 0x5a, 0x55, 0x8e, 0xee, 0x54, 0x06, 0xf5, 0xdd, 0xf8, 0xf0, 0x12, 0x2f, 0xc1, 0xa7,
	0x04, 0x3e, 0xbf, 0x19, 0x11, 0x0d, 0xe1, 0x9b, 0x8d, 0xe7, 0xf8, 0x76, 0xe0, 0x6e, 0xd6, 0xe1,
	0x76, 0xfd, 0xc3, 0x73, 0x16, 0xee, 0xca, 0x75, 0x96, 0x83, 0x27, 0xe8, 0x15, 0x7c, 0xb9, 0xf0,
	0x1d, 0x3b, 0x70, 0xc2, 0xd5, 0x76, 0xbd, 0x28, 0x03, 0x03, 0x50, 0xc2, 0xad, 0xb7, 0xbc, 0x81,
	0xad, 0x12, 0x2e, 0x9d, 0x6f, 0xce, 0x35, 0x6c, 0xcf, 0xff, 0x02, 0x38, 0xde, 0x8b, 0xd3, 0x23,
	0x37, 0x36, 0x37, 0x56, 0x35, 0x6a, 0x06, 0xcb, 0xbc, 0x72, 0x09, 0x0f, 0xfc, 0xfa, 0xaa, 0xdd,
	0x58, 0x1c, 0x29, 0x8f, 0xb1, 0x48, 0x63, 0x12, 0x33, 0x5e, 0xad, 0xa8, 0xdf, 0x1d, 0x95, 0x49,
	0x76, 0xdf, 0xdb, 0x9b, 0x35, 0x64, 0xd7, 0xad, 0x9c, 0x2f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x4b, 0xf7, 0x57, 0x72, 0xae, 0x02, 0x00, 0x00,
}
