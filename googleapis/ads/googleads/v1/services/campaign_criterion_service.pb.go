// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [CampaignCriterionService.GetCampaignCriterion][google.ads.googleads.v1.services.CampaignCriterionService.GetCampaignCriterion].
type GetCampaignCriterionRequest struct {
	// The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignCriterionRequest) Reset()         { *m = GetCampaignCriterionRequest{} }
func (m *GetCampaignCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignCriterionRequest) ProtoMessage()    {}
func (*GetCampaignCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{0}
}

func (m *GetCampaignCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignCriterionRequest.Unmarshal(m, b)
}
func (m *GetCampaignCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignCriterionRequest.Merge(m, src)
}
func (m *GetCampaignCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignCriterionRequest.Size(m)
}
func (m *GetCampaignCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignCriterionRequest proto.InternalMessageInfo

func (m *GetCampaignCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignCriterionService.MutateCampaignCriteria][google.ads.googleads.v1.services.CampaignCriterionService.MutateCampaignCriteria].
type MutateCampaignCriteriaRequest struct {
	// The ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual criteria.
	Operations []*CampaignCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCampaignCriteriaRequest) Reset()         { *m = MutateCampaignCriteriaRequest{} }
func (m *MutateCampaignCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriteriaRequest) ProtoMessage()    {}
func (*MutateCampaignCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{1}
}

func (m *MutateCampaignCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateCampaignCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriteriaRequest.Merge(m, src)
}
func (m *MutateCampaignCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriteriaRequest.Size(m)
}
func (m *MutateCampaignCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriteriaRequest proto.InternalMessageInfo

func (m *MutateCampaignCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignCriteriaRequest) GetOperations() []*CampaignCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign criterion.
type CampaignCriterionOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignCriterionOperation_Create
	//	*CampaignCriterionOperation_Update
	//	*CampaignCriterionOperation_Remove
	Operation            isCampaignCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CampaignCriterionOperation) Reset()         { *m = CampaignCriterionOperation{} }
func (m *CampaignCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionOperation) ProtoMessage()    {}
func (*CampaignCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{2}
}

func (m *CampaignCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionOperation.Unmarshal(m, b)
}
func (m *CampaignCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionOperation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionOperation.Merge(m, src)
}
func (m *CampaignCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionOperation.Size(m)
}
func (m *CampaignCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionOperation proto.InternalMessageInfo

func (m *CampaignCriterionOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignCriterionOperation_Operation interface {
	isCampaignCriterionOperation_Operation()
}

type CampaignCriterionOperation_Create struct {
	Create *resources.CampaignCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignCriterionOperation_Update struct {
	Update *resources.CampaignCriterion `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignCriterionOperation_Create) isCampaignCriterionOperation_Operation() {}

func (*CampaignCriterionOperation_Update) isCampaignCriterionOperation_Operation() {}

func (*CampaignCriterionOperation_Remove) isCampaignCriterionOperation_Operation() {}

func (m *CampaignCriterionOperation) GetOperation() isCampaignCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignCriterionOperation) GetCreate() *resources.CampaignCriterion {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignCriterionOperation) GetUpdate() *resources.CampaignCriterion {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionOperation_Create)(nil),
		(*CampaignCriterionOperation_Update)(nil),
		(*CampaignCriterionOperation_Remove)(nil),
	}
}

// Response message for campaign criterion mutate.
type MutateCampaignCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MutateCampaignCriteriaResponse) Reset()         { *m = MutateCampaignCriteriaResponse{} }
func (m *MutateCampaignCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriteriaResponse) ProtoMessage()    {}
func (*MutateCampaignCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{3}
}

func (m *MutateCampaignCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateCampaignCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriteriaResponse.Merge(m, src)
}
func (m *MutateCampaignCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriteriaResponse.Size(m)
}
func (m *MutateCampaignCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriteriaResponse proto.InternalMessageInfo

func (m *MutateCampaignCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignCriteriaResponse) GetResults() []*MutateCampaignCriterionResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCampaignCriterionResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignCriterionResult) Reset()         { *m = MutateCampaignCriterionResult{} }
func (m *MutateCampaignCriterionResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignCriterionResult) ProtoMessage()    {}
func (*MutateCampaignCriterionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a3c2b9bfd5dd6f9, []int{4}
}

func (m *MutateCampaignCriterionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignCriterionResult.Unmarshal(m, b)
}
func (m *MutateCampaignCriterionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignCriterionResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignCriterionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignCriterionResult.Merge(m, src)
}
func (m *MutateCampaignCriterionResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignCriterionResult.Size(m)
}
func (m *MutateCampaignCriterionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignCriterionResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignCriterionResult proto.InternalMessageInfo

func (m *MutateCampaignCriterionResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignCriterionRequest)(nil), "google.ads.googleads.v1.services.GetCampaignCriterionRequest")
	proto.RegisterType((*MutateCampaignCriteriaRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignCriteriaRequest")
	proto.RegisterType((*CampaignCriterionOperation)(nil), "google.ads.googleads.v1.services.CampaignCriterionOperation")
	proto.RegisterType((*MutateCampaignCriteriaResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignCriteriaResponse")
	proto.RegisterType((*MutateCampaignCriterionResult)(nil), "google.ads.googleads.v1.services.MutateCampaignCriterionResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_criterion_service.proto", fileDescriptor_3a3c2b9bfd5dd6f9)
}

var fileDescriptor_3a3c2b9bfd5dd6f9 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x6b, 0xd4, 0x4c,
	0x1c, 0xc7, 0x9f, 0x64, 0x1f, 0xfa, 0x3c, 0x9d, 0xad, 0x0a, 0xe3, 0xbf, 0xb0, 0xda, 0xba, 0xc4,
	0x82, 0x65, 0x0f, 0x09, 0xbb, 0xd6, 0x4b, 0x6a, 0xb1, 0xbb, 0xd5, 0xb6, 0x1e, 0xfa, 0x87, 0x14,
	0x0a, 0xca, 0x42, 0x98, 0x26, 0xd3, 0x10, 0x9a, 0x64, 0xe2, 0xcc, 0x64, 0xa5, 0x94, 0x5e, 0xc4,
	0x77, 0xe0, 0x1b, 0x10, 0x8f, 0xbe, 0x0d, 0x05, 0xc1, 0xab, 0x67, 0x6f, 0x9e, 0xc4, 0x17, 0x21,
	0x93, 0xc9, 0xac, 0x6d, 0x77, 0xe3, 0x4a, 0xbd, 0xfd, 0xf2, 0x9b, 0x6f, 0x3e, 0xf3, 0xfb, 0x37,
	0x33, 0xa0, 0x1b, 0x12, 0x12, 0xc6, 0xd8, 0x46, 0x01, 0xb3, 0xa5, 0x29, 0xac, 0x41, 0xdb, 0x66,
	0x98, 0x0e, 0x22, 0x1f, 0x33, 0xdb, 0x47, 0x49, 0x86, 0xa2, 0x30, 0xf5, 0x7c, 0x1a, 0x71, 0x4c,
	0x23, 0x92, 0x7a, 0xe5, 0x9a, 0x95, 0x51, 0xc2, 0x09, 0x6c, 0xca, 0xff, 0x2c, 0x14, 0x30, 0x6b,
	0x88, 0xb0, 0x06, 0x6d, 0x4b, 0x21, 0x1a, 0x4e, 0xd5, 0x26, 0x14, 0x33, 0x92, 0xd3, 0xf1, 0xbb,
	0x48, 0x7a, 0xe3, 0xb6, 0xfa, 0x37, 0x8b, 0x6c, 0x94, 0xa6, 0x84, 0x23, 0x1e, 0x91, 0x94, 0x95,
	0xab, 0xe5, 0xde, 0x76, 0xf1, 0xb5, 0x9f, 0x1f, 0xd8, 0x07, 0x11, 0x8e, 0x03, 0x2f, 0x41, 0xec,
	0xb0, 0x54, 0xcc, 0x9d, 0x57, 0xbc, 0xa4, 0x28, 0xcb, 0x30, 0x55, 0x84, 0x9b, 0xe5, 0x3a, 0xcd,
	0x7c, 0x9b, 0x71, 0xc4, 0xf3, 0x72, 0xc1, 0xec, 0x81, 0x5b, 0xeb, 0x98, 0xaf, 0x96, 0x71, 0xad,
	0xaa, 0xb0, 0x5c, 0xfc, 0x22, 0xc7, 0x8c, 0xc3, 0xbb, 0xe0, 0x92, 0x8a, 0xde, 0x4b, 0x51, 0x82,
	0x0d, 0xad, 0xa9, 0x2d, 0x4c, 0xbb, 0x33, 0xca, 0xb9, 0x85, 0x12, 0x6c, 0xfe, 0xd0, 0xc0, 0xec,
	0x66, 0xce, 0x11, 0xc7, 0xe7, 0x38, 0x48, 0x61, 0xee, 0x80, 0xba, 0x9f, 0x33, 0x4e, 0x12, 0x4c,
	0xbd, 0x28, 0x28, 0x21, 0x40, 0xb9, 0x9e, 0x06, 0xb0, 0x0f, 0x00, 0xc9, 0x30, 0x95, 0x59, 0x1b,
	0x7a, 0xb3, 0xb6, 0x50, 0xef, 0x3c, 0xb4, 0x26, 0x95, 0xdc, 0x1a, 0x89, 0x7b, 0x5b, 0x41, 0xdc,
	0x53, 0x3c, 0x78, 0x0f, 0x5c, 0xc9, 0x10, 0xe5, 0x11, 0x8a, 0xbd, 0x03, 0x14, 0xc5, 0x39, 0xc5,
	0x46, 0xad, 0xa9, 0x2d, 0xfc, 0xef, 0x5e, 0x2e, 0xdd, 0x6b, 0xd2, 0x2b, 0xd2, 0x1d, 0xa0, 0x38,
	0x0a, 0x10, 0xc7, 0x1e, 0x49, 0xe3, 0x23, 0xe3, 0xdf, 0x42, 0x36, 0xa3, 0x9c, 0xdb, 0x69, 0x7c,
	0x64, 0xbe, 0xd5, 0x41, 0xa3, 0x7a, 0x63, 0xb8, 0x04, 0xea, 0x79, 0x56, 0x10, 0x44, 0x7f, 0x0a,
	0x42, 0xbd, 0xd3, 0x50, 0xb9, 0xa8, 0x06, 0x59, 0x6b, 0xa2, 0x85, 0x9b, 0x88, 0x1d, 0xba, 0x40,
	0xca, 0x85, 0x0d, 0xb7, 0xc0, 0x94, 0x4f, 0x31, 0xe2, 0xb2, 0xd0, 0xf5, 0xce, 0x62, 0x65, 0x0d,
	0x86, 0x43, 0x35, 0x5a, 0x84, 0x8d, 0x7f, 0xdc, 0x92, 0x22, 0x78, 0x92, 0x6e, 0xe8, 0x7f, 0xc7,
	0x93, 0x14, 0x68, 0x80, 0x29, 0x8a, 0x13, 0x32, 0x90, 0x05, 0x9c, 0x16, 0x2b, 0xf2, 0xbb, 0x57,
	0x07, 0xd3, 0xc3, 0x8a, 0x9b, 0x1f, 0x34, 0x30, 0x57, 0x35, 0x11, 0x2c, 0x23, 0x29, 0xc3, 0x70,
	0x0d, 0x5c, 0x3f, 0xd7, 0x13, 0x0f, 0x53, 0x4a, 0x68, 0x01, 0xae, 0x77, 0xa0, 0x0a, 0x94, 0x66,
	0xbe, 0xb5, 0x5b, 0x4c, 0xac, 0x7b, 0xf5, 0x6c, 0xb7, 0x9e, 0x08, 0x39, 0x7c, 0x06, 0xfe, 0xa3,
	0x98, 0xe5, 0x31, 0x57, 0x63, 0xf3, 0x68, 0xf2, 0xd8, 0x8c, 0x0d, 0x4d, 0x0c, 0xbd, 0xe0, 0xb8,
	0x8a, 0x67, 0x3e, 0xae, 0x18, 0x6b, 0xa5, 0xfc, 0xa3, 0xd3, 0xd1, 0xf9, 0x58, 0x03, 0xc6, 0x08,
	0x60, 0x57, 0x86, 0x02, 0x3f, 0x69, 0xe0, 0xda, 0xb8, 0xf3, 0x07, 0x97, 0x27, 0x67, 0xf1, 0x9b,
	0x73, 0xdb, 0xb8, 0x50, 0x9f, 0x4d, 0xe7, 0xd5, 0x97, 0x6f, 0x6f, 0xf4, 0x45, 0xd8, 0x11, 0xb7,
	0xd6, 0xf1, 0x99, 0xd4, 0x96, 0xd5, 0x61, 0x65, 0x76, 0x6b, 0x78, 0x8d, 0xa9, 0xa6, 0xda, 0xad,
	0x13, 0xf8, 0x55, 0x03, 0x37, 0xc6, 0xb7, 0x1c, 0x5e, 0xb4, 0x23, 0xea, 0xfa, 0x68, 0xac, 0x5c,
	0x1c, 0x20, 0xa7, 0xcd, 0x5c, 0x29, 0x32, 0x73, 0xcc, 0x07, 0x22, 0xb3, 0x5f, 0xa9, 0x1c, 0x9f,
	0xba, 0x95, 0x96, 0x5b, 0x27, 0x23, 0x89, 0x39, 0x49, 0x81, 0x75, 0xb4, 0x56, 0xef, 0xb5, 0x0e,
	0xe6, 0x7d, 0x92, 0x4c, 0x8c, 0xa4, 0x37, 0x5b, 0xd5, 0xec, 0x1d, 0x71, 0xf4, 0x77, 0xb4, 0xe7,
	0x1b, 0x25, 0x22, 0x24, 0x31, 0x4a, 0x43, 0x8b, 0xd0, 0xd0, 0x0e, 0x71, 0x5a, 0x5c, 0x0c, 0xea,
	0xdd, 0xc8, 0x22, 0x56, 0xfd, 0x56, 0x2d, 0x29, 0xe3, 0x9d, 0x5e, 0x5b, 0xef, 0x76, 0xdf, 0xeb,
	0xcd, 0x75, 0x09, 0xec, 0x06, 0xcc, 0x92, 0xa6, 0xb0, 0xf6, 0xda, 0x56, 0xb9, 0x31, 0xfb, 0xac,
	0x24, 0xfd, 0x6e, 0xc0, 0xfa, 0x43, 0x49, 0x7f, 0xaf, 0xdd, 0x57, 0x92, 0xef, 0xfa, 0xbc, 0xf4,
	0x3b, 0x4e, 0x37, 0x60, 0x8e, 0x33, 0x14, 0x39, 0xce, 0x5e, 0xdb, 0x71, 0x94, 0x6c, 0x7f, 0xaa,
	0x88, 0xf3, 0xfe, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x31, 0xc3, 0xac, 0x52, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignCriterionServiceClient is the client API for CampaignCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetCampaignCriterion(ctx context.Context, in *GetCampaignCriterionRequest, opts ...grpc.CallOption) (*resources.CampaignCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error)
}

type campaignCriterionServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignCriterionServiceClient(cc *grpc.ClientConn) CampaignCriterionServiceClient {
	return &campaignCriterionServiceClient{cc}
}

func (c *campaignCriterionServiceClient) GetCampaignCriterion(ctx context.Context, in *GetCampaignCriterionRequest, opts ...grpc.CallOption) (*resources.CampaignCriterion, error) {
	out := new(resources.CampaignCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignCriterionService/GetCampaignCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignCriterionServiceClient) MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error) {
	out := new(MutateCampaignCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignCriterionService/MutateCampaignCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCriterionServiceServer is the server API for CampaignCriterionService service.
type CampaignCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetCampaignCriterion(context.Context, *GetCampaignCriterionRequest) (*resources.CampaignCriterion, error)
	// Creates, updates, or removes criteria. Operation statuses are returned.
	MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error)
}

func RegisterCampaignCriterionServiceServer(s *grpc.Server, srv CampaignCriterionServiceServer) {
	s.RegisterService(&_CampaignCriterionService_serviceDesc, srv)
}

func _CampaignCriterionService_GetCampaignCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).GetCampaignCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignCriterionService/GetCampaignCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).GetCampaignCriterion(ctx, req.(*GetCampaignCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignCriterionService_MutateCampaignCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignCriterionService/MutateCampaignCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, req.(*MutateCampaignCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignCriterionService",
	HandlerType: (*CampaignCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignCriterion",
			Handler:    _CampaignCriterionService_GetCampaignCriterion_Handler,
		},
		{
			MethodName: "MutateCampaignCriteria",
			Handler:    _CampaignCriterionService_MutateCampaignCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_criterion_service.proto",
}
