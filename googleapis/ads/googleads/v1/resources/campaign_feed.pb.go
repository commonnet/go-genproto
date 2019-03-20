// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_feed.proto

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

// A campaign feed.
type CampaignFeed struct {
	// The resource name of the campaign feed.
	// Campaign feed resource names have the form:
	//
	// `customers/{customer_id}/campaignFeeds/{campaign_id}~{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed to which the CampaignFeed belongs.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The campaign to which the CampaignFeed belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// campaign. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v1.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the CampaignFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the campaign feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignFeed) Reset()         { *m = CampaignFeed{} }
func (m *CampaignFeed) String() string { return proto.CompactTextString(m) }
func (*CampaignFeed) ProtoMessage()    {}
func (*CampaignFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_10f784f3692d2150, []int{0}
}

func (m *CampaignFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeed.Unmarshal(m, b)
}
func (m *CampaignFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeed.Marshal(b, m, deterministic)
}
func (m *CampaignFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeed.Merge(m, src)
}
func (m *CampaignFeed) XXX_Size() int {
	return xxx_messageInfo_CampaignFeed.Size(m)
}
func (m *CampaignFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeed.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeed proto.InternalMessageInfo

func (m *CampaignFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *CampaignFeed) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *CampaignFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *CampaignFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignFeed)(nil), "google.ads.googleads.v1.resources.CampaignFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_feed.proto", fileDescriptor_10f784f3692d2150)
}

var fileDescriptor_10f784f3692d2150 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0x6e, 0x5d, 0x74, 0xac, 0xa5, 0x9b, 0xab, 0x50, 0x8a, 0x6c, 0x95, 0xc2, 0x5e,
	0xcd, 0x34, 0xeb, 0x07, 0x12, 0x6f, 0xcc, 0x8a, 0x2d, 0x88, 0xca, 0x92, 0xca, 0x22, 0xb2, 0x12,
	0xa6, 0xc9, 0xd9, 0x69, 0x68, 0xe6, 0x83, 0x4c, 0x52, 0xe9, 0xeb, 0x78, 0xe9, 0x9b, 0xe8, 0xa3,
	0xf8, 0x10, 0x22, 0x49, 0x66, 0xa2, 0x5d, 0xd9, 0x76, 0xef, 0xce, 0x9e, 0xf9, 0xff, 0xce, 0x9e,
	0x73, 0xfe, 0x27, 0xe8, 0x19, 0x93, 0x92, 0xe5, 0x40, 0x68, 0xaa, 0x49, 0x1b, 0xd6, 0xd1, 0xa5,
	0x4f, 0x0a, 0xd0, 0xb2, 0x2a, 0x12, 0xd0, 0x24, 0xa1, 0x5c, 0xd1, 0x8c, 0x89, 0x78, 0x09, 0x90,
	0x62, 0x55, 0xc8, 0x52, 0xba, 0x07, 0xad, 0x16, 0xd3, 0x54, 0xe3, 0x0e, 0xc3, 0x97, 0x3e, 0xee,
	0xb0, 0xbd, 0xe7, 0xeb, 0x2a, 0x27, 0x92, 0x73, 0x29, 0x08, 0xa7, 0x65, 0x72, 0x9e, 0x09, 0x16,
	0x2f, 0x2b, 0x91, 0x94, 0x99, 0x14, 0x6d, 0xe9, 0xbd, 0xa7, 0xeb, 0x38, 0x10, 0x15, 0xd7, 0xa4,
	0x6e, 0x22, 0xce, 0x33, 0x71, 0x11, 0xeb, 0x92, 0x96, 0x95, 0xde, 0x8c, 0x52, 0x39, 0x4d, 0xe0,
	0x5c, 0xe6, 0x29, 0x14, 0x71, 0x79, 0xa5, 0xc0, 0x50, 0x0f, 0x0d, 0xd5, 0xfc, 0x3a, 0xab, 0x96,
	0xe4, 0x6b, 0x41, 0x95, 0x82, 0xc2, 0x56, 0xdd, 0xb7, 0x55, 0x55, 0x46, 0xa8, 0x10, 0xb2, 0xa4,
	0x75, 0xa3, 0xe6, 0xf5, 0xd1, 0x8f, 0x3e, 0xda, 0x7e, 0x6d, 0x96, 0x73, 0x0c, 0x90, 0xba, 0x8f,
	0xd1, 0x03, 0x3b, 0x7f, 0x2c, 0x28, 0x07, 0xcf, 0x19, 0x39, 0xe3, 0x7b, 0xd1, 0xb6, 0x4d, 0x7e,
	0xa0, 0x1c, 0xdc, 0x23, 0xb4, 0x55, 0xcf, 0xe0, 0xf5, 0x46, 0xce, 0xf8, 0xfe, 0x64, 0xdf, 0xac,
	0x0f, 0xdb, 0x16, 0xf0, 0x69, 0x59, 0x64, 0x82, 0xcd, 0x69, 0x5e, 0x41, 0xd4, 0x28, 0xdd, 0x17,
	0xe8, 0xae, 0xf5, 0xc0, 0xeb, 0x6f, 0x40, 0x75, 0x6a, 0x57, 0xa2, 0xe1, 0xea, 0xe4, 0xda, 0xdb,
	0x1a, 0xf5, 0xc7, 0x3b, 0x93, 0x29, 0x5e, 0x67, 0x61, 0xb3, 0x31, 0x3c, 0xfb, 0xcb, 0x7d, 0xbc,
	0x52, 0xf0, 0x46, 0x54, 0x7c, 0x35, 0x17, 0xed, 0xaa, 0xeb, 0x09, 0xed, 0x7e, 0x41, 0xc3, 0xff,
	0x7c, 0xf5, 0xee, 0x34, 0x3d, 0x1f, 0xad, 0xfd, 0xc3, 0xf6, 0x20, 0xf0, 0x7b, 0x03, 0x1e, 0x1b,
	0x2e, 0xda, 0xe5, 0x2b, 0x19, 0xf7, 0x13, 0x1a, 0xb4, 0xae, 0x7b, 0x83, 0x91, 0x33, 0xde, 0x99,
	0xbc, 0xba, 0x65, 0x88, 0xda, 0x95, 0x77, 0x99, 0xb8, 0x38, 0x6d, 0xa0, 0x66, 0x86, 0xeb, 0xa9,
	0xc8, 0xd4, 0x9b, 0xfe, 0x76, 0xd0, 0x61, 0x22, 0x39, 0xbe, 0xf5, 0xae, 0xa7, 0xc3, 0x7f, 0x2d,
	0x9f, 0xd5, 0xfb, 0x9f, 0x39, 0x9f, 0xdf, 0x1a, 0x8e, 0xc9, 0x9c, 0x0a, 0x86, 0x65, 0xc1, 0x08,
	0x03, 0xd1, 0xb8, 0x63, 0xcf, 0x51, 0x65, 0xfa, 0x86, 0xaf, 0xec, 0x65, 0x17, 0x7d, 0xeb, 0xf5,
	0x4f, 0xc2, 0xf0, 0x7b, 0xef, 0xe0, 0xa4, 0x2d, 0x19, 0xa6, 0x1a, 0xb7, 0x61, 0x1d, 0xcd, 0x7d,
	0x1c, 0x59, 0xe5, 0x4f, 0xab, 0x59, 0x84, 0xa9, 0x5e, 0x74, 0x9a, 0xc5, 0xdc, 0x5f, 0x74, 0x9a,
	0x5f, 0xbd, 0xc3, 0xf6, 0x21, 0x08, 0xc2, 0x54, 0x07, 0x41, 0xa7, 0x0a, 0x82, 0xb9, 0x1f, 0x04,
	0x9d, 0xee, 0x6c, 0xd0, 0x34, 0xfb, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xf6, 0xe4,
	0x5c, 0x11, 0x04, 0x00, 0x00,
}
