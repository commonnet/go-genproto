// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_shared_set_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Request message for [CampaignSharedSetService.GetCampaignSharedSet][google.ads.googleads.v1.services.CampaignSharedSetService.GetCampaignSharedSet].
type GetCampaignSharedSetRequest struct {
	// The resource name of the campaign shared set to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignSharedSetRequest) Reset()         { *m = GetCampaignSharedSetRequest{} }
func (m *GetCampaignSharedSetRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignSharedSetRequest) ProtoMessage()    {}
func (*GetCampaignSharedSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eb5d89b68dc46d0, []int{0}
}

func (m *GetCampaignSharedSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignSharedSetRequest.Unmarshal(m, b)
}
func (m *GetCampaignSharedSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignSharedSetRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignSharedSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignSharedSetRequest.Merge(m, src)
}
func (m *GetCampaignSharedSetRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignSharedSetRequest.Size(m)
}
func (m *GetCampaignSharedSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignSharedSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignSharedSetRequest proto.InternalMessageInfo

func (m *GetCampaignSharedSetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignSharedSetService.MutateCampaignSharedSets][google.ads.googleads.v1.services.CampaignSharedSetService.MutateCampaignSharedSets].
type MutateCampaignSharedSetsRequest struct {
	// The ID of the customer whose campaign shared sets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign shared sets.
	Operations []*CampaignSharedSetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCampaignSharedSetsRequest) Reset()         { *m = MutateCampaignSharedSetsRequest{} }
func (m *MutateCampaignSharedSetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignSharedSetsRequest) ProtoMessage()    {}
func (*MutateCampaignSharedSetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eb5d89b68dc46d0, []int{1}
}

func (m *MutateCampaignSharedSetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignSharedSetsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignSharedSetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignSharedSetsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignSharedSetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignSharedSetsRequest.Merge(m, src)
}
func (m *MutateCampaignSharedSetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignSharedSetsRequest.Size(m)
}
func (m *MutateCampaignSharedSetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignSharedSetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignSharedSetsRequest proto.InternalMessageInfo

func (m *MutateCampaignSharedSetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignSharedSetsRequest) GetOperations() []*CampaignSharedSetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignSharedSetsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignSharedSetsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an campaign shared set.
type CampaignSharedSetOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignSharedSetOperation_Create
	//	*CampaignSharedSetOperation_Remove
	Operation            isCampaignSharedSetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CampaignSharedSetOperation) Reset()         { *m = CampaignSharedSetOperation{} }
func (m *CampaignSharedSetOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetOperation) ProtoMessage()    {}
func (*CampaignSharedSetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eb5d89b68dc46d0, []int{2}
}

func (m *CampaignSharedSetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetOperation.Unmarshal(m, b)
}
func (m *CampaignSharedSetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetOperation.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetOperation.Merge(m, src)
}
func (m *CampaignSharedSetOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetOperation.Size(m)
}
func (m *CampaignSharedSetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetOperation proto.InternalMessageInfo

type isCampaignSharedSetOperation_Operation interface {
	isCampaignSharedSetOperation_Operation()
}

type CampaignSharedSetOperation_Create struct {
	Create *resources.CampaignSharedSet `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignSharedSetOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignSharedSetOperation_Create) isCampaignSharedSetOperation_Operation() {}

func (*CampaignSharedSetOperation_Remove) isCampaignSharedSetOperation_Operation() {}

func (m *CampaignSharedSetOperation) GetOperation() isCampaignSharedSetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignSharedSetOperation) GetCreate() *resources.CampaignSharedSet {
	if x, ok := m.GetOperation().(*CampaignSharedSetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignSharedSetOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignSharedSetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignSharedSetOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignSharedSetOperation_Create)(nil),
		(*CampaignSharedSetOperation_Remove)(nil),
	}
}

// Response message for a campaign shared set mutate.
type MutateCampaignSharedSetsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignSharedSetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateCampaignSharedSetsResponse) Reset()         { *m = MutateCampaignSharedSetsResponse{} }
func (m *MutateCampaignSharedSetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignSharedSetsResponse) ProtoMessage()    {}
func (*MutateCampaignSharedSetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eb5d89b68dc46d0, []int{3}
}

func (m *MutateCampaignSharedSetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignSharedSetsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignSharedSetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignSharedSetsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignSharedSetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignSharedSetsResponse.Merge(m, src)
}
func (m *MutateCampaignSharedSetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignSharedSetsResponse.Size(m)
}
func (m *MutateCampaignSharedSetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignSharedSetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignSharedSetsResponse proto.InternalMessageInfo

func (m *MutateCampaignSharedSetsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignSharedSetsResponse) GetResults() []*MutateCampaignSharedSetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign shared set mutate.
type MutateCampaignSharedSetResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignSharedSetResult) Reset()         { *m = MutateCampaignSharedSetResult{} }
func (m *MutateCampaignSharedSetResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignSharedSetResult) ProtoMessage()    {}
func (*MutateCampaignSharedSetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eb5d89b68dc46d0, []int{4}
}

func (m *MutateCampaignSharedSetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignSharedSetResult.Unmarshal(m, b)
}
func (m *MutateCampaignSharedSetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignSharedSetResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignSharedSetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignSharedSetResult.Merge(m, src)
}
func (m *MutateCampaignSharedSetResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignSharedSetResult.Size(m)
}
func (m *MutateCampaignSharedSetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignSharedSetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignSharedSetResult proto.InternalMessageInfo

func (m *MutateCampaignSharedSetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignSharedSetRequest)(nil), "google.ads.googleads.v1.services.GetCampaignSharedSetRequest")
	proto.RegisterType((*MutateCampaignSharedSetsRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignSharedSetsRequest")
	proto.RegisterType((*CampaignSharedSetOperation)(nil), "google.ads.googleads.v1.services.CampaignSharedSetOperation")
	proto.RegisterType((*MutateCampaignSharedSetsResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignSharedSetsResponse")
	proto.RegisterType((*MutateCampaignSharedSetResult)(nil), "google.ads.googleads.v1.services.MutateCampaignSharedSetResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_shared_set_service.proto", fileDescriptor_4eb5d89b68dc46d0)
}

var fileDescriptor_4eb5d89b68dc46d0 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x41, 0x6b, 0xd4, 0x4e,
	0x14, 0xff, 0x67, 0xf7, 0x4f, 0xb5, 0xb3, 0x55, 0x61, 0x54, 0x0c, 0xab, 0xb5, 0x4b, 0x2c, 0x58,
	0xf6, 0x90, 0xb0, 0x6b, 0x51, 0x48, 0x5b, 0x24, 0x51, 0xdb, 0x7a, 0xb0, 0x2d, 0x59, 0x58, 0x50,
	0x16, 0xc2, 0x34, 0x99, 0xc6, 0x40, 0x92, 0x89, 0x33, 0x93, 0x95, 0x52, 0x7a, 0x11, 0xbf, 0x41,
	0xbf, 0x81, 0x47, 0xbf, 0x87, 0x07, 0x3d, 0x78, 0xf1, 0x2b, 0xe8, 0xc5, 0x83, 0x9f, 0x41, 0x92,
	0xc9, 0xa4, 0xad, 0xdb, 0x74, 0xa5, 0xde, 0x5e, 0xde, 0x7b, 0xf3, 0x7b, 0xef, 0x37, 0xbf, 0x37,
	0x2f, 0xc0, 0x0e, 0x08, 0x09, 0x22, 0x6c, 0x20, 0x9f, 0x19, 0xc2, 0xcc, 0xad, 0x71, 0xcf, 0x60,
	0x98, 0x8e, 0x43, 0x0f, 0x33, 0xc3, 0x43, 0x71, 0x8a, 0xc2, 0x20, 0x71, 0xd9, 0x6b, 0x44, 0xb1,
	0xef, 0x32, 0xcc, 0xdd, 0x32, 0xa8, 0xa7, 0x94, 0x70, 0x02, 0x3b, 0xe2, 0xa0, 0x8e, 0x7c, 0xa6,
	0x57, 0x18, 0xfa, 0xb8, 0xa7, 0x4b, 0x8c, 0xf6, 0x4a, 0x5d, 0x15, 0x8a, 0x19, 0xc9, 0x68, 0x4d,
	0x19, 0x01, 0xdf, 0xbe, 0x23, 0x0f, 0xa7, 0xa1, 0x81, 0x92, 0x84, 0x70, 0xc4, 0x43, 0x92, 0xb0,
	0x32, 0x7a, 0xb7, 0x8c, 0x16, 0x5f, 0xbb, 0xd9, 0x9e, 0xf1, 0x96, 0xa2, 0x34, 0xc5, 0x54, 0xc6,
	0x6f, 0x95, 0x71, 0x9a, 0x7a, 0x06, 0xe3, 0x88, 0x67, 0x65, 0x40, 0xb3, 0xc1, 0xed, 0x0d, 0xcc,
	0x9f, 0x94, 0x65, 0x07, 0x45, 0xd5, 0x01, 0xe6, 0x0e, 0x7e, 0x93, 0x61, 0xc6, 0xe1, 0x3d, 0x70,
	0x45, 0x36, 0xe7, 0x26, 0x28, 0xc6, 0xaa, 0xd2, 0x51, 0x96, 0x66, 0x9d, 0x39, 0xe9, 0xdc, 0x42,
	0x31, 0xd6, 0x7e, 0x29, 0x60, 0xe1, 0x45, 0xc6, 0x11, 0xc7, 0x13, 0x38, 0x4c, 0x02, 0x2d, 0x80,
	0x96, 0x97, 0x31, 0x4e, 0x62, 0x4c, 0xdd, 0xd0, 0x2f, 0x61, 0x80, 0x74, 0x3d, 0xf7, 0xe1, 0x08,
	0x00, 0x92, 0x62, 0x2a, 0x58, 0xa9, 0x8d, 0x4e, 0x73, 0xa9, 0xd5, 0x5f, 0xd5, 0xa7, 0xdd, 0xa9,
	0x3e, 0x51, 0x71, 0x5b, 0x82, 0x38, 0x27, 0xf0, 0xe0, 0x7d, 0x70, 0x2d, 0x45, 0x94, 0x87, 0x28,
	0x72, 0xf7, 0x50, 0x18, 0x65, 0x14, 0xab, 0xcd, 0x8e, 0xb2, 0x74, 0xd9, 0xb9, 0x5a, 0xba, 0xd7,
	0x85, 0x37, 0x27, 0x3c, 0x46, 0x51, 0xe8, 0x23, 0x8e, 0x5d, 0x92, 0x44, 0xfb, 0xea, 0xff, 0x45,
	0xda, 0x9c, 0x74, 0x6e, 0x27, 0xd1, 0xbe, 0x76, 0xa4, 0x80, 0x76, 0x7d, 0x61, 0xb8, 0x05, 0x66,
	0x3c, 0x8a, 0x11, 0x17, 0xb7, 0xd5, 0xea, 0x2f, 0xd7, 0xd2, 0xa8, 0x84, 0x9f, 0xe4, 0xb1, 0xf9,
	0x9f, 0x53, 0xa2, 0x40, 0x15, 0xcc, 0x50, 0x1c, 0x93, 0xb1, 0xe8, 0x79, 0x36, 0x8f, 0x88, 0x6f,
	0xbb, 0x05, 0x66, 0x2b, 0x92, 0xda, 0x27, 0x05, 0x74, 0xea, 0x65, 0x60, 0x29, 0x49, 0x18, 0x86,
	0xeb, 0xe0, 0xe6, 0x1f, 0x17, 0xe1, 0x62, 0x4a, 0x09, 0x2d, 0xa0, 0x5b, 0x7d, 0x28, 0x5b, 0xa5,
	0xa9, 0xa7, 0x0f, 0x8a, 0x41, 0x71, 0xae, 0x9f, 0xbe, 0xa2, 0x67, 0x79, 0x3a, 0x7c, 0x09, 0x2e,
	0x51, 0xcc, 0xb2, 0x88, 0x4b, 0xad, 0x1e, 0x4f, 0xd7, 0xaa, 0xa6, 0x39, 0xa7, 0xc0, 0x71, 0x24,
	0x9e, 0xf6, 0x14, 0xcc, 0x9f, 0x9b, 0xf9, 0x57, 0x43, 0xd9, 0xff, 0xda, 0x04, 0xea, 0x04, 0xc0,
	0x40, 0xb4, 0x02, 0x3f, 0x2b, 0xe0, 0xc6, 0x59, 0x63, 0x0f, 0xd7, 0xa6, 0xb3, 0x38, 0xe7, 0xb9,
	0xb4, 0x2f, 0xa4, 0xb4, 0xb6, 0xfa, 0xee, 0xdb, 0xf7, 0xa3, 0xc6, 0x43, 0xb8, 0x9c, 0xef, 0x82,
	0x83, 0x53, 0xd4, 0xd6, 0xe4, 0x0b, 0x61, 0x46, 0xb7, 0x5a, 0x0e, 0xc7, 0xb2, 0x1a, 0xdd, 0x43,
	0xf8, 0x43, 0x01, 0x6a, 0x9d, 0xec, 0xd0, 0xba, 0xb0, 0x2a, 0xf2, 0xe5, 0xb6, 0xed, 0x7f, 0x81,
	0x10, 0x53, 0xa7, 0xd9, 0x05, 0xc3, 0x55, 0xed, 0x51, 0xce, 0xf0, 0x98, 0xd2, 0xc1, 0x89, 0x95,
	0xb0, 0xd6, 0x3d, 0x3c, 0x83, 0xa0, 0x19, 0x17, 0xd0, 0xa6, 0xd2, 0xb5, 0xdf, 0x37, 0xc0, 0xa2,
	0x47, 0xe2, 0xa9, 0xdd, 0xd8, 0xf3, 0x75, 0xb2, 0xef, 0xe4, 0x1b, 0x6f, 0x47, 0x79, 0xb5, 0x59,
	0x42, 0x04, 0x24, 0x42, 0x49, 0xa0, 0x13, 0x1a, 0x18, 0x01, 0x4e, 0x8a, 0x7d, 0x28, 0xf7, 0x72,
	0x1a, 0xb2, 0xfa, 0x9f, 0xc1, 0x8a, 0x34, 0x3e, 0x34, 0x9a, 0x1b, 0x96, 0xf5, 0xb1, 0xd1, 0xd9,
	0x10, 0x80, 0x96, 0xcf, 0x74, 0x61, 0xe6, 0xd6, 0xb0, 0xa7, 0x97, 0x85, 0xd9, 0x17, 0x99, 0x32,
	0xb2, 0x7c, 0x36, 0xaa, 0x52, 0x46, 0xc3, 0xde, 0x48, 0xa6, 0xfc, 0x6c, 0x2c, 0x0a, 0xbf, 0x69,
	0x5a, 0x3e, 0x33, 0xcd, 0x2a, 0xc9, 0x34, 0x87, 0x3d, 0xd3, 0x94, 0x69, 0xbb, 0x33, 0x45, 0x9f,
	0x0f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x69, 0x26, 0xa6, 0xb3, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignSharedSetServiceClient is the client API for CampaignSharedSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignSharedSetServiceClient interface {
	// Returns the requested campaign shared set in full detail.
	GetCampaignSharedSet(ctx context.Context, in *GetCampaignSharedSetRequest, opts ...grpc.CallOption) (*resources.CampaignSharedSet, error)
	// Creates or removes campaign shared sets. Operation statuses are returned.
	MutateCampaignSharedSets(ctx context.Context, in *MutateCampaignSharedSetsRequest, opts ...grpc.CallOption) (*MutateCampaignSharedSetsResponse, error)
}

type campaignSharedSetServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignSharedSetServiceClient(cc *grpc.ClientConn) CampaignSharedSetServiceClient {
	return &campaignSharedSetServiceClient{cc}
}

func (c *campaignSharedSetServiceClient) GetCampaignSharedSet(ctx context.Context, in *GetCampaignSharedSetRequest, opts ...grpc.CallOption) (*resources.CampaignSharedSet, error) {
	out := new(resources.CampaignSharedSet)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignSharedSetService/GetCampaignSharedSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignSharedSetServiceClient) MutateCampaignSharedSets(ctx context.Context, in *MutateCampaignSharedSetsRequest, opts ...grpc.CallOption) (*MutateCampaignSharedSetsResponse, error) {
	out := new(MutateCampaignSharedSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignSharedSetService/MutateCampaignSharedSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignSharedSetServiceServer is the server API for CampaignSharedSetService service.
type CampaignSharedSetServiceServer interface {
	// Returns the requested campaign shared set in full detail.
	GetCampaignSharedSet(context.Context, *GetCampaignSharedSetRequest) (*resources.CampaignSharedSet, error)
	// Creates or removes campaign shared sets. Operation statuses are returned.
	MutateCampaignSharedSets(context.Context, *MutateCampaignSharedSetsRequest) (*MutateCampaignSharedSetsResponse, error)
}

func RegisterCampaignSharedSetServiceServer(s *grpc.Server, srv CampaignSharedSetServiceServer) {
	s.RegisterService(&_CampaignSharedSetService_serviceDesc, srv)
}

func _CampaignSharedSetService_GetCampaignSharedSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignSharedSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignSharedSetServiceServer).GetCampaignSharedSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignSharedSetService/GetCampaignSharedSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignSharedSetServiceServer).GetCampaignSharedSet(ctx, req.(*GetCampaignSharedSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignSharedSetService_MutateCampaignSharedSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignSharedSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignSharedSetServiceServer).MutateCampaignSharedSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignSharedSetService/MutateCampaignSharedSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignSharedSetServiceServer).MutateCampaignSharedSets(ctx, req.(*MutateCampaignSharedSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignSharedSetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignSharedSetService",
	HandlerType: (*CampaignSharedSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignSharedSet",
			Handler:    _CampaignSharedSetService_GetCampaignSharedSet_Handler,
		},
		{
			MethodName: "MutateCampaignSharedSets",
			Handler:    _CampaignSharedSetService_MutateCampaignSharedSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_shared_set_service.proto",
}
