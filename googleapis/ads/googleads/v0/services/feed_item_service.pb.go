// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/feed_item_service.proto

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
// [FeedItemService.GetFeedItem][google.ads.googleads.v0.services.FeedItemService.GetFeedItem].
type GetFeedItemRequest struct {
	// The resource name of the feed item to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedItemRequest) Reset()         { *m = GetFeedItemRequest{} }
func (m *GetFeedItemRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedItemRequest) ProtoMessage()    {}
func (*GetFeedItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_098a5bc3e1f83ec0, []int{0}
}

func (m *GetFeedItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedItemRequest.Unmarshal(m, b)
}
func (m *GetFeedItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedItemRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedItemRequest.Merge(m, src)
}
func (m *GetFeedItemRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedItemRequest.Size(m)
}
func (m *GetFeedItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedItemRequest proto.InternalMessageInfo

func (m *GetFeedItemRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [FeedItemService.MutateFeedItems][google.ads.googleads.v0.services.FeedItemService.MutateFeedItems].
type MutateFeedItemsRequest struct {
	// The ID of the customer whose feed items are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feed items.
	Operations []*FeedItemOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateFeedItemsRequest) Reset()         { *m = MutateFeedItemsRequest{} }
func (m *MutateFeedItemsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemsRequest) ProtoMessage()    {}
func (*MutateFeedItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_098a5bc3e1f83ec0, []int{1}
}

func (m *MutateFeedItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemsRequest.Unmarshal(m, b)
}
func (m *MutateFeedItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemsRequest.Merge(m, src)
}
func (m *MutateFeedItemsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemsRequest.Size(m)
}
func (m *MutateFeedItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemsRequest proto.InternalMessageInfo

func (m *MutateFeedItemsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedItemsRequest) GetOperations() []*FeedItemOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedItemsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedItemsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an feed item.
type FeedItemOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedItemOperation_Create
	//	*FeedItemOperation_Update
	//	*FeedItemOperation_Remove
	Operation            isFeedItemOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FeedItemOperation) Reset()         { *m = FeedItemOperation{} }
func (m *FeedItemOperation) String() string { return proto.CompactTextString(m) }
func (*FeedItemOperation) ProtoMessage()    {}
func (*FeedItemOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_098a5bc3e1f83ec0, []int{2}
}

func (m *FeedItemOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemOperation.Unmarshal(m, b)
}
func (m *FeedItemOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemOperation.Marshal(b, m, deterministic)
}
func (m *FeedItemOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemOperation.Merge(m, src)
}
func (m *FeedItemOperation) XXX_Size() int {
	return xxx_messageInfo_FeedItemOperation.Size(m)
}
func (m *FeedItemOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemOperation proto.InternalMessageInfo

func (m *FeedItemOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedItemOperation_Operation interface {
	isFeedItemOperation_Operation()
}

type FeedItemOperation_Create struct {
	Create *resources.FeedItem `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedItemOperation_Update struct {
	Update *resources.FeedItem `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedItemOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedItemOperation_Create) isFeedItemOperation_Operation() {}

func (*FeedItemOperation_Update) isFeedItemOperation_Operation() {}

func (*FeedItemOperation_Remove) isFeedItemOperation_Operation() {}

func (m *FeedItemOperation) GetOperation() isFeedItemOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedItemOperation) GetCreate() *resources.FeedItem {
	if x, ok := m.GetOperation().(*FeedItemOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedItemOperation) GetUpdate() *resources.FeedItem {
	if x, ok := m.GetOperation().(*FeedItemOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedItemOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedItemOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedItemOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedItemOperation_Create)(nil),
		(*FeedItemOperation_Update)(nil),
		(*FeedItemOperation_Remove)(nil),
	}
}

// Response message for an feed item mutate.
type MutateFeedItemsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedItemResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MutateFeedItemsResponse) Reset()         { *m = MutateFeedItemsResponse{} }
func (m *MutateFeedItemsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemsResponse) ProtoMessage()    {}
func (*MutateFeedItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_098a5bc3e1f83ec0, []int{3}
}

func (m *MutateFeedItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemsResponse.Unmarshal(m, b)
}
func (m *MutateFeedItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemsResponse.Merge(m, src)
}
func (m *MutateFeedItemsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemsResponse.Size(m)
}
func (m *MutateFeedItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemsResponse proto.InternalMessageInfo

func (m *MutateFeedItemsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedItemsResponse) GetResults() []*MutateFeedItemResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed item mutate.
type MutateFeedItemResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedItemResult) Reset()         { *m = MutateFeedItemResult{} }
func (m *MutateFeedItemResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemResult) ProtoMessage()    {}
func (*MutateFeedItemResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_098a5bc3e1f83ec0, []int{4}
}

func (m *MutateFeedItemResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemResult.Unmarshal(m, b)
}
func (m *MutateFeedItemResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedItemResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemResult.Merge(m, src)
}
func (m *MutateFeedItemResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemResult.Size(m)
}
func (m *MutateFeedItemResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemResult proto.InternalMessageInfo

func (m *MutateFeedItemResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedItemRequest)(nil), "google.ads.googleads.v0.services.GetFeedItemRequest")
	proto.RegisterType((*MutateFeedItemsRequest)(nil), "google.ads.googleads.v0.services.MutateFeedItemsRequest")
	proto.RegisterType((*FeedItemOperation)(nil), "google.ads.googleads.v0.services.FeedItemOperation")
	proto.RegisterType((*MutateFeedItemsResponse)(nil), "google.ads.googleads.v0.services.MutateFeedItemsResponse")
	proto.RegisterType((*MutateFeedItemResult)(nil), "google.ads.googleads.v0.services.MutateFeedItemResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/feed_item_service.proto", fileDescriptor_098a5bc3e1f83ec0)
}

var fileDescriptor_098a5bc3e1f83ec0 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xfe, 0xed, 0xf6, 0x47, 0xb5, 0xb3, 0xd5, 0xe2, 0x58, 0x6d, 0x08, 0xa2, 0x61, 0x2d, 0x58,
	0x52, 0xdc, 0x8d, 0x89, 0x48, 0xbb, 0xa5, 0x87, 0x14, 0x9a, 0xb6, 0x87, 0xda, 0xb2, 0x85, 0x1e,
	0x24, 0xb0, 0x4c, 0xb3, 0x6f, 0xc2, 0xd2, 0xdd, 0x9d, 0x75, 0x66, 0x36, 0x52, 0x4a, 0x2f, 0x82,
	0x9f, 0xc0, 0x6f, 0x20, 0x78, 0xf1, 0xea, 0x27, 0xf0, 0xea, 0xd5, 0xab, 0x47, 0x4f, 0x7e, 0x05,
	0x11, 0x64, 0x77, 0x76, 0xd2, 0x26, 0x35, 0xc4, 0xf6, 0xf6, 0xce, 0x3b, 0xcf, 0xf3, 0xbc, 0xef,
	0xbc, 0x7f, 0x06, 0xad, 0xf4, 0x28, 0xed, 0x85, 0x60, 0x13, 0x9f, 0xdb, 0xd2, 0xcc, 0xac, 0x7e,
	0xcd, 0xe6, 0xc0, 0xfa, 0x41, 0x07, 0xb8, 0xdd, 0x05, 0xf0, 0xbd, 0x40, 0x40, 0xe4, 0x15, 0x2e,
	0x2b, 0x61, 0x54, 0x50, 0x5c, 0x91, 0x70, 0x8b, 0xf8, 0xdc, 0x1a, 0x30, 0xad, 0x7e, 0xcd, 0x52,
	0xcc, 0xf2, 0xb3, 0x71, 0xda, 0x0c, 0x38, 0x4d, 0xd9, 0x90, 0xb8, 0x14, 0x2d, 0x3f, 0x50, 0x94,
	0x24, 0xb0, 0x49, 0x1c, 0x53, 0x41, 0x44, 0x40, 0x63, 0x5e, 0xdc, 0x16, 0x21, 0xed, 0xfc, 0x74,
	0x94, 0x76, 0xed, 0x6e, 0x00, 0xa1, 0xef, 0x45, 0x84, 0x1f, 0x17, 0x88, 0x87, 0xa3, 0x88, 0x37,
	0x8c, 0x24, 0x09, 0x30, 0xa5, 0xb0, 0x50, 0xdc, 0xb3, 0xa4, 0x63, 0x73, 0x41, 0x44, 0x5a, 0x5c,
	0x98, 0xab, 0x08, 0x6f, 0x81, 0x68, 0x01, 0xf8, 0x3b, 0x02, 0x22, 0x17, 0x5e, 0xa7, 0xc0, 0x05,
	0x7e, 0x8c, 0x6e, 0xa9, 0x5c, 0xbd, 0x98, 0x44, 0x50, 0xd2, 0x2a, 0xda, 0xd2, 0x8c, 0x3b, 0xab,
	0x9c, 0x2f, 0x49, 0x04, 0xe6, 0x77, 0x0d, 0xdd, 0xdf, 0x4d, 0x05, 0x11, 0xa0, 0xe8, 0x5c, 0xf1,
	0x1f, 0x21, 0xa3, 0x93, 0x72, 0x41, 0x23, 0x60, 0x5e, 0xe0, 0x17, 0x6c, 0xa4, 0x5c, 0x3b, 0x3e,
	0x3e, 0x40, 0x88, 0x26, 0xc0, 0xe4, 0x2b, 0x4b, 0x7a, 0x65, 0x6a, 0xc9, 0xa8, 0x37, 0xac, 0x49,
	0x95, 0xb5, 0x54, 0xa0, 0x3d, 0xc5, 0x75, 0x2f, 0xc8, 0xe0, 0x27, 0x68, 0x2e, 0x21, 0x4c, 0x04,
	0x24, 0xf4, 0xba, 0x24, 0x08, 0x53, 0x06, 0xa5, 0xa9, 0x8a, 0xb6, 0x74, 0xd3, 0xbd, 0x5d, 0xb8,
	0x5b, 0xd2, 0x9b, 0x3d, 0xaf, 0x4f, 0xc2, 0xc0, 0x27, 0x02, 0x3c, 0x1a, 0x87, 0x27, 0xa5, 0xff,
	0x73, 0xd8, 0xac, 0x72, 0xee, 0xc5, 0xe1, 0x89, 0xf9, 0x4e, 0x47, 0x77, 0x2e, 0xc5, 0xc3, 0x6b,
	0xc8, 0x48, 0x93, 0x9c, 0x98, 0x55, 0x3f, 0x27, 0x1a, 0xf5, 0xb2, 0xca, 0x5c, 0x95, 0xdf, 0x6a,
	0x65, 0x0d, 0xda, 0x25, 0xfc, 0xd8, 0x45, 0x12, 0x9e, 0xd9, 0x78, 0x13, 0x4d, 0x77, 0x18, 0x10,
	0x21, 0xeb, 0x69, 0xd4, 0x97, 0xc7, 0xbe, 0x78, 0x30, 0x29, 0x83, 0x27, 0x6f, 0xff, 0xe7, 0x16,
	0xe4, 0x4c, 0x46, 0x8a, 0x96, 0xf4, 0x6b, 0xc9, 0x48, 0x32, 0x2e, 0xa1, 0x69, 0x06, 0x11, 0xed,
	0xcb, 0x2a, 0xcd, 0x64, 0x37, 0xf2, 0xbc, 0x61, 0xa0, 0x99, 0x41, 0x59, 0xcd, 0xcf, 0x1a, 0x5a,
	0xb8, 0xd4, 0x66, 0x9e, 0xd0, 0x98, 0x03, 0x6e, 0xa1, 0x7b, 0x23, 0x15, 0xf7, 0x80, 0x31, 0xca,
	0x72, 0x45, 0xa3, 0x8e, 0x55, 0x62, 0x2c, 0xe9, 0x58, 0x07, 0xf9, 0xd8, 0xb9, 0x77, 0x87, 0x7b,
	0xb1, 0x99, 0xc1, 0xf1, 0x3e, 0xba, 0xc1, 0x80, 0xa7, 0xa1, 0x50, 0xb3, 0xf0, 0x62, 0xf2, 0x2c,
	0x0c, 0xe7, 0xe4, 0xe6, 0x74, 0x57, 0xc9, 0x98, 0x6b, 0x68, 0xfe, 0x6f, 0x80, 0x7f, 0x9a, 0xec,
	0xfa, 0x6f, 0x1d, 0xcd, 0x29, 0xde, 0x81, 0x8c, 0x87, 0x3f, 0x6a, 0xc8, 0xb8, 0xb0, 0x29, 0xf8,
	0xf9, 0xe4, 0x0c, 0x2f, 0x2f, 0x56, 0xf9, 0x2a, 0xad, 0x32, 0x1b, 0x6f, 0xbf, 0xfd, 0x78, 0xaf,
	0x3f, 0xc5, 0xcb, 0xd9, 0xdf, 0x71, 0x3a, 0x94, 0xf6, 0xba, 0xda, 0x25, 0x6e, 0x57, 0xf3, 0xcf,
	0x24, 0xef, 0x8b, 0x5d, 0x3d, 0xc3, 0x5f, 0x34, 0x34, 0x37, 0xd2, 0x2e, 0xbc, 0x72, 0xd5, 0x6a,
	0xaa, 0x45, 0x2e, 0xaf, 0x5e, 0x83, 0x29, 0x67, 0xc3, 0x5c, 0xcd, 0xb3, 0x6f, 0x98, 0x56, 0x96,
	0xfd, 0x79, 0xba, 0xa7, 0x17, 0x3e, 0x86, 0xf5, 0xea, 0xd9, 0x79, 0xf2, 0x4e, 0x94, 0x0b, 0x39,
	0x5a, 0x75, 0xe3, 0x97, 0x86, 0x16, 0x3b, 0x34, 0x9a, 0x18, 0x7b, 0x63, 0x7e, 0xa4, 0x4b, 0xfb,
	0xd9, 0xfe, 0xed, 0x6b, 0xaf, 0xb6, 0x0b, 0x66, 0x8f, 0x86, 0x24, 0xee, 0x59, 0x94, 0xf5, 0xec,
	0x1e, 0xc4, 0xf9, 0x76, 0xaa, 0x1f, 0x39, 0x09, 0xf8, 0xf8, 0xcf, 0x7f, 0x4d, 0x19, 0x1f, 0xf4,
	0xa9, 0xad, 0x66, 0xf3, 0x93, 0x5e, 0xd9, 0x92, 0x82, 0x4d, 0x9f, 0x5b, 0xd2, 0xcc, 0xac, 0xc3,
	0x9a, 0x55, 0x04, 0xe6, 0x5f, 0x15, 0xa4, 0xdd, 0xf4, 0x79, 0x7b, 0x00, 0x69, 0x1f, 0xd6, 0xda,
	0x0a, 0xf2, 0x53, 0x5f, 0x94, 0x7e, 0xc7, 0x69, 0xfa, 0xdc, 0x71, 0x06, 0x20, 0xc7, 0x39, 0xac,
	0x39, 0x8e, 0x82, 0x1d, 0x4d, 0xe7, 0x79, 0x36, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0x3b,
	0xdb, 0x9f, 0xa3, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedItemServiceClient is the client API for FeedItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedItemServiceClient interface {
	// Returns the requested feed item in full detail.
	GetFeedItem(ctx context.Context, in *GetFeedItemRequest, opts ...grpc.CallOption) (*resources.FeedItem, error)
	// Creates, updates, or removes feed items. Operation statuses are
	// returned.
	MutateFeedItems(ctx context.Context, in *MutateFeedItemsRequest, opts ...grpc.CallOption) (*MutateFeedItemsResponse, error)
}

type feedItemServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedItemServiceClient(cc *grpc.ClientConn) FeedItemServiceClient {
	return &feedItemServiceClient{cc}
}

func (c *feedItemServiceClient) GetFeedItem(ctx context.Context, in *GetFeedItemRequest, opts ...grpc.CallOption) (*resources.FeedItem, error) {
	out := new(resources.FeedItem)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedItemService/GetFeedItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedItemServiceClient) MutateFeedItems(ctx context.Context, in *MutateFeedItemsRequest, opts ...grpc.CallOption) (*MutateFeedItemsResponse, error) {
	out := new(MutateFeedItemsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedItemService/MutateFeedItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemServiceServer is the server API for FeedItemService service.
type FeedItemServiceServer interface {
	// Returns the requested feed item in full detail.
	GetFeedItem(context.Context, *GetFeedItemRequest) (*resources.FeedItem, error)
	// Creates, updates, or removes feed items. Operation statuses are
	// returned.
	MutateFeedItems(context.Context, *MutateFeedItemsRequest) (*MutateFeedItemsResponse, error)
}

func RegisterFeedItemServiceServer(s *grpc.Server, srv FeedItemServiceServer) {
	s.RegisterService(&_FeedItemService_serviceDesc, srv)
}

func _FeedItemService_GetFeedItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemServiceServer).GetFeedItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedItemService/GetFeedItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemServiceServer).GetFeedItem(ctx, req.(*GetFeedItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedItemService_MutateFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemServiceServer).MutateFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedItemService/MutateFeedItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemServiceServer).MutateFeedItems(ctx, req.(*MutateFeedItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.FeedItemService",
	HandlerType: (*FeedItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedItem",
			Handler:    _FeedItemService_GetFeedItem_Handler,
		},
		{
			MethodName: "MutateFeedItems",
			Handler:    _FeedItemService_MutateFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/feed_item_service.proto",
}
