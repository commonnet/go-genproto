// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/remarketing_action_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// Request message for
// [RemarketingActionService.GetRemarketingAction][google.ads.googleads.v0.services.RemarketingActionService.GetRemarketingAction].
type GetRemarketingActionRequest struct {
	// The resource name of the remarketing action to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRemarketingActionRequest) Reset()         { *m = GetRemarketingActionRequest{} }
func (m *GetRemarketingActionRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemarketingActionRequest) ProtoMessage()    {}
func (*GetRemarketingActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8000a2ef14449c51, []int{0}
}

func (m *GetRemarketingActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemarketingActionRequest.Unmarshal(m, b)
}
func (m *GetRemarketingActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemarketingActionRequest.Marshal(b, m, deterministic)
}
func (m *GetRemarketingActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemarketingActionRequest.Merge(m, src)
}
func (m *GetRemarketingActionRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemarketingActionRequest.Size(m)
}
func (m *GetRemarketingActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemarketingActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemarketingActionRequest proto.InternalMessageInfo

func (m *GetRemarketingActionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [RemarketingActionService.MutateRemarketingActions][google.ads.googleads.v0.services.RemarketingActionService.MutateRemarketingActions].
type MutateRemarketingActionsRequest struct {
	// The ID of the customer whose remarketing actions are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual remarketing actions.
	Operations []*RemarketingActionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateRemarketingActionsRequest) Reset()         { *m = MutateRemarketingActionsRequest{} }
func (m *MutateRemarketingActionsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateRemarketingActionsRequest) ProtoMessage()    {}
func (*MutateRemarketingActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8000a2ef14449c51, []int{1}
}

func (m *MutateRemarketingActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRemarketingActionsRequest.Unmarshal(m, b)
}
func (m *MutateRemarketingActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRemarketingActionsRequest.Marshal(b, m, deterministic)
}
func (m *MutateRemarketingActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRemarketingActionsRequest.Merge(m, src)
}
func (m *MutateRemarketingActionsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateRemarketingActionsRequest.Size(m)
}
func (m *MutateRemarketingActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRemarketingActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRemarketingActionsRequest proto.InternalMessageInfo

func (m *MutateRemarketingActionsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateRemarketingActionsRequest) GetOperations() []*RemarketingActionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateRemarketingActionsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateRemarketingActionsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update) on a remarketing action.
type RemarketingActionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*RemarketingActionOperation_Create
	//	*RemarketingActionOperation_Update
	Operation            isRemarketingActionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RemarketingActionOperation) Reset()         { *m = RemarketingActionOperation{} }
func (m *RemarketingActionOperation) String() string { return proto.CompactTextString(m) }
func (*RemarketingActionOperation) ProtoMessage()    {}
func (*RemarketingActionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8000a2ef14449c51, []int{2}
}

func (m *RemarketingActionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingActionOperation.Unmarshal(m, b)
}
func (m *RemarketingActionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingActionOperation.Marshal(b, m, deterministic)
}
func (m *RemarketingActionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingActionOperation.Merge(m, src)
}
func (m *RemarketingActionOperation) XXX_Size() int {
	return xxx_messageInfo_RemarketingActionOperation.Size(m)
}
func (m *RemarketingActionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingActionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingActionOperation proto.InternalMessageInfo

func (m *RemarketingActionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isRemarketingActionOperation_Operation interface {
	isRemarketingActionOperation_Operation()
}

type RemarketingActionOperation_Create struct {
	Create *resources.RemarketingAction `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type RemarketingActionOperation_Update struct {
	Update *resources.RemarketingAction `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*RemarketingActionOperation_Create) isRemarketingActionOperation_Operation() {}

func (*RemarketingActionOperation_Update) isRemarketingActionOperation_Operation() {}

func (m *RemarketingActionOperation) GetOperation() isRemarketingActionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *RemarketingActionOperation) GetCreate() *resources.RemarketingAction {
	if x, ok := m.GetOperation().(*RemarketingActionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *RemarketingActionOperation) GetUpdate() *resources.RemarketingAction {
	if x, ok := m.GetOperation().(*RemarketingActionOperation_Update); ok {
		return x.Update
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RemarketingActionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RemarketingActionOperation_Create)(nil),
		(*RemarketingActionOperation_Update)(nil),
	}
}

// Response message for remarketing action mutate.
type MutateRemarketingActionsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateRemarketingActionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateRemarketingActionsResponse) Reset()         { *m = MutateRemarketingActionsResponse{} }
func (m *MutateRemarketingActionsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateRemarketingActionsResponse) ProtoMessage()    {}
func (*MutateRemarketingActionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8000a2ef14449c51, []int{3}
}

func (m *MutateRemarketingActionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRemarketingActionsResponse.Unmarshal(m, b)
}
func (m *MutateRemarketingActionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRemarketingActionsResponse.Marshal(b, m, deterministic)
}
func (m *MutateRemarketingActionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRemarketingActionsResponse.Merge(m, src)
}
func (m *MutateRemarketingActionsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateRemarketingActionsResponse.Size(m)
}
func (m *MutateRemarketingActionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRemarketingActionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRemarketingActionsResponse proto.InternalMessageInfo

func (m *MutateRemarketingActionsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateRemarketingActionsResponse) GetResults() []*MutateRemarketingActionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the remarketing action mutate.
type MutateRemarketingActionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateRemarketingActionResult) Reset()         { *m = MutateRemarketingActionResult{} }
func (m *MutateRemarketingActionResult) String() string { return proto.CompactTextString(m) }
func (*MutateRemarketingActionResult) ProtoMessage()    {}
func (*MutateRemarketingActionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8000a2ef14449c51, []int{4}
}

func (m *MutateRemarketingActionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateRemarketingActionResult.Unmarshal(m, b)
}
func (m *MutateRemarketingActionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateRemarketingActionResult.Marshal(b, m, deterministic)
}
func (m *MutateRemarketingActionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateRemarketingActionResult.Merge(m, src)
}
func (m *MutateRemarketingActionResult) XXX_Size() int {
	return xxx_messageInfo_MutateRemarketingActionResult.Size(m)
}
func (m *MutateRemarketingActionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateRemarketingActionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateRemarketingActionResult proto.InternalMessageInfo

func (m *MutateRemarketingActionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRemarketingActionRequest)(nil), "google.ads.googleads.v0.services.GetRemarketingActionRequest")
	proto.RegisterType((*MutateRemarketingActionsRequest)(nil), "google.ads.googleads.v0.services.MutateRemarketingActionsRequest")
	proto.RegisterType((*RemarketingActionOperation)(nil), "google.ads.googleads.v0.services.RemarketingActionOperation")
	proto.RegisterType((*MutateRemarketingActionsResponse)(nil), "google.ads.googleads.v0.services.MutateRemarketingActionsResponse")
	proto.RegisterType((*MutateRemarketingActionResult)(nil), "google.ads.googleads.v0.services.MutateRemarketingActionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/remarketing_action_service.proto", fileDescriptor_8000a2ef14449c51)
}

var fileDescriptor_8000a2ef14449c51 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0xc7, 0x4d, 0x56, 0xaa, 0x9d, 0x54, 0x85, 0x51, 0x31, 0xac, 0xd6, 0x2e, 0xb1, 0x60, 0xd9,
	0x43, 0xb2, 0xac, 0x45, 0x21, 0x6d, 0x91, 0x2c, 0xda, 0xd6, 0x43, 0xff, 0x90, 0x42, 0x41, 0x59,
	0x08, 0xd3, 0x64, 0x1a, 0x42, 0x93, 0x4c, 0x9c, 0x99, 0xac, 0x94, 0xd2, 0x8b, 0x88, 0x2f, 0xe0,
	0x1b, 0x78, 0xf4, 0x3d, 0x3c, 0xe8, 0xc1, 0x8b, 0xaf, 0xa0, 0x17, 0x0f, 0x3e, 0x83, 0x24, 0x93,
	0xd9, 0xfe, 0xd9, 0xa6, 0x2b, 0xed, 0x6d, 0xf6, 0x37, 0xdf, 0xfd, 0xfc, 0xfe, 0x7c, 0x67, 0x26,
	0xc0, 0x09, 0x09, 0x09, 0x63, 0x6c, 0xa1, 0x80, 0x59, 0x62, 0x59, 0xac, 0x06, 0x1d, 0x8b, 0x61,
	0x3a, 0x88, 0x7c, 0xcc, 0x2c, 0x8a, 0x13, 0x44, 0xf7, 0x30, 0x8f, 0xd2, 0xd0, 0x43, 0x3e, 0x8f,
	0x48, 0xea, 0x55, 0x7b, 0x66, 0x46, 0x09, 0x27, 0xb0, 0x25, 0xfe, 0x67, 0xa2, 0x80, 0x99, 0x43,
	0x84, 0x39, 0xe8, 0x98, 0x12, 0xd1, 0xb4, 0xeb, 0x92, 0x50, 0xcc, 0x48, 0x4e, 0xcf, 0xce, 0x22,
	0xe8, 0xcd, 0x07, 0xf2, 0xbf, 0x59, 0x64, 0xa1, 0x34, 0x25, 0x1c, 0x15, 0x9b, 0xac, 0xda, 0xad,
	0x72, 0x5b, 0xe5, 0xaf, 0x9d, 0x7c, 0xd7, 0xda, 0x8d, 0x70, 0x1c, 0x78, 0x09, 0x62, 0x7b, 0x95,
	0xe2, 0xe1, 0x69, 0xc5, 0x3b, 0x8a, 0xb2, 0x0c, 0x53, 0x49, 0xb8, 0x57, 0xed, 0xd3, 0xcc, 0xb7,
	0x18, 0x47, 0x3c, 0xaf, 0x36, 0x8c, 0x1e, 0xb8, 0xbf, 0x82, 0xb9, 0x7b, 0x54, 0x97, 0x53, 0x96,
	0xe5, 0xe2, 0xb7, 0x39, 0x66, 0x1c, 0x3e, 0x02, 0x37, 0x64, 0xf5, 0x5e, 0x8a, 0x12, 0xac, 0x2b,
	0x2d, 0x65, 0x6e, 0xd2, 0x9d, 0x92, 0xc1, 0x75, 0x94, 0x60, 0xe3, 0xaf, 0x02, 0x66, 0xd6, 0x72,
	0x8e, 0x38, 0x1e, 0xe1, 0x30, 0x09, 0x9a, 0x01, 0x9a, 0x9f, 0x33, 0x4e, 0x12, 0x4c, 0xbd, 0x28,
	0xa8, 0x30, 0x40, 0x86, 0x5e, 0x05, 0xb0, 0x0f, 0x00, 0xc9, 0x30, 0x15, 0x7d, 0xeb, 0x6a, 0xab,
	0x31, 0xa7, 0x75, 0x17, 0xcd, 0x71, 0x43, 0x37, 0x47, 0x32, 0x6e, 0x48, 0x88, 0x7b, 0x8c, 0x07,
	0x1f, 0x83, 0x5b, 0x19, 0xa2, 0x3c, 0x42, 0xb1, 0xb7, 0x8b, 0xa2, 0x38, 0xa7, 0x58, 0x6f, 0xb4,
	0x94, 0xb9, 0xeb, 0xee, 0xcd, 0x2a, 0xbc, 0x2c, 0xa2, 0x45, 0xc3, 0x03, 0x14, 0x47, 0x01, 0xe2,
	0xd8, 0x23, 0x69, 0xbc, 0xaf, 0x5f, 0x2d, 0x65, 0x53, 0x32, 0xb8, 0x91, 0xc6, 0xfb, 0xc6, 0x47,
	0x15, 0x34, 0xeb, 0x13, 0xc3, 0x05, 0xa0, 0xe5, 0x59, 0x49, 0x28, 0x1c, 0x2a, 0x09, 0x5a, 0xb7,
	0x29, 0x7b, 0x91, 0x16, 0x99, 0xcb, 0x85, 0x89, 0x6b, 0x88, 0xed, 0xb9, 0x40, 0xc8, 0x8b, 0x35,
	0x5c, 0x07, 0x13, 0x3e, 0xc5, 0x88, 0x8b, 0x51, 0x6b, 0xdd, 0xf9, 0xda, 0x19, 0x0c, 0x8f, 0xd5,
	0xe8, 0x10, 0x56, 0xaf, 0xb8, 0x15, 0xa5, 0xe0, 0x09, 0xba, 0xae, 0x5e, 0x8e, 0x27, 0x28, 0x3d,
	0x0d, 0x4c, 0x0e, 0xe7, 0x6a, 0x7c, 0x55, 0x40, 0xab, 0xde, 0x79, 0x96, 0x91, 0x94, 0x61, 0xb8,
	0x0c, 0xee, 0x9e, 0x9a, 0xbd, 0x87, 0x29, 0x25, 0xb4, 0x74, 0x40, 0xeb, 0x42, 0x59, 0x10, 0xcd,
	0x7c, 0x73, 0xab, 0x3c, 0x9b, 0xee, 0xed, 0x93, 0xae, 0xbc, 0x2c, 0xe4, 0xf0, 0x35, 0xb8, 0x46,
	0x31, 0xcb, 0x63, 0x2e, 0x8f, 0xc7, 0xf3, 0xf1, 0xc7, 0xa3, 0xa6, 0x38, 0xb7, 0xe4, 0xb8, 0x92,
	0x67, 0xbc, 0x00, 0xd3, 0xe7, 0x2a, 0xff, 0xeb, 0x1e, 0x74, 0x7f, 0x34, 0x80, 0x3e, 0x02, 0xd8,
	0x12, 0xa5, 0xc0, 0x6f, 0x0a, 0xb8, 0x73, 0xd6, 0x4d, 0x83, 0x4b, 0xe3, 0xbb, 0x38, 0xe7, 0x86,
	0x36, 0x2f, 0xe4, 0xa7, 0xb1, 0xf8, 0xfe, 0xe7, 0xaf, 0x4f, 0xea, 0x53, 0x38, 0x5f, 0xbc, 0x4f,
	0x07, 0x27, 0x5a, 0x5b, 0x92, 0x97, 0x92, 0x59, 0xed, 0xe3, 0x0f, 0x56, 0x65, 0xab, 0xd5, 0x3e,
	0x84, 0xbf, 0x15, 0xa0, 0xd7, 0xd9, 0x0e, 0x9d, 0x0b, 0xbb, 0x22, 0x1f, 0x8b, 0x66, 0xef, 0x32,
	0x08, 0x71, 0xea, 0x8c, 0x5e, 0xd9, 0xe1, 0xa2, 0xf1, 0xac, 0xe8, 0xf0, 0xa8, 0xa5, 0x83, 0x63,
	0xaf, 0xd0, 0x52, 0xfb, 0xf0, 0x8c, 0x06, 0xed, 0xa4, 0x44, 0xdb, 0x4a, 0xbb, 0xf7, 0x41, 0x05,
	0xb3, 0x3e, 0x49, 0xc6, 0x56, 0xd3, 0x9b, 0xae, 0xb3, 0x7d, 0xb3, 0xb8, 0xec, 0x9b, 0xca, 0x9b,
	0xd5, 0x0a, 0x11, 0x92, 0x18, 0xa5, 0xa1, 0x49, 0x68, 0x68, 0x85, 0x38, 0x2d, 0x9f, 0x02, 0xf9,
	0xad, 0xc8, 0x22, 0x56, 0xff, 0x7d, 0x5a, 0x90, 0x8b, 0xcf, 0x6a, 0x63, 0xc5, 0x71, 0xbe, 0xa8,
	0xad, 0x15, 0x01, 0x74, 0x02, 0x66, 0x8a, 0x65, 0xb1, 0xda, 0xee, 0x98, 0x55, 0x62, 0xf6, 0x5d,
	0x4a, 0xfa, 0x4e, 0xc0, 0xfa, 0x43, 0x49, 0x7f, 0xbb, 0xd3, 0x97, 0x92, 0x3f, 0xea, 0xac, 0x88,
	0xdb, 0xb6, 0x13, 0x30, 0xdb, 0x1e, 0x8a, 0x6c, 0x7b, 0xbb, 0x63, 0xdb, 0x52, 0xb6, 0x33, 0x51,
	0xd6, 0xf9, 0xe4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x57, 0x2e, 0x59, 0x46, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemarketingActionServiceClient is the client API for RemarketingActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemarketingActionServiceClient interface {
	// Returns the requested remarketing action in full detail.
	GetRemarketingAction(ctx context.Context, in *GetRemarketingActionRequest, opts ...grpc.CallOption) (*resources.RemarketingAction, error)
	// Creates or updates remarketing actions. Operation statuses are returned.
	MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error)
}

type remarketingActionServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemarketingActionServiceClient(cc *grpc.ClientConn) RemarketingActionServiceClient {
	return &remarketingActionServiceClient{cc}
}

func (c *remarketingActionServiceClient) GetRemarketingAction(ctx context.Context, in *GetRemarketingActionRequest, opts ...grpc.CallOption) (*resources.RemarketingAction, error) {
	out := new(resources.RemarketingAction)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.RemarketingActionService/GetRemarketingAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remarketingActionServiceClient) MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error) {
	out := new(MutateRemarketingActionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.RemarketingActionService/MutateRemarketingActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemarketingActionServiceServer is the server API for RemarketingActionService service.
type RemarketingActionServiceServer interface {
	// Returns the requested remarketing action in full detail.
	GetRemarketingAction(context.Context, *GetRemarketingActionRequest) (*resources.RemarketingAction, error)
	// Creates or updates remarketing actions. Operation statuses are returned.
	MutateRemarketingActions(context.Context, *MutateRemarketingActionsRequest) (*MutateRemarketingActionsResponse, error)
}

func RegisterRemarketingActionServiceServer(s *grpc.Server, srv RemarketingActionServiceServer) {
	s.RegisterService(&_RemarketingActionService_serviceDesc, srv)
}

func _RemarketingActionService_GetRemarketingAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemarketingActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarketingActionServiceServer).GetRemarketingAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.RemarketingActionService/GetRemarketingAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarketingActionServiceServer).GetRemarketingAction(ctx, req.(*GetRemarketingActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemarketingActionService_MutateRemarketingActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRemarketingActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.RemarketingActionService/MutateRemarketingActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, req.(*MutateRemarketingActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemarketingActionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.RemarketingActionService",
	HandlerType: (*RemarketingActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRemarketingAction",
			Handler:    _RemarketingActionService_GetRemarketingAction_Handler,
		},
		{
			MethodName: "MutateRemarketingActions",
			Handler:    _RemarketingActionService_MutateRemarketingActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/remarketing_action_service.proto",
}
