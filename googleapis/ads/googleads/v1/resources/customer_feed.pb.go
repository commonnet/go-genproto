// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_feed.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A customer feed.
type CustomerFeed struct {
	// The resource name of the customer feed.
	// Customer feed resource names have the form:
	//
	// `customers/{customer_id}/customerFeeds/{feed_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed being linked to the customer.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// customer. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,3,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v1.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the CustomerFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,4,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the customer feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CustomerFeed) Reset()         { *m = CustomerFeed{} }
func (m *CustomerFeed) String() string { return proto.CompactTextString(m) }
func (*CustomerFeed) ProtoMessage()    {}
func (*CustomerFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_51bd0ae9893f1b0c, []int{0}
}

func (m *CustomerFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFeed.Unmarshal(m, b)
}
func (m *CustomerFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFeed.Marshal(b, m, deterministic)
}
func (m *CustomerFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFeed.Merge(m, src)
}
func (m *CustomerFeed) XXX_Size() int {
	return xxx_messageInfo_CustomerFeed.Size(m)
}
func (m *CustomerFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFeed.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFeed proto.InternalMessageInfo

func (m *CustomerFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *CustomerFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *CustomerFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *CustomerFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CustomerFeed)(nil), "google.ads.googleads.v1.resources.CustomerFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_feed.proto", fileDescriptor_51bd0ae9893f1b0c)
}

var fileDescriptor_51bd0ae9893f1b0c = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0xbb, 0xb5, 0x60, 0xac, 0xa5, 0x9b, 0xab, 0x50, 0x8a, 0x6c, 0x95, 0xc2, 0x5e, 0xcd,
	0x34, 0xeb, 0xcf, 0x45, 0xbc, 0x31, 0x2b, 0xb6, 0x20, 0x2a, 0x4b, 0x2a, 0x8b, 0xc8, 0x4a, 0x98,
	0x26, 0x67, 0xd3, 0xd0, 0xcc, 0x0f, 0x33, 0x93, 0x4a, 0x9f, 0xc2, 0x77, 0xf0, 0xd2, 0x47, 0xf1,
	0x51, 0x7c, 0x08, 0x91, 0x64, 0x66, 0xa2, 0x5d, 0x59, 0xeb, 0xdd, 0xc9, 0x99, 0xef, 0xfb, 0xf2,
	0x9d, 0xf3, 0xcd, 0xf8, 0x4f, 0x4b, 0xce, 0xcb, 0x1a, 0x30, 0x29, 0x14, 0x36, 0x65, 0x5b, 0x5d,
	0x45, 0x58, 0x82, 0xe2, 0x8d, 0xcc, 0x41, 0xe1, 0xbc, 0x51, 0x9a, 0x53, 0x90, 0xd9, 0x0a, 0xa0,
	0x40, 0x42, 0x72, 0xcd, 0x83, 0x43, 0x83, 0x45, 0xa4, 0x50, 0xa8, 0xa7, 0xa1, 0xab, 0x08, 0xf5,
	0xb4, 0xfd, 0x67, 0x9b, 0x94, 0x73, 0x4e, 0x29, 0x67, 0x98, 0x12, 0x9d, 0x5f, 0x54, 0xac, 0xcc,
	0x56, 0x0d, 0xcb, 0x75, 0xc5, 0x99, 0x91, 0xde, 0x7f, 0xb2, 0x89, 0x07, 0xac, 0xa1, 0x0a, 0xb7,
	0x26, 0xb2, 0xba, 0x62, 0x97, 0x99, 0xd2, 0x44, 0x37, 0xea, 0xff, 0x58, 0xa2, 0x26, 0x39, 0x5c,
	0xf0, 0xba, 0x00, 0x99, 0xe9, 0x6b, 0x01, 0x96, 0xf5, 0xc0, 0xb2, 0xba, 0xaf, 0xf3, 0x66, 0x85,
	0x3f, 0x4b, 0x22, 0x04, 0x48, 0xa7, 0x7a, 0xe0, 0x54, 0x45, 0x85, 0x09, 0x63, 0x5c, 0x93, 0xd6,
	0xa8, 0x3d, 0x7d, 0xf8, 0x65, 0xe8, 0xef, 0xbc, 0xb4, 0xcb, 0x39, 0x01, 0x28, 0x82, 0x47, 0xfe,
	0x7d, 0x37, 0x7f, 0xc6, 0x08, 0x85, 0xd0, 0x1b, 0x7b, 0x93, 0xbb, 0xe9, 0x8e, 0x6b, 0xbe, 0x23,
	0x14, 0x82, 0x63, 0x7f, 0xab, 0x9d, 0x21, 0x1c, 0x8c, 0xbd, 0xc9, 0xbd, 0xe9, 0x81, 0x5d, 0x1f,
	0x72, 0x16, 0xd0, 0x99, 0x96, 0x15, 0x2b, 0x17, 0xa4, 0x6e, 0x20, 0xed, 0x90, 0x01, 0xf7, 0x47,
	0xeb, 0xfe, 0x55, 0x38, 0x1c, 0x0f, 0x27, 0xbb, 0xd3, 0x19, 0xda, 0x14, 0x44, 0x37, 0x37, 0x9a,
	0xff, 0xe6, 0xbd, 0xbf, 0x16, 0xf0, 0x8a, 0x35, 0x74, 0xbd, 0x97, 0xee, 0x89, 0x9b, 0x0d, 0x15,
	0x7c, 0xf2, 0x47, 0x7f, 0xa5, 0x13, 0x6e, 0x75, 0x7e, 0x8f, 0x37, 0xfe, 0xd0, 0xc4, 0x8a, 0xde,
	0x5a, 0xe2, 0x89, 0xe5, 0xa5, 0x7b, 0x74, 0xad, 0x13, 0x7c, 0xf0, 0xb7, 0x4d, 0x76, 0xe1, 0x9d,
	0xb1, 0x37, 0xd9, 0x9d, 0xbe, 0xb8, 0x65, 0x88, 0x76, 0xb7, 0x6f, 0x2a, 0x76, 0x79, 0xd6, 0x91,
	0xba, 0x19, 0x6e, 0xb6, 0x52, 0xab, 0x37, 0xfb, 0xe9, 0xf9, 0x47, 0x39, 0xa7, 0xe8, 0xd6, 0xdb,
	0x39, 0x1b, 0xfd, 0x19, 0xdc, 0xbc, 0xdd, 0xfd, 0xdc, 0xfb, 0xf8, 0xda, 0xf2, 0x4a, 0x5e, 0x13,
	0x56, 0x22, 0x2e, 0x4b, 0x5c, 0x02, 0xeb, 0x92, 0x71, 0x97, 0x4a, 0x54, 0xea, 0x1f, 0x6f, 0xe5,
	0x79, 0x5f, 0x7d, 0x1d, 0x0c, 0x4f, 0x93, 0xe4, 0xdb, 0xe0, 0xf0, 0xd4, 0x48, 0x26, 0x85, 0x42,
	0xa6, 0x6c, 0xab, 0x45, 0x84, 0x52, 0x87, 0xfc, 0xee, 0x30, 0xcb, 0xa4, 0x50, 0xcb, 0x1e, 0xb3,
	0x5c, 0x44, 0xcb, 0x1e, 0xf3, 0x63, 0x70, 0x64, 0x0e, 0xe2, 0x38, 0x29, 0x54, 0x1c, 0xf7, 0xa8,
	0x38, 0x5e, 0x44, 0x71, 0xdc, 0xe3, 0xce, 0xb7, 0x3b, 0xb3, 0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x67, 0xbb, 0xa2, 0x22, 0xd7, 0x03, 0x00, 0x00,
}
